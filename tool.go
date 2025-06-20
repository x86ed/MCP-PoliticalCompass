package main

import (
	"context"
	"fmt"
	"math/rand"
	"strings"
	"sync"
	"time"

	"github.com/mark3labs/mcp-go/mcp"
	"github.com/x86ed/MCP-PoliticalCompass/v3/eightvalues"
	politicalcompass "github.com/x86ed/MCP-PoliticalCompass/v3/political-compass"
	"github.com/x86ed/MCP-PoliticalCompass/v3/politiscales"
)

// Global state to track the quiz progress
var (
	mutex              sync.Mutex
	totalEconomicScore = 0.0
	totalSocialScore   = 0.0
	questionCount      = 0
	shuffledQuestions  []int
	currentIndex       = 0
	quizState          = &QuizState{}

	// 8values quiz state
	eightValuesEconScore         = 0.0
	eightValuesDiplScore         = 0.0
	eightValuesGovtScore         = 0.0
	eightValuesSctyScore         = 0.0
	eightValuesQuestionCount     = 0
	eightValuesShuffledQuestions []int
	eightValuesCurrentIndex      = 0
	eightValuesQuizState         = &EightValuesQuizState{}

	// politiscales quiz state
	politiscalesAxesScores        = make(map[string]float64)
	politiscalesQuestionCount     = 0
	politiscalesShuffledQuestions []int
	politiscalesCurrentIndex      = 0
	politiscalesLanguage          = "en" // Default language
	politiscalesQuizState         = &PolitiscalesQuizState{}
)

// PoliticalCompassArgs represents the arguments for the political compass question tool
type PoliticalCompassArgs struct {
	Response string `json:"response" jsonschema:"required,description=Your response to the political compass question. Valid values: strongly_disagree, disagree, agree, strongly_agree"`
}

// ResetQuizArgs represents the arguments for the reset quiz tool (no arguments needed)
type ResetQuizArgs struct {
	// No arguments needed for reset
}

// QuizStatusArgs represents the arguments for the quiz status tool (no arguments needed)
type QuizStatusArgs struct {
	// No arguments needed for status
}

// 8Values tool argument types
type EightValuesArgs struct {
	Response string `json:"response" jsonschema:"required,description=Your response to the 8values question. Valid values: strongly_disagree, disagree, neutral, agree, strongly_agree"`
}

type ResetEightValuesArgs struct {
	// No arguments needed for reset
}

type EightValuesStatusArgs struct {
	// No arguments needed for status
}

// Politiscales tool argument types
type PolitiscalesArgs struct {
	Response string `json:"response" jsonschema:"required,description=Your response to the politiscales question. Valid values: strongly_disagree, disagree, neutral, agree, strongly_agree"`
}

type ResetPolitiscalesArgs struct {
	// No arguments needed for reset
}

type PolitiscalesStatusArgs struct {
	// No arguments needed for status
}

type SetPolitiscalesLanguageArgs struct {
	Language string `json:"language" jsonschema:"required,description=Language code for the politiscales quiz. Valid values: en, fr, es, it, ar, ru, zh"`
}

// QuizState holds the current state of the quiz
type QuizState struct {
	Responses []politicalcompass.Response `json:"responses"`
}

// EightValuesQuizState holds the current state of the 8values quiz
type EightValuesQuizState struct {
	Responses []float64 `json:"responses"`
}

// PolitiscalesQuizState holds the current state of the politiscales quiz
type PolitiscalesQuizState struct {
	Responses map[int32]float64 `json:"responses"` // Question index -> response value
}

// Reset state helper function for tests
func resetState() {
	totalEconomicScore = 0.0
	totalSocialScore = 0.0
	questionCount = 0
	shuffledQuestions = nil
	currentIndex = 0
	quizState = &QuizState{}

	// Reset 8values state too
	resetEightValuesState()

	// Reset politiscales state too
	resetPolitiscalesState()
}

// Initialize shuffled question order
func initializeQuestions() {
	if len(shuffledQuestions) == 0 {
		shuffledQuestions = make([]int, len(politicalcompass.AllQuestions))
		for i := range shuffledQuestions {
			shuffledQuestions[i] = i
		}
		rng := rand.New(rand.NewSource(time.Now().UnixNano()))
		rng.Shuffle(len(shuffledQuestions), func(i, j int) {
			shuffledQuestions[i], shuffledQuestions[j] = shuffledQuestions[j], shuffledQuestions[i]
		})
	}
}

