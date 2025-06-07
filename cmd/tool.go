package main

import (
	"fmt"
	"math/rand"
	"time"

	mcp "github.com/metoro-io/mcp-golang"
	"github.com/x86ed/MCP-PoliticalCompass/polcomp"
)

// Global state to track the quiz progress
var (
	totalEconomicScore = 0.0
	totalSocialScore   = 0.0
	questionCount      = 0
	shuffledQuestions  []int
	currentIndex       = 0
)

// PoliticalCompassArgs represents the arguments for the political compass question tool
type PoliticalCompassArgs struct {
	Response string `json:"response" jsonschema:"required,enum=Strongly Disagree,enum=Disagree,enum=Agree,enum=Strongly Agree,description=Your response to the political compass question"`
}

// ResetQuizArgs represents the arguments for the reset quiz tool (no arguments needed)
type ResetQuizArgs struct {
	// No arguments needed for reset
}

// Reset state helper function for tests
func resetState() {
	totalEconomicScore = 0.0
	totalSocialScore = 0.0
	questionCount = 0
	shuffledQuestions = nil
	currentIndex = 0
}

// Initialize shuffled question order
func initializeQuestions() {
	if len(shuffledQuestions) == 0 {
		shuffledQuestions = make([]int, len(polcomp.AllQuestions))
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
	// Initialize questions if not done already
	initializeQuestions()

	var question polcomp.Question
	var isFirstQuestion = false

	// If this is a response to a previous question, process it first
	if questionCount > 0 {
		// Get the last asked question to calculate scores
		lastQuestionIndex := shuffledQuestions[currentIndex-1]
		lastQuestion := polcomp.AllQuestions[lastQuestionIndex]

		// Parse the response
		var response polcomp.Response
		switch args.Response {
		case "Strongly Disagree":
			response = polcomp.StronglyDisagree
		case "Disagree":
			response = polcomp.Disagree
		case "Agree":
			response = polcomp.Agree
		case "Strongly Agree":
			response = polcomp.StronglyAgree
		default:
			return nil, fmt.Errorf("invalid response: %s", args.Response)
		}

		// Calculate and accumulate scores
		economicScore := lastQuestion.Economic[int(response)]
		socialScore := lastQuestion.Social[int(response)]
		totalEconomicScore += economicScore
		totalSocialScore += socialScore
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
	question = polcomp.AllQuestions[questionIndex]
	currentIndex++
	questionCount++

	var message string
	if isFirstQuestion {
		message = fmt.Sprintf("🗳️ Political Compass Quiz Started!\n\n"+
			"Question %d of %d:\n%s\n\n"+
			"Please respond with: Strongly Disagree, Disagree, Agree, or Strongly Agree\n"+
			"After answering, call this tool again to continue to the next question.",
			questionCount, len(polcomp.AllQuestions), question.Text)
	} else {
		// Show progress and current scores
		avgEconomicScore := totalEconomicScore / float64(questionCount-1)
		avgSocialScore := totalSocialScore / float64(questionCount-1)

		message = fmt.Sprintf("✅ Response recorded!\n\n"+
			"Progress: %d of %d questions completed\n"+
			"Current Economic Score: %.2f (Left: + | Right: -)\n"+
			"Current Social Score: %.2f (Libertarian: + | Authoritarian: -)\n\n"+
			"Question %d of %d:\n%s\n\n"+
			"Please respond with: Strongly Disagree, Disagree, Agree, or Strongly Agree\n"+
			"After answering, call this tool again to continue to the next question.",
			questionCount-1, len(polcomp.AllQuestions), avgEconomicScore, avgSocialScore,
			questionCount, len(polcomp.AllQuestions), question.Text)
	}

	return mcp.NewToolResponse(mcp.NewTextContent(message)), nil
}

// Handler function for reset quiz tool
func handleResetQuiz(args ResetQuizArgs) (*mcp.ToolResponse, error) {
	// Reset all quiz state
	totalEconomicScore = 0.0
	totalSocialScore = 0.0
	questionCount = 0
	shuffledQuestions = nil
	currentIndex = 0

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
		userX = width-margin
	}
	if userY < margin {
		userY = margin
	}
	if userY > height-margin {
		userY = height-margin
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
		width, height,                                    // SVG dimensions
		width, height,                                    // Background rect
		margin, margin, centerX-margin, centerY-margin,  // Auth Left
		centerX, margin, centerX-margin, centerY-margin, // Auth Right  
		margin, centerY, centerX-margin, centerY-margin, // Lib Left
		centerX, centerY, centerX-margin, centerY-margin, // Lib Right
		centerX, margin, centerX, height-margin,          // Vertical center line
		margin, centerY, width-margin, centerY,           // Horizontal center line
		margin+(centerX-margin)/2, centerY-5, margin+(centerX-margin)/2, centerY+5,     // Economic left mark
		centerX+(centerX-margin)/2, centerY-5, centerX+(centerX-margin)/2, centerY+5,   // Economic right mark
		centerX-5, margin+(centerY-margin)/2, centerX+5, margin+(centerY-margin)/2,     // Social top mark
		centerX-5, centerY+(centerY-margin)/2, centerX+5, centerY+(centerY-margin)/2,   // Social bottom mark
		centerX, height-15,                               // Economic label
		margin-5, centerY+15,                             // Left label
		width-margin+5, centerY+15,                       // Right label
		15, centerY, 15, centerY,                         // Social label
		15, margin+(centerY-margin)/2, 15, margin+(centerY-margin)/2,     // Libertarian label
		15, centerY+(centerY-margin)/2, 15, centerY+(centerY-margin)/2,   // Authoritarian label
		margin+(centerX-margin)/2, margin+15,             // Auth Left label
		margin+(centerX-margin)/2, margin+25,             // Auth Left label 2
		centerX+(centerX-margin)/2, margin+15,            // Auth Right label
		centerX+(centerX-margin)/2, margin+25,            // Auth Right label 2
		margin+(centerX-margin)/2, height-margin+15,      // Lib Left label
		margin+(centerX-margin)/2, height-margin+25,      // Lib Left label 2
		centerX+(centerX-margin)/2, height-margin+15,     // Lib Right label
		centerX+(centerX-margin)/2, height-margin+25,     // Lib Right label 2
		userX, userY,                                     // User position circle
		userX, userY+1,                                   // User position text
		margin, height-30,                                // Position text
		economicScore, socialScore,                       // Position coordinates
	)

	return svg
}
