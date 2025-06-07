package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"

	mcp "github.com/metoro-io/mcp-golang"
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

// QuizState holds the current state of the quiz
type QuizState struct {
	Responses []politicalcompass.Response `json:"responses"`
}

// Reset state helper function for tests
func resetState() {
	totalEconomicScore = 0.0
	totalSocialScore = 0.0
	questionCount = 0
	shuffledQuestions = nil
	currentIndex = 0
	quizState = &QuizState{}
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
			"%s\n\n"+
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
			"Please respond with: strongly_disagree, disagree, agree, or strongly_agree\n"+
			"After answering, call this tool again to continue to the next question.",
			questionCount, len(politicalcompass.AllQuestions), question.Text)
	} else {
		message = fmt.Sprintf("✅ Response recorded!\n\n"+
			"Progress: %d of %d questions completed\n\n"+
			"Question %d of %d:\n%s\n\n"+
			"Please respond with: strongly_disagree, disagree, agree, or strongly_agree\n"+
			"After answering, call this tool again to continue to the next question.",
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

**Current Scores:** (based on %d responses)
- Economic axis: %.2f (%.2f%% toward %s)
- Social axis: %.2f (%.2f%% toward %s)

**Current Quadrant:** %s

**Response Distribution:**
`, answered, totalQuestions, remaining,
		float64(answered)/float64(totalQuestions)*100,
		answered,
		economicScore, abs(economicScore)/10*100,
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