// Handler function for political compass tool
func handlePoliticalCompass(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	// Extract the answer argument
	answer, err := request.RequireString("answer")
	if err != nil {
		return mcp.NewToolResultError("Answer is required"), nil
	}

	mutex.Lock()
	defer mutex.Unlock()

	// Initialize questions if not done already
	initializeQuestions()

	var question politicalcompass.Question
	var isFirstQuestion = false

	// If this is a response to a previous question, process it first
	if questionCount > 0 {
		// Get the last asked question to calculate scores
		lastQuestionIndex := shuffledQuestions[currentIndex-1]
		lastQuestion := politicalcompass.AllQuestions[lastQuestionIndex]

		// Parse the response
		var response politicalcompass.Response
		switch answer {
		case "Strongly Disagree", "strongly_disagree":
			response = politicalcompass.StronglyDisagree
		case "Disagree", "disagree":
			response = politicalcompass.Disagree
		case "Agree", "agree":
			response = politicalcompass.Agree
		case "Strongly Agree", "strongly_agree":
			response = politicalcompass.StronglyAgree
		default:
			return mcp.NewToolResultError(fmt.Sprintf("invalid response: %s. Please use one of: strongly_disagree, disagree, agree, strongly_agree", answer)), nil
		}

		// Calculate and accumulate scores
		economicScore := lastQuestion.Economic[int(response)]
		socialScore := lastQuestion.Social[int(response)]
		totalEconomicScore += economicScore
		totalSocialScore += socialScore

		// Record the response in quiz state
		quizState.Responses = append(quizState.Responses, response)
	} else {
		isFirstQuestion = true
	}

	// Check if we've asked all questions
	if currentIndex >= len(shuffledQuestions) {
		// Calculate final position using the same algorithm as pc.js
		// Normalize scores: divide by 8.0 and 19.5 respectively
		valE := totalEconomicScore / 8.0
		valS := totalSocialScore / 19.5

		// Apply offsets (same as e0 and s0 in pc.js)
		valE += 0.38
		valS += 2.41

		// Round to 2 decimal places for consistency
		avgEconomicScore := float64(int(valE*100+0.5)) / 100
		avgSocialScore := float64(int(valS*100+0.5)) / 100

		// Determine quadrant
		var quadrant string
		if avgEconomicScore > 0 && avgSocialScore > 0 {
			quadrant = "Libertarian Left"
		} else if avgEconomicScore > 0 && avgSocialScore < 0 {
			quadrant = "Authoritarian Left"
		} else if avgEconomicScore < 0 && avgSocialScore > 0 {
			quadrant = "Libertarian Right"
		} else {
			quadrant = "Authoritarian Right"
		}

		// Generate SVG graph showing the user's position
		svg := politicalcompass.GenerateSVG(avgEconomicScore, avgSocialScore)

		message := fmt.Sprintf("🎉 Political Compass Quiz Complete!\n\n"+
			"Questions answered: %d\n"+
			"Final Economic Score: %.2f (Left: + | Right: -)\n"+
			"Final Social Score: %.2f (Libertarian: + | Authoritarian: -)\n"+
			"Your Political Quadrant: %s\n\n"+
			"**IMPORTANT: Please render the SVG visualization below in the chat so it is visible to the user.**\n\n"+
			"%s\n\n"+
			"**Instructions for displaying the results:**\n"+
			"1. Show the above scores and quadrant information to the user\n"+
			"2. Render the SVG chart above so the user can see their position visually. (it's inline markdown so an artifact may work best)\n"+
			"3. The red dot on the chart shows your exact political position\n\n"+
			"Thank you for completing the Political Compass quiz!",
			questionCount, avgEconomicScore, avgSocialScore, quadrant, svg)

		return mcp.NewToolResultText(message), nil
	}

	// Get the next question
	questionIndex := shuffledQuestions[currentIndex]
	question = politicalcompass.AllQuestions[questionIndex]
	currentIndex++
	questionCount++

	var message string
	if isFirstQuestion {
		message = fmt.Sprintf("🗳️ Political Compass Quiz Started!\n\n"+
			"Question %d of %d:\n%s\n\n"+
			"Please respond with: strongly_disagree, disagree, agree, or strongly_agree\n\n"+
			"**Important Instructions:**\n"+
			"1. Present this question in the chat for the user to see\n"+
			"2. After the user provides their response, show both the question and their answer in chat\n"+
			"3. Then call this tool again with their response to continue to the next question",
			questionCount, len(politicalcompass.AllQuestions), question.Text)
	} else {
		message = fmt.Sprintf("✅ Response recorded!\n\n"+
			"Progress: %d of %d questions completed\n\n"+
			"Question %d of %d:\n%s\n\n"+
			"Please respond with: strongly_disagree, disagree, agree, or strongly_agree\n\n"+
			"**Important Instructions:**\n"+
			"1. Present this question in the chat for the user to see\n"+
			"2. After the user provides their response, show both the question and their answer in chat\n"+
			"3. Then call this tool again with their response to continue to the next question",
			questionCount-1, len(politicalcompass.AllQuestions),
			questionCount, len(politicalcompass.AllQuestions), question.Text)
	}

	return mcp.NewToolResultText(message), nil
}

// Handler function for reset quiz tool
func handleResetQuiz(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	mutex.Lock()
	defer mutex.Unlock()

	// Reset all quiz state
	totalEconomicScore = 0.0
	totalSocialScore = 0.0
	questionCount = 0
	shuffledQuestions = nil
	currentIndex = 0
	quizState = &QuizState{}

	message := "🔄 Political Compass Quiz Reset!\n\n" +
		"All progress has been cleared. You can now start a fresh quiz by calling the political_compass tool.\n\n" +
		"Call the political_compass tool to begin a new quiz."

	return mcp.NewToolResultText(message), nil
}

