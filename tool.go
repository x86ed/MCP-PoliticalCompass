package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"

	mcp "github.com/metoro-io/mcp-golang"
	"github.com/x86ed/MCP-PoliticalCompass/v2/eightvalues"
	politicalcompass "github.com/x86ed/MCP-PoliticalCompass/v2/political-compass"
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

// QuizState holds the current state of the quiz
type QuizState struct {
	Responses []politicalcompass.Response `json:"responses"`
}

// EightValuesQuizState holds the current state of the 8values quiz
type EightValuesQuizState struct {
	Responses []float64 `json:"responses"`
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
func handlePoliticalCompass(args PoliticalCompassArgs) (*mcp.ToolResponse, error) {
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
		switch args.Response {
		case "Strongly Disagree", "strongly_disagree":
			response = politicalcompass.StronglyDisagree
		case "Disagree", "disagree":
			response = politicalcompass.Disagree
		case "Agree", "agree":
			response = politicalcompass.Agree
		case "Strongly Agree", "strongly_agree":
			response = politicalcompass.StronglyAgree
		default:
			return nil, fmt.Errorf("invalid response: %s. Please use one of: strongly_disagree, disagree, agree, strongly_agree", args.Response)
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
		svg := generatePoliticalCompassSVG(avgEconomicScore, avgSocialScore)

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

		return mcp.NewToolResponse(mcp.NewTextContent(message)), nil
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

	return mcp.NewToolResponse(mcp.NewTextContent(message)), nil
}

// Handler function for reset quiz tool
func handleResetQuiz(args ResetQuizArgs) (*mcp.ToolResponse, error) {
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

	return mcp.NewToolResponse(mcp.NewTextContent(message)), nil
}

// generatePoliticalCompassSVG creates an SVG visualization of the user's political position
func generatePoliticalCompassSVG(economicScore, socialScore float64) string {
	// SVG dimensions and margins
	width := 400
	height := 400
	margin := 50

	// Calculate center point
	centerX := width / 2
	centerY := height / 2

	// Calculate user position on the graph
	// Economic score: -10 to +10 maps to left-right (0 to width-2*margin)
	// Social score: -10 to +10 maps to top-bottom (inverted: +10 is top)
	userX := centerX + int(economicScore*(float64(width-2*margin)/20.0))
	userY := centerY - int(socialScore*(float64(height-2*margin)/20.0))

	// Clamp values to stay within bounds
	if userX < margin {
		userX = margin
	}
	if userX > width-margin {
		userX = width - margin
	}
	if userY < margin {
		userY = margin
	}
	if userY > height-margin {
		userY = height - margin
	}

	svg := fmt.Sprintf(`<svg width="%d" height="%d" xmlns="http://www.w3.org/2000/svg">
  <!-- Background -->
  <rect width="%d" height="%d" fill="#f8f9fa" stroke="#dee2e6" stroke-width="1"/>
  
  <!-- Quadrant backgrounds -->
  <!-- Authoritarian Left (top-left) -->
  <rect x="%d" y="%d" width="%d" height="%d" fill="#ffebee" opacity="0.7"/>
  <!-- Authoritarian Right (top-right) -->
  <rect x="%d" y="%d" width="%d" height="%d" fill="#fff3e0" opacity="0.7"/>
  <!-- Libertarian Left (bottom-left) -->
  <rect x="%d" y="%d" width="%d" height="%d" fill="#e8f5e8" opacity="0.7"/>
  <!-- Libertarian Right (bottom-right) -->
  <rect x="%d" y="%d" width="%d" height="%d" fill="#e3f2fd" opacity="0.7"/>
  
  <!-- Grid lines -->
  <!-- Vertical center line -->
  <line x1="%d" y1="%d" x2="%d" y2="%d" stroke="#6c757d" stroke-width="2"/>
  <!-- Horizontal center line -->
  <line x1="%d" y1="%d" x2="%d" y2="%d" stroke="#6c757d" stroke-width="2"/>
  
  <!-- Grid marks -->
  <!-- Economic axis marks -->
  <line x1="%d" y1="%d" x2="%d" y2="%d" stroke="#adb5bd" stroke-width="1"/>
  <line x1="%d" y1="%d" x2="%d" y2="%d" stroke="#adb5bd" stroke-width="1"/>
  <!-- Social axis marks -->
  <line x1="%d" y1="%d" x2="%d" y2="%d" stroke="#adb5bd" stroke-width="1"/>
  <line x1="%d" y1="%d" x2="%d" y2="%d" stroke="#adb5bd" stroke-width="1"/>
  
  <!-- Axis labels -->
  <text x="%d" y="%d" text-anchor="middle" font-family="Arial, sans-serif" font-size="12" fill="#495057">Economic</text>
  <text x="%d" y="%d" text-anchor="middle" font-family="Arial, sans-serif" font-size="12" fill="#495057">Left</text>
  <text x="%d" y="%d" text-anchor="middle" font-family="Arial, sans-serif" font-size="12" fill="#495057">Right</text>
  <text x="%d" y="%d" text-anchor="middle" font-family="Arial, sans-serif" font-size="12" fill="#495057" transform="rotate(-90 %d %d)">Social</text>
  <text x="%d" y="%d" text-anchor="middle" font-family="Arial, sans-serif" font-size="12" fill="#495057" transform="rotate(-90 %d %d)">Libertarian</text>
  <text x="%d" y="%d" text-anchor="middle" font-family="Arial, sans-serif" font-size="12" fill="#495057" transform="rotate(-90 %d %d)">Authoritarian</text>
  
  <!-- Quadrant labels -->
  <text x="%d" y="%d" text-anchor="middle" font-family="Arial, sans-serif" font-size="11" font-weight="bold" fill="#666">Authoritarian</text>
  <text x="%d" y="%d" text-anchor="middle" font-family="Arial, sans-serif" font-size="11" font-weight="bold" fill="#666">Left</text>
  <text x="%d" y="%d" text-anchor="middle" font-family="Arial, sans-serif" font-size="11" font-weight="bold" fill="#666">Authoritarian</text>
  <text x="%d" y="%d" text-anchor="middle" font-family="Arial, sans-serif" font-size="11" font-weight="bold" fill="#666">Right</text>
  <text x="%d" y="%d" text-anchor="middle" font-family="Arial, sans-serif" font-size="11" font-weight="bold" fill="#666">Libertarian</text>
  <text x="%d" y="%d" text-anchor="middle" font-family="Arial, sans-serif" font-size="11" font-weight="bold" fill="#666">Left</text>
  <text x="%d" y="%d" text-anchor="middle" font-family="Arial, sans-serif" font-size="11" font-weight="bold" fill="#666">Libertarian</text>
  <text x="%d" y="%d" text-anchor="middle" font-family="Arial, sans-serif" font-size="11" font-weight="bold" fill="#666">Right</text>
  
  <!-- User position -->
  <circle cx="%d" cy="%d" r="8" fill="#dc3545" stroke="#ffffff" stroke-width="2"/>
  <text x="%d" y="%d" text-anchor="middle" font-family="Arial, sans-serif" font-size="10" font-weight="bold" fill="#ffffff">●</text>
  
  <!-- Position coordinates -->
  <text x="%d" y="%d" text-anchor="start" font-family="Arial, sans-serif" font-size="10" fill="#495057">
    Position: (%.2f, %.2f)
  </text>
</svg>`,
		width, height, // SVG dimensions
		width, height, // Background rect
		margin, margin, centerX-margin, centerY-margin, // Auth Left
		centerX, margin, centerX-margin, centerY-margin, // Auth Right
		margin, centerY, centerX-margin, centerY-margin, // Lib Left
		centerX, centerY, centerX-margin, centerY-margin, // Lib Right
		centerX, margin, centerX, height-margin, // Vertical center line
		margin, centerY, width-margin, centerY, // Horizontal center line
		margin+(centerX-margin)/2, centerY-5, margin+(centerX-margin)/2, centerY+5, // Economic left mark
		centerX+(centerX-margin)/2, centerY-5, centerX+(centerX-margin)/2, centerY+5, // Economic right mark
		centerX-5, margin+(centerY-margin)/2, centerX+5, margin+(centerY-margin)/2, // Social top mark
		centerX-5, centerY+(centerY-margin)/2, centerX+5, centerY+(centerY-margin)/2, // Social bottom mark
		centerX, height-15, // Economic label
		margin-5, centerY+15, // Left label
		width-margin+5, centerY+15, // Right label
		15, centerY, 15, centerY, // Social label
		15, margin+(centerY-margin)/2, 15, margin+(centerY-margin)/2, // Libertarian label
		15, centerY+(centerY-margin)/2, 15, centerY+(centerY-margin)/2, // Authoritarian label
		margin+(centerX-margin)/2, margin+15, // Auth Left label
		margin+(centerX-margin)/2, margin+25, // Auth Left label 2
		centerX+(centerX-margin)/2, margin+15, // Auth Right label
		centerX+(centerX-margin)/2, margin+25, // Auth Right label 2
		margin+(centerX-margin)/2, height-margin+15, // Lib Left label
		margin+(centerX-margin)/2, height-margin+25, // Lib Left label 2
		centerX+(centerX-margin)/2, height-margin+15, // Lib Right label
		centerX+(centerX-margin)/2, height-margin+25, // Lib Right label 2
		userX, userY, // User position circle
		userX, userY+1, // User position text
		margin, height-30, // Position text
		economicScore, socialScore, // Position coordinates
	)

	return svg
}

// handleQuizStatus shows the current quiz progress and statistics
func handleQuizStatus(args QuizStatusArgs) (*mcp.ToolResponse, error) {
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

	return mcp.NewToolResponse(mcp.NewTextContent(statusText)), nil
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
func handleEightValues(args EightValuesArgs) (*mcp.ToolResponse, error) {
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
		switch args.Response {
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
			return nil, fmt.Errorf("invalid response: %s. Please use one of: strongly_disagree, disagree, neutral, agree, strongly_agree", args.Response)
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

		return mcp.NewToolResponse(mcp.NewTextContent(message)), nil
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

	return mcp.NewToolResponse(mcp.NewTextContent(message)), nil
}

// Handler function for reset 8values quiz tool
func handleResetEightValues(args ResetEightValuesArgs) (*mcp.ToolResponse, error) {
	mutex.Lock()
	defer mutex.Unlock()

	// Reset all 8values quiz state
	resetEightValuesState()

	message := "🔄 8values Quiz Reset!\n\n" +
		"All progress has been cleared. You can now start a fresh 8values quiz by calling the eight_values tool.\n\n" +
		"Call the eight_values tool to begin a new quiz."

	return mcp.NewToolResponse(mcp.NewTextContent(message)), nil
}

// handleEightValuesStatus shows the current 8values quiz progress and statistics
func handleEightValuesStatus(args EightValuesStatusArgs) (*mcp.ToolResponse, error) {
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

	return mcp.NewToolResponse(mcp.NewTextContent(statusText)), nil
}