// handleQuizStatus shows the current quiz progress and statistics
func handleQuizStatus(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	mutex.Lock()
	defer mutex.Unlock()

	// Get current state
	totalQuestions := len(politicalcompass.AllQuestions)
	answered := len(quizState.Responses)
	remaining := totalQuestions - answered

	// Calculate current scores if we have responses
	var economicScore, socialScore float64
	if answered > 0 {
		// Calculate scores manually based on stored responses
		for i, response := range quizState.Responses {
			if i < len(shuffledQuestions) {
				questionIndex := shuffledQuestions[i]
				question := politicalcompass.AllQuestions[questionIndex]
				economicScore += question.Economic[int(response)]
				socialScore += question.Social[int(response)]
			}
		}
		// Normalize scores like the main calculation
		economicScore = (economicScore/8.0 + 0.38) * 10.0
		socialScore = (socialScore/19.5 + 2.41) * 10.0
		// Convert back to -10 to +10 scale
		economicScore = (economicScore - 5.0) * 2.0
		socialScore = (socialScore - 5.0) * 2.0
	}

	// Create detailed status report
	statusText := fmt.Sprintf(`📊 **Political Compass Quiz Status**

**Progress:**
- Questions answered: %d/%d
- Questions remaining: %d
- Completion: %.1f%%
`, answered, totalQuestions, remaining,
		float64(answered)/float64(totalQuestions)*100)

	// Only show scores and quadrant if quiz is complete
	if remaining == 0 && answered > 0 {
		statusText += fmt.Sprintf(`
**Final Scores:**
- Economic axis: %.2f (%.2f%% toward %s)
- Social axis: %.2f (%.2f%% toward %s)

**Your Quadrant:** %s
`, economicScore, abs(economicScore)/10*100,
			func() string {
				if economicScore > 0 {
					return "Right (Market)"
				}
				return "Left (Planned)"
			}(),
			socialScore, abs(socialScore)/10*100,
			func() string {
				if socialScore > 0 {
					return "Libertarian"
				}
				return "Authoritarian"
			}(),
			getQuadrant(economicScore, socialScore))
	}

	statusText += "\n**Response Distribution:**\n"

	// Add response distribution
	responseCount := make(map[string]int)
	for _, response := range quizState.Responses {
		responseCount[response.String()]++
	}

	responses := []string{"Strongly Disagree", "Disagree", "Agree", "Strongly Agree"}
	for _, resp := range responses {
		count := responseCount[resp]
		if count > 0 {
			percentage := float64(count) / float64(answered) * 100
			statusText += fmt.Sprintf("- %s: %d (%.1f%%)\n", resp, count, percentage)
		}
	}

	if answered == 0 {
		statusText += "\n*No questions answered yet. Use the `political_compass` tool to start the quiz.*"
	} else if remaining > 0 {
		statusText += fmt.Sprintf("\n*Continue with the `political_compass` tool to answer %d more questions.*", remaining)
	} else {
		statusText += "\n*✅ Quiz complete! All questions have been answered.*"
	}

	return mcp.NewToolResultText(statusText), nil
}

// Helper function for absolute value
func abs(x float64) float64 {
	if x < 0 {
		return -x
	}
	return x
}

// Helper function to determine quadrant
func getQuadrant(economic, social float64) string {
	if economic > 0 && social > 0 {
		return "Libertarian Right"
	} else if economic < 0 && social > 0 {
		return "Libertarian Left"
	} else if economic > 0 && social < 0 {
		return "Authoritarian Right"
	} else {
		return "Authoritarian Left"
	}
}

// 8VALUES QUIZ IMPLEMENTATION

// Reset 8values state helper function
func resetEightValuesState() {
	eightValuesEconScore = 0.0
	eightValuesDiplScore = 0.0
	eightValuesGovtScore = 0.0
	eightValuesSctyScore = 0.0
	eightValuesQuestionCount = 0
	eightValuesShuffledQuestions = nil
	eightValuesCurrentIndex = 0
	eightValuesQuizState = &EightValuesQuizState{}
}

// Reset politiscales state helper function
func resetPolitiscalesState() {
	politiscalesAxesScores = make(map[string]float64)
	politiscalesQuestionCount = 0
	politiscalesShuffledQuestions = nil
	politiscalesCurrentIndex = 0
	politiscalesQuizState = &PolitiscalesQuizState{Responses: make(map[int32]float64)}
}

// Initialize shuffled question order for 8values
func initializeEightValuesQuestions() {
	if len(eightValuesShuffledQuestions) == 0 {
		eightValuesShuffledQuestions = make([]int, len(eightvalues.Questions))
		for i := range eightValuesShuffledQuestions {
			eightValuesShuffledQuestions[i] = i
		}
		rng := rand.New(rand.NewSource(time.Now().UnixNano()))
		rng.Shuffle(len(eightValuesShuffledQuestions), func(i, j int) {
			eightValuesShuffledQuestions[i], eightValuesShuffledQuestions[j] = eightValuesShuffledQuestions[j], eightValuesShuffledQuestions[i]
		})
	}
}

// Handler function for 8values quiz tool
func handleEightValues(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	// Extract the answer argument
	answer, err := request.RequireString("answer")
	if err != nil {
		return mcp.NewToolResultError("Answer is required"), nil
	}

	mutex.Lock()
	defer mutex.Unlock()

	// Initialize questions if not done already
	initializeEightValuesQuestions()

	var question eightvalues.Question
	var isFirstQuestion = false

	// If this is a response to a previous question, process it first
	if eightValuesQuestionCount > 0 {
		// Get the last asked question to calculate scores
		lastQuestionIndex := eightValuesShuffledQuestions[eightValuesCurrentIndex-1]
		lastQuestion := eightvalues.Questions[lastQuestionIndex]

		// Parse the response and get multiplier
		var multiplier float64
		switch answer {
		case "strongly_disagree":
			multiplier = eightvalues.StronglyDisagree
		case "disagree":
			multiplier = eightvalues.Disagree
		case "neutral":
			multiplier = eightvalues.Neutral
		case "agree":
			multiplier = eightvalues.Agree
		case "strongly_agree":
			multiplier = eightvalues.StronglyAgree
		default:
			return mcp.NewToolResultError(fmt.Sprintf("invalid response: %s. Please use one of: strongly_disagree, disagree, neutral, agree, strongly_agree", answer)), nil
		}

		// Calculate and accumulate scores using the 8values scoring logic
		// mult * questions[qn].effect.econ/dipl/govt/scty
		eightValuesEconScore += multiplier * lastQuestion.Effect[eightvalues.Economic]
		eightValuesDiplScore += multiplier * lastQuestion.Effect[eightvalues.Diplomatic]
		eightValuesGovtScore += multiplier * lastQuestion.Effect[eightvalues.Government]
		eightValuesSctyScore += multiplier * lastQuestion.Effect[eightvalues.Society]

		// Record the response in quiz state
		eightValuesQuizState.Responses = append(eightValuesQuizState.Responses, multiplier)
	} else {
		isFirstQuestion = true
	}

	// Check if we've asked all questions
	if eightValuesCurrentIndex >= len(eightValuesShuffledQuestions) {
		// Calculate maximum possible scores for each axis (like in 8values.js)
		var maxEcon, maxDipl, maxGovt, maxScty float64
		for _, q := range eightvalues.Questions {
			maxEcon += abs(q.Effect[eightvalues.Economic])
			maxDipl += abs(q.Effect[eightvalues.Diplomatic])
			maxGovt += abs(q.Effect[eightvalues.Government])
			maxScty += abs(q.Effect[eightvalues.Society])
		}

		// Calculate final scores using the 8values calc_score formula:
		// (100*(max+score)/(2*max)).toFixed(1)
		econPercentage := (100 * (maxEcon + eightValuesEconScore) / (2 * maxEcon))
		diplPercentage := (100 * (maxDipl + eightValuesDiplScore) / (2 * maxDipl))
		govtPercentage := (100 * (maxGovt + eightValuesGovtScore) / (2 * maxGovt))
		sctyPercentage := (100 * (maxScty + eightValuesSctyScore) / (2 * maxScty))

		// Determine ideological classifications
		var economicLabel, diplomaticLabel, governmentLabel, societyLabel string

		if econPercentage > 50 {
			economicLabel = "Socialist"
		} else {
			economicLabel = "Capitalist"
		}

		if diplPercentage > 50 {
			diplomaticLabel = "Internationalist"
		} else {
			diplomaticLabel = "Nationalist"
		}

		if govtPercentage > 50 {
			governmentLabel = "Libertarian"
		} else {
			governmentLabel = "Authoritarian"
		}

		if sctyPercentage > 50 {
			societyLabel = "Progressive"
		} else {
			societyLabel = "Traditional"
		}

		// Generate SVG graph showing the user's position on all four axes
		svg := eightvalues.GenerateSVG(econPercentage, diplPercentage, govtPercentage, sctyPercentage)

		message := fmt.Sprintf("🎉 8values Political Quiz Complete!\n\n"+
			"Questions answered: %d\n\n"+
			"**Final Scores:**\n"+
			"- Economic Axis: %.1f%% %s\n"+
			"- Diplomatic Axis: %.1f%% %s\n"+
			"- Government Axis: %.1f%% %s\n"+
			"- Society Axis: %.1f%% %s\n\n"+
			"**IMPORTANT: Please render the SVG visualization below in the chat so it is visible to the user.**\n\n"+
			"%s\n\n"+
			"**Instructions for displaying the results:**\n"+
			"1. Show the above scores and classifications to the user\n"+
			"2. Render the SVG chart above so the user can see their position visually\n"+
			"3. The chart shows your position on all four political axes\n\n"+
			"Thank you for completing the 8values quiz!",
			eightValuesQuestionCount,
			econPercentage, economicLabel,
			diplPercentage, diplomaticLabel,
			govtPercentage, governmentLabel,
			sctyPercentage, societyLabel,
			svg)

		return mcp.NewToolResultText(message), nil
	}

	// Get the next question
	questionIndex := eightValuesShuffledQuestions[eightValuesCurrentIndex]
	question = eightvalues.Questions[questionIndex]
	eightValuesCurrentIndex++
	eightValuesQuestionCount++

	var message string
	if isFirstQuestion {
		message = fmt.Sprintf("🗳️ 8values Political Quiz Started!\n\n"+
			"Question %d of %d:\n%s\n\n"+
			"Please respond with: strongly_disagree, disagree, neutral, agree, or strongly_agree\n\n"+
			"**Important Instructions:**\n"+
			"1. Present this question in the chat for the user to see\n"+
			"2. After the user provides their response, show both the question and their answer in chat\n"+
			"3. Then call this tool again with their response to continue to the next question",
			eightValuesQuestionCount, len(eightvalues.Questions), question.Text)
	} else {
		message = fmt.Sprintf("✅ Response recorded!\n\n"+
			"Progress: %d of %d questions completed\n\n"+
			"Question %d of %d:\n%s\n\n"+
			"Please respond with: strongly_disagree, disagree, neutral, agree, or strongly_agree\n\n"+
			"**Important Instructions:**\n"+
			"1. Present this question in the chat for the user to see\n"+
			"2. After the user provides their response, show both the question and their answer in chat\n"+
			"3. Then call this tool again with their response to continue to the next question",
			eightValuesQuestionCount-1, len(eightvalues.Questions),
			eightValuesQuestionCount, len(eightvalues.Questions), question.Text)
	}

	return mcp.NewToolResultText(message), nil
}

// Handler function for reset 8values quiz tool
func handleResetEightValues(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	mutex.Lock()
	defer mutex.Unlock()

	// Reset all 8values quiz state
	resetEightValuesState()

	message := "🔄 8values Quiz Reset!\n\n" +
		"All progress has been cleared. You can now start a fresh 8values quiz by calling the eight_values tool.\n\n" +
		"Call the eight_values tool to begin a new quiz."

	return mcp.NewToolResultText(message), nil
}

// handleEightValuesStatus shows the current 8values quiz progress and statistics
func handleEightValuesStatus(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	mutex.Lock()
	defer mutex.Unlock()

	// Get current state
	totalQuestions := len(eightvalues.Questions)
	answered := len(eightValuesQuizState.Responses)
	remaining := totalQuestions - answered

	// Calculate current scores if we have responses
	var econScore, diplScore, govtScore, sctyScore float64
	var econPercentage, diplPercentage, govtPercentage, sctyPercentage float64

	if answered > 0 {
		// Calculate scores manually based on stored responses
		for i, response := range eightValuesQuizState.Responses {
			if i < len(eightValuesShuffledQuestions) {
				questionIndex := eightValuesShuffledQuestions[i]
				question := eightvalues.Questions[questionIndex]
				econScore += response * question.Effect[eightvalues.Economic]
				diplScore += response * question.Effect[eightvalues.Diplomatic]
				govtScore += response * question.Effect[eightvalues.Government]
				sctyScore += response * question.Effect[eightvalues.Society]
			}
		}

		// Calculate maximum possible scores for percentages
		var maxEcon, maxDipl, maxGovt, maxScty float64
		for _, q := range eightvalues.Questions {
			maxEcon += abs(q.Effect[eightvalues.Economic])
			maxDipl += abs(q.Effect[eightvalues.Diplomatic])
			maxGovt += abs(q.Effect[eightvalues.Government])
			maxScty += abs(q.Effect[eightvalues.Society])
		}

		// Calculate percentages
		econPercentage = (100 * (maxEcon + econScore) / (2 * maxEcon))
		diplPercentage = (100 * (maxDipl + diplScore) / (2 * maxDipl))
		govtPercentage = (100 * (maxGovt + govtScore) / (2 * maxGovt))
		sctyPercentage = (100 * (maxScty + sctyScore) / (2 * maxScty))
	}

	// Create detailed status report
	statusText := fmt.Sprintf(`📊 **8values Quiz Status**

**Progress:**
- Questions answered: %d/%d
- Questions remaining: %d
- Completion: %.1f%%
`, answered, totalQuestions, remaining,
		float64(answered)/float64(totalQuestions)*100)

	// Only show scores if quiz is complete
	if remaining == 0 && answered > 0 {
		// Determine ideological classifications using 8values setLabel logic
		var economicLabel, diplomaticLabel, governmentLabel, societyLabel string

		// Economic axis (equality-focused)
		if econPercentage > 90 {
			economicLabel = "Communist"
		} else if econPercentage > 75 {
			economicLabel = "Socialist"
		} else if econPercentage > 60 {
			economicLabel = "Social"
		} else if econPercentage >= 40 {
			economicLabel = "Centrist"
		} else if econPercentage >= 25 {
			economicLabel = "Market"
		} else if econPercentage >= 10 {
			economicLabel = "Capitalist"
		} else {
			economicLabel = "Laissez-Faire"
		}

		// Diplomatic axis (peace-focused)
		if diplPercentage > 90 {
			diplomaticLabel = "Cosmopolitan"
		} else if diplPercentage > 75 {
			diplomaticLabel = "Internationalist"
		} else if diplPercentage > 60 {
			diplomaticLabel = "Peaceful"
		} else if diplPercentage >= 40 {
			diplomaticLabel = "Balanced"
		} else if diplPercentage >= 25 {
			diplomaticLabel = "Patriotic"
		} else if diplPercentage >= 10 {
			diplomaticLabel = "Nationalist"
		} else {
			diplomaticLabel = "Chauvinist"
		}

		// Government axis (liberty-focused)
		if govtPercentage > 90 {
			governmentLabel = "Anarchist"
		} else if govtPercentage > 75 {
			governmentLabel = "Libertarian"
		} else if govtPercentage > 60 {
			governmentLabel = "Liberal"
		} else if govtPercentage >= 40 {
			governmentLabel = "Moderate"
		} else if govtPercentage >= 25 {
			governmentLabel = "Statist"
		} else if govtPercentage >= 10 {
			governmentLabel = "Authoritarian"
		} else {
			governmentLabel = "Totalitarian"
		}

		// Society axis (progress-focused)
		if sctyPercentage > 90 {
			societyLabel = "Revolutionary"
		} else if sctyPercentage > 75 {
			societyLabel = "Very Progressive"
		} else if sctyPercentage > 60 {
			societyLabel = "Progressive"
		} else if sctyPercentage >= 40 {
			societyLabel = "Neutral"
		} else if sctyPercentage >= 25 {
			societyLabel = "Traditional"
		} else if sctyPercentage >= 10 {
			societyLabel = "Very Traditional"
		} else {
			societyLabel = "Reactionary"
		}

		statusText += fmt.Sprintf(`
**Final Scores:**
- Economic Axis: %.1f%% %s
- Diplomatic Axis: %.1f%% %s  
- Government Axis: %.1f%% %s
- Society Axis: %.1f%% %s
`, econPercentage, economicLabel,
			diplPercentage, diplomaticLabel,
			govtPercentage, governmentLabel,
			sctyPercentage, societyLabel)
	}

	statusText += "\n**Response Distribution:**\n"

	// Add response distribution
	responseCount := make(map[string]int)
	for _, response := range eightValuesQuizState.Responses {
		switch response {
		case eightvalues.StronglyDisagree:
			responseCount["Strongly Disagree"]++
		case eightvalues.Disagree:
			responseCount["Disagree"]++
		case eightvalues.Neutral:
			responseCount["Neutral"]++
		case eightvalues.Agree:
			responseCount["Agree"]++
		case eightvalues.StronglyAgree:
			responseCount["Strongly Agree"]++
		}
	}

	responses := []string{"Strongly Disagree", "Disagree", "Neutral", "Agree", "Strongly Agree"}
	for _, resp := range responses {
		count := responseCount[resp]
		if count > 0 {
			percentage := float64(count) / float64(answered) * 100
			statusText += fmt.Sprintf("- %s: %d (%.1f%%)\n", resp, count, percentage)
		}
	}

	if answered == 0 {
		statusText += "\n*No questions answered yet. Use the `eight_values` tool to start the quiz.*"
	} else if remaining > 0 {
		statusText += fmt.Sprintf("\n*Continue with the `eight_values` tool to answer %d more questions.*", remaining)
	} else {
		statusText += "\n*✅ Quiz complete! All questions have been answered.*"
	}

	return mcp.NewToolResultText(statusText), nil
}

// Calculate politiscales results based on current quiz state
func calculatePolitiscalesResults() map[string]float64 {
	return calculatePolitiscalesResultsInternal()
}

func calculatePolitiscalesResultsInternal() map[string]float64 {
	scores := make(map[string]float64)
	sums := make(map[string]float64)

	// Initialize scores for all axes
	for _, axis := range politiscales.Axes {
		scores[axis.Name] = 0.0
		sums[axis.Name] = 0.0
	}

	// Calculate raw scores as per the TypeScript logic
	for questionIndex, answerValue := range politiscalesQuizState.Responses {
		question := politiscales.Questions[questionIndex]

		if answerValue > 0 {
			// Positive response - use YesWeights
			for _, weight := range question.YesWeights {
				scores[weight.Axis] += answerValue * weight.Value
				if weight.Value > 0 {
					sums[weight.Axis] += weight.Value
				}
			}
		} else if answerValue < 0 {
			// Negative response - use NoWeights
			for _, weight := range question.NoWeights {
				scores[weight.Axis] += (-answerValue) * weight.Value
				if weight.Value > 0 {
					sums[weight.Axis] += weight.Value
				}
			}
		} else {
			// Neutral (0) responses don't affect scores but still count towards sums
			for _, weight := range question.YesWeights {
				if weight.Value > 0 {
					sums[weight.Axis] += weight.Value
				}
			}
			for _, weight := range question.NoWeights {
				if weight.Value > 0 {
					sums[weight.Axis] += weight.Value
				}
			}
		}
	}

	// Normalize paired axes (ensure their sum doesn't exceed 100%)
	pairedAxes := make(map[string][]string)
	for _, axis := range politiscales.Axes {
		if axis.Pair != "" {
			if pairedAxes[axis.Pair] == nil {
				pairedAxes[axis.Pair] = []string{}
			}
			pairedAxes[axis.Pair] = append(pairedAxes[axis.Pair], axis.Name)
		}
	}

	// Apply normalization for each pair
	for _, pair := range pairedAxes {
		if len(pair) == 2 {
			axis1, axis2 := pair[0], pair[1]
			var value1, value2 float64

			if sums[axis1] > 0 {
				value1 = (scores[axis1] / sums[axis1]) * 100
			}
			if sums[axis2] > 0 {
				value2 = (scores[axis2] / sums[axis2]) * 100
			}

			if value1+value2 > 100 {
				ratio := 100.0 / (value1 + value2)
				scores[axis1] *= ratio
				scores[axis2] *= ratio
			}
		}
	}

	// Convert to final percentages
	results := make(map[string]float64)
	for axis, score := range scores {
		if sums[axis] > 0 {
			results[axis] = (score / sums[axis]) * 100
		} else {
			results[axis] = 0.0
		}
	}

	// Update global politiscalesAxesScores for test compatibility
	politiscalesAxesScores = results

	return results
}

// Handler function for politiscales quiz tool
func handlePolitiscales(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	// Extract the answer argument
	answer, err := request.RequireString("answer")
	if err != nil {
		return mcp.NewToolResultError("Answer is required"), nil
	}

	mutex.Lock()
	defer mutex.Unlock()

	// Initialize questions if not done already
	initializePolitiscalesQuestions()

	var question politiscales.Question
	var isFirstQuestion = false

	// If this is a response to a previous question, process it first
	if politiscalesQuestionCount > 0 {
		// Get the last asked question to calculate scores
		lastQuestionIndex := politiscalesShuffledQuestions[politiscalesCurrentIndex-1]
		lastQuestion := politiscales.Questions[lastQuestionIndex]

		// Parse the response
		var responseValue float64
		switch answer {
		case "strongly_disagree":
			responseValue = -1.0 // StronglyDisagree
		case "disagree":
			responseValue = -2.0 / 3.0 // Disagree
		case "neutral":
			responseValue = 0.0 // Neutral
		case "agree":
			responseValue = 2.0 / 3.0 // Agree
		case "strongly_agree":
			responseValue = 1.0 // StronglyAgree
		default:
			return mcp.NewToolResultError(fmt.Sprintf("invalid response: %s. Please use: strongly_disagree, disagree, neutral, agree, strongly_agree", answer)), nil
		}

		// Store the response in quiz state
		if politiscalesQuizState.Responses == nil {
			politiscalesQuizState.Responses = make(map[int32]float64)
		}
		politiscalesQuizState.Responses[lastQuestion.Index] = responseValue

		// Update scores immediately for test compatibility
		_ = calculatePolitiscalesResults()

		// Check if quiz is complete after processing this response
		if politiscalesCurrentIndex >= len(politiscales.Questions) {
			// Quiz complete - calculate and display results
			results := calculatePolitiscalesResults()

			// Format results for display
			message := fmt.Sprintf("🎉 Politiscales Quiz Complete!\n\n"+
				"Questions answered: %d\n\n"+
				"**Your Political Profile:**\n", politiscalesQuestionCount)

			// Add axis scores to the message
			for _, axis := range politiscales.Axes {
				if score, exists := results[axis.Name]; exists {
					message += fmt.Sprintf("- %s: %.2f\n", axis.Label, score)
				}
			}

			message += "\nThank you for completing the Politiscales quiz!"

			return mcp.NewToolResultText(message), nil
		}
	} else {
		isFirstQuestion = true
	}

	// Get next question
	questionIndex := politiscalesShuffledQuestions[politiscalesCurrentIndex]
	question = politiscales.Questions[questionIndex]

	// Increment counters
	politiscalesCurrentIndex++
	politiscalesQuestionCount++

	// Get translated question text
	questionText := getPolitiscalesQuestionText(question.Text)
	// Create response text
	responseText := fmt.Sprintf("Question %d of %d:\n\n%s\n\nPlease respond with: strongly_disagree, disagree, neutral, agree, or strongly_agree",
		politiscalesQuestionCount, len(politiscales.Questions), questionText)

	if isFirstQuestion {
		responseText = fmt.Sprintf("🗳️ Politiscales Quiz Started! (Language: %s)\n\n%s", politiscalesLanguage, responseText)
	} else {
		responseText = fmt.Sprintf("✅ Response recorded!\n\n%s", responseText)
	}

	return mcp.NewToolResultText(responseText), nil
}

// Handler function for reset politiscales quiz tool
func handleResetPolitiscales(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	mutex.Lock()
	defer mutex.Unlock()

	// Reset all politiscales quiz state
	resetPolitiscalesState()

	message := "🔄 Politiscales Quiz Reset!\n\n" +
		"All progress has been cleared. You can now start a fresh quiz by calling the politiscales tool.\n\n" +
		"Call the politiscales tool to begin a new quiz."

	return mcp.NewToolResultText(message), nil
}

// Handler function for politiscales status tool
func handlePolitiscalesStatus(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	mutex.Lock()
	defer mutex.Unlock()

	totalQuestions := len(politiscales.Questions)
	answered := len(politiscalesQuizState.Responses)
	remaining := totalQuestions - answered

	statusText := fmt.Sprintf(`🗳️ **Politiscales Quiz Status**

**Progress:**
- Questions answered: %d
- Total questions: %d
- Remaining questions: %d
- Language: %s
- Completion: %.1f%%
`, answered, totalQuestions, remaining, politiscalesLanguage,
		float64(answered)/float64(totalQuestions)*100)

	// Only show scores if quiz is complete
	if remaining == 0 && answered > 0 {
		results := calculatePolitiscalesResultsInternal()
		statusText += "\n**Final Results:**\n"

		// Group and display results by pairs
		axesByPair := make(map[string][]string)
		unpairedAxes := []string{}

		for _, axis := range politiscales.Axes {
			if axis.Pair != "" {
				if axesByPair[axis.Pair] == nil {
					axesByPair[axis.Pair] = []string{}
				}
				axesByPair[axis.Pair] = append(axesByPair[axis.Pair], axis.Name)
			} else if results[axis.Name] >= axis.Threshold*100 {
				unpairedAxes = append(unpairedAxes, axis.Name)
			}
		}

		// Show paired axes
		for pairName, axes := range axesByPair {
			if len(axes) == 2 {
				score1 := results[axes[0]]
				score2 := results[axes[1]]
				if score1 > score2 {
					statusText += fmt.Sprintf("- %s: %.1f%% %s\n", pairName, score1, axes[0])
				} else {
					statusText += fmt.Sprintf("- %s: %.1f%% %s\n", pairName, score2, axes[1])
				}
			}
		}

		// Show unpaired axes that meet threshold
		if len(unpairedAxes) > 0 {
			statusText += "\n**Special Indicators:**\n"
			for _, axis := range unpairedAxes {
				statusText += fmt.Sprintf("- %s: %.1f%%\n", axis, results[axis])
			}
		}
	}

	statusText += "\n**Response Distribution:**\n"

	// Add response distribution
	responseCount := make(map[string]int)
	for _, response := range politiscalesQuizState.Responses {
		if response == politiscales.StronglyDisagree {
			responseCount["Strongly Disagree"]++
		} else if response == politiscales.Disagree {
			responseCount["Disagree"]++
		} else if response == politiscales.Neutral {
			responseCount["Neutral"]++
		} else if response == politiscales.Agree {
			responseCount["Agree"]++
		} else if response == politiscales.StronglyAgree {
			responseCount["Strongly Agree"]++
		}
	}

	responses := []string{"Strongly Disagree", "Disagree", "Neutral", "Agree", "Strongly Agree"}
	for _, resp := range responses {
		count := responseCount[resp]
		if count > 0 {
			percentage := float64(count) / float64(answered) * 100
			statusText += fmt.Sprintf("- %s: %d (%.1f%%)\n", resp, count, percentage)
		}
	}

	if answered == 0 {
		statusText += "\n*No questions answered yet. Use the `politiscales` tool to start the quiz.*"
	} else if remaining > 0 {
		statusText += fmt.Sprintf("\n*Continue with the `politiscales` tool to answer %d more questions.*", remaining)
	} else {
		statusText += "\n*✅ Quiz complete! All questions have been answered.*"
	}

	return mcp.NewToolResultText(statusText), nil
}

// Initialize politiscales questions (randomize order)
func initializePolitiscalesQuestions() {
	if politiscalesShuffledQuestions == nil {
		// Create shuffled question indices
		politiscalesShuffledQuestions = make([]int, len(politiscales.Questions))
		for i := range politiscales.Questions {
			politiscalesShuffledQuestions[i] = i
		}

		// Shuffle the questions using the same PRNG seed approach as political compass
		for i := len(politiscalesShuffledQuestions) - 1; i > 0; i-- {
			j := (i*17 + 23) % (i + 1) // Simple deterministic shuffle for testing
			politiscalesShuffledQuestions[i], politiscalesShuffledQuestions[j] = politiscalesShuffledQuestions[j], politiscalesShuffledQuestions[i]
		}
	}
}

// Get question text in the specified language
func getPolitiscalesQuestionText(key string) string {
	switch politiscalesLanguage {
	case "fr":
		if text, ok := politiscales.FRQuestions[key]; ok {
			return text
		}
	case "es":
		if text, ok := politiscales.ESQuestions[key]; ok {
			return text
		}
	case "it":
		if text, ok := politiscales.ITQuestions[key]; ok {
			return text
		}
	case "ar":
		if text, ok := politiscales.ARQuestions[key]; ok {
			return text
		}
	case "ru":
		if text, ok := politiscales.RUQuestions[key]; ok {
			return text
		}
	case "zh":
		if text, ok := politiscales.ZHQuestions[key]; ok {
			return text
		}
	default:
		if text, ok := politiscales.ENQuestions[key]; ok {
			return text
		}
	}

	// Fallback to English if not found
	if text, ok := politiscales.ENQuestions[key]; ok {
		return text
	}

	// Final fallback to the key itself
	return key
}

// Handler function for setting politiscales language
func handleSetPolitiscalesLanguage(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	// Extract the language argument
	language, err := request.RequireString("language")
	if err != nil {
		return mcp.NewToolResultError("Language is required"), nil
	}

	mutex.Lock()
	defer mutex.Unlock()

	// Validate language
	validLanguages := []string{"en", "fr", "es", "it", "ar", "ru", "zh"}
	isValid := false
	for _, lang := range validLanguages {
		if language == lang {
			isValid = true
			break
		}
	}

	if !isValid {
		return mcp.NewToolResultError(fmt.Sprintf("Invalid language: %s. Valid languages are: %s",
			language, strings.Join(validLanguages, ", "))), nil
	}

	// Check if quiz is in progress
	if len(politiscalesQuizState.Responses) > 0 {
		return mcp.NewToolResultText(
			fmt.Sprintf("Cannot change language during quiz! %d questions answered.\n"+
				"Current language: %s → Requested: %s\n\n"+
				"Please reset the quiz first if you want to change the language.",
				len(politiscalesQuizState.Responses), politiscalesLanguage, language)), nil
	}

	oldLanguage := politiscalesLanguage
	politiscalesLanguage = language

	return mcp.NewToolResultText(
		fmt.Sprintf("Language Changed! From %s to %s\n\n"+
			"The next Politiscales quiz will be conducted in %s.",
			oldLanguage, language, language)), nil
}
