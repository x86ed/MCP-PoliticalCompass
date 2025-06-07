package main

import (
	"strings"
	"testing"

	politicalcompass "github.com/x86ed/MCP-PoliticalCompass/v2/political-compass"
)

func TestPoliticalCompassToolStart(t *testing.T) {
	resetState()

	response, err := handlePoliticalCompass(PoliticalCompassArgs{Response: ""})
	if err != nil {
		t.Fatalf("unexpected error starting quiz: %v", err)
	}

	if response == nil {
		t.Fatal("response is nil")
	}

	if len(response.Content) == 0 {
		t.Fatal("response content is empty")
	}

	content := response.Content[0]
	if content.TextContent == nil {
		t.Fatal("response content is not TextContent")
	}

	responseText := content.TextContent.Text
	if !strings.Contains(responseText, "Political Compass Quiz Started!") {
		t.Error("start response should contain quiz started message")
	}

	if questionCount != 1 {
		t.Errorf("expected questionCount to be 1, got %d", questionCount)
	}

	if len(shuffledQuestions) != len(politicalcompass.AllQuestions) {
		t.Errorf("expected %d shuffled questions, got %d", len(politicalcompass.AllQuestions), len(shuffledQuestions))
	}
}

func TestPoliticalCompassInvalidResponse(t *testing.T) {
	resetState()

	// Start quiz first
	handlePoliticalCompass(PoliticalCompassArgs{Response: ""})

	// Try invalid response
	_, err := handlePoliticalCompass(PoliticalCompassArgs{Response: "Invalid Response"})
	if err == nil {
		t.Error("expected error for invalid response")
	}

	expectedError := "invalid response: Invalid Response. Please use one of: strongly_disagree, disagree, agree, strongly_agree"
	if err.Error() != expectedError {
		t.Errorf("expected error message '%s', got '%s'", expectedError, err.Error())
	}
}

func TestPoliticalCompassAllResponseTypes(t *testing.T) {
	resetState()

	responses := []string{"Strongly Disagree", "Disagree", "Agree", "Strongly Agree"}

	for _, resp := range responses {
		t.Run(resp, func(t *testing.T) {
			resetState()

			// Start quiz
			handlePoliticalCompass(PoliticalCompassArgs{Response: ""})

			// Test response
			response, err := handlePoliticalCompass(PoliticalCompassArgs{Response: resp})
			if err != nil {
				t.Fatalf("unexpected error for response '%s': %v", resp, err)
			}

			if response == nil {
				t.Fatal("response is nil")
			}

			content := response.Content[0].TextContent.Text
			if !strings.Contains(content, "Response recorded!") {
				t.Errorf("response should contain 'Response recorded!' for '%s'", resp)
			}
		})
	}
}

func TestPoliticalCompassProgressDisplay(t *testing.T) {
	resetState()

	// Start quiz
	handlePoliticalCompass(PoliticalCompassArgs{Response: ""})

	// Answer first question
	response, err := handlePoliticalCompass(PoliticalCompassArgs{Response: "Agree"})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	content := response.Content[0].TextContent.Text

	// Should show progress
	if !strings.Contains(content, "Response recorded!") {
		t.Error("should show response recorded message")
	}

	if !strings.Contains(content, "Progress: 1 of") {
		t.Error("should show progress information")
	}

	// Should NOT show intermediate scores (only final results)
	if strings.Contains(content, "Current Economic Score:") {
		t.Error("should NOT show current economic score during quiz")
	}

	if strings.Contains(content, "Current Social Score:") {
		t.Error("should NOT show current social score during quiz")
	}

	if !strings.Contains(content, "Question 2 of") {
		t.Error("should show next question number")
	}
}

func TestPoliticalCompassQuizCompletion(t *testing.T) {
	resetState()

	// Simulate answering all questions
	handlePoliticalCompass(PoliticalCompassArgs{Response: ""}) // Start

	// Answer all questions except the last one
	for i := 0; i < len(politicalcompass.AllQuestions)-1; i++ {
		_, err := handlePoliticalCompass(PoliticalCompassArgs{Response: "Agree"})
		if err != nil {
			t.Fatalf("unexpected error on question %d: %v", i+1, err)
		}
	}

	// Answer the final question
	response, err := handlePoliticalCompass(PoliticalCompassArgs{Response: "Agree"})
	if err != nil {
		t.Fatalf("unexpected error on final question: %v", err)
	}

	content := response.Content[0].TextContent.Text

	// Check completion message
	if !strings.Contains(content, "Political Compass Quiz Complete!") {
		t.Error("should show quiz completion message")
	}

	if !strings.Contains(content, "Questions answered:") {
		t.Error("should show total questions answered")
	}

	if !strings.Contains(content, "Final Economic Score:") {
		t.Error("should show final economic score")
	}

	if !strings.Contains(content, "Final Social Score:") {
		t.Error("should show final social score")
	}

	if !strings.Contains(content, "Your Political Quadrant:") {
		t.Error("should show political quadrant")
	}
}

func TestPoliticalCompassCompletionWithSVG(t *testing.T) {
	resetState()

	// Simulate answering all questions
	handlePoliticalCompass(PoliticalCompassArgs{Response: ""}) // Start

	// Answer all questions
	for i := 0; i < len(politicalcompass.AllQuestions); i++ {
		_, err := handlePoliticalCompass(PoliticalCompassArgs{Response: "Agree"})
		if err != nil {
			t.Fatalf("unexpected error on question %d: %v", i+1, err)
		}
	}

	// Get the last response which should contain the SVG
	response, err := handlePoliticalCompass(PoliticalCompassArgs{Response: "Agree"})
	if err != nil {
		t.Fatalf("unexpected error on completion: %v", err)
	}

	content := response.Content[0].TextContent.Text

	// Check that SVG is included
	if !strings.Contains(content, "<svg") {
		t.Error("completion response should contain SVG")
	}

	if !strings.Contains(content, "</svg>") {
		t.Error("completion response should contain closing SVG tag")
	}

	if !strings.Contains(content, "xmlns=\"http://www.w3.org/2000/svg\"") {
		t.Error("SVG should have proper namespace")
	}

	if !strings.Contains(content, "Political Compass Quiz Complete!") {
		t.Error("should show quiz completion message")
	}

	// Check that position coordinates are included in SVG
	if !strings.Contains(content, "Position:") {
		t.Error("SVG should contain position coordinates")
	}
}

func TestPoliticalCompassQuadrantCalculations(t *testing.T) {
	testCases := []struct {
		name          string
		economicScore float64
		socialScore   float64
		expectedQuad  string
	}{
		{"Libertarian Left", 1.0, 1.0, "Libertarian Left"},
		{"Authoritarian Left", 1.0, -1.0, "Authoritarian Left"},
		{"Libertarian Right", -1.0, 1.0, "Libertarian Right"},
		{"Authoritarian Right", -1.0, -1.0, "Authoritarian Right"},
		{"Center-Right Authoritarian", -0.1, -0.1, "Authoritarian Right"},
		{"Center-Left Libertarian", 0.1, 0.1, "Libertarian Left"},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			resetState()

			// Start the quiz first
			handlePoliticalCompass(PoliticalCompassArgs{Response: ""})

			// Manually set scores to simulate reaching the desired final scores
			totalEconomicScore = (tc.economicScore - 0.38) * 8.0
			totalSocialScore = (tc.socialScore - 2.41) * 19.5
			questionCount = len(politicalcompass.AllQuestions)
			currentIndex = len(politicalcompass.AllQuestions) // Set to completion point

			// Call with a valid response to trigger completion logic
			response, err := handlePoliticalCompass(PoliticalCompassArgs{Response: "Agree"})
			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}

			content := response.Content[0].TextContent.Text
			if !strings.Contains(content, tc.expectedQuad) {
				t.Errorf("expected quadrant '%s' but content was: %s", tc.expectedQuad, content)
			}
		})
	}
}

func TestPoliticalCompassScoreAccumulation(t *testing.T) {
	resetState()

	// Start quiz
	handlePoliticalCompass(PoliticalCompassArgs{Response: ""})

	initialEconomic := totalEconomicScore
	initialSocial := totalSocialScore

	// Answer with "Agree" which should affect scores
	handlePoliticalCompass(PoliticalCompassArgs{Response: "Agree"})

	// Scores should have changed (unless the first question has all zeros, which is unlikely)
	if totalEconomicScore == initialEconomic && totalSocialScore == initialSocial {
		// This could happen if the first question has zero scores for "Agree"
		// Let's check a few more questions to ensure score accumulation works
		handlePoliticalCompass(PoliticalCompassArgs{Response: "Strongly Agree"})
		handlePoliticalCompass(PoliticalCompassArgs{Response: "Disagree"})
		handlePoliticalCompass(PoliticalCompassArgs{Response: "Strongly Disagree"})
	}

	// After answering several questions, at least one should have affected the scores
	if questionCount > 1 && totalEconomicScore == 0.0 && totalSocialScore == 0.0 {
		t.Error("expected some score accumulation after answering questions")
	}
}

func TestResetQuizTool(t *testing.T) {
	resetState()
	handlePoliticalCompass(PoliticalCompassArgs{Response: ""})
	handlePoliticalCompass(PoliticalCompassArgs{Response: "Agree"})

	if questionCount == 0 {
		t.Fatal("quiz should have started")
	}

	response, err := handleResetQuiz(ResetQuizArgs{})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if response == nil {
		t.Fatal("response is nil")
	}

	if questionCount != 0 || totalEconomicScore != 0.0 || totalSocialScore != 0.0 || currentIndex != 0 {
		t.Error("quiz state was not properly reset")
	}

	if len(shuffledQuestions) != 0 {
		t.Error("shuffled questions should be empty after reset")
	}
}

func TestPoliticalCompassEdgeCases(t *testing.T) {
	t.Run("Empty response on first question", func(t *testing.T) {
		resetState()

		response, err := handlePoliticalCompass(PoliticalCompassArgs{Response: ""})
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}

		content := response.Content[0].TextContent.Text
		if !strings.Contains(content, "Political Compass Quiz Started!") {
			t.Error("should start quiz with empty response")
		}
	})

	t.Run("Shuffled questions are different each time", func(t *testing.T) {
		resetState()
		handlePoliticalCompass(PoliticalCompassArgs{Response: ""})
		firstShuffle := make([]int, len(shuffledQuestions))
		copy(firstShuffle, shuffledQuestions)

		resetState()
		handlePoliticalCompass(PoliticalCompassArgs{Response: ""})
		secondShuffle := make([]int, len(shuffledQuestions))
		copy(secondShuffle, shuffledQuestions)

		// While technically they could be the same due to randomization,
		// the probability is extremely low with 62 questions
		allSame := true
		for i := range firstShuffle {
			if firstShuffle[i] != secondShuffle[i] {
				allSame = false
				break
			}
		}

		if allSame && len(firstShuffle) > 10 {
			t.Log("Warning: Shuffled questions were identical - this is extremely unlikely but possible")
		}
	})

	t.Run("Initialization only happens once", func(t *testing.T) {
		resetState()

		// Call multiple times
		handlePoliticalCompass(PoliticalCompassArgs{Response: ""})
		firstShuffle := make([]int, len(shuffledQuestions))
		copy(firstShuffle, shuffledQuestions)

		// Call again without reset - should use same shuffled order
		handlePoliticalCompass(PoliticalCompassArgs{Response: "Agree"})

		// Shuffled questions should be unchanged
		for i := range firstShuffle {
			if firstShuffle[i] != shuffledQuestions[i] {
				t.Error("shuffled questions should not change after initialization")
				break
			}
		}
	})
}

func TestPoliticalCompassScoreCalculationDetails(t *testing.T) {
	resetState()

	// Test that scores are calculated correctly
	handlePoliticalCompass(PoliticalCompassArgs{Response: ""})

	// Get the first question
	firstQuestionIndex := shuffledQuestions[0]
	firstQuestion := politicalcompass.AllQuestions[firstQuestionIndex]

	// Answer with "Agree" (index 2)
	expectedEconomic := firstQuestion.Economic[2]
	expectedSocial := firstQuestion.Social[2]

	handlePoliticalCompass(PoliticalCompassArgs{Response: "Agree"})

	if totalEconomicScore != expectedEconomic {
		t.Errorf("expected economic score %f, got %f", expectedEconomic, totalEconomicScore)
	}

	if totalSocialScore != expectedSocial {
		t.Errorf("expected social score %f, got %f", expectedSocial, totalSocialScore)
	}
}

func TestPoliticalCompassBoundaryScores(t *testing.T) {
	resetState()

	// Test extreme scores that result in boundary quadrant calculations
	testCases := []struct {
		name         string
		economic     float64
		social       float64
		expectedQuad string
	}{
		{"Clear Libertarian Left", 0.5, 0.5, "Libertarian Left"},
		{"Clear Authoritarian Left", 0.5, -0.5, "Authoritarian Left"},
		{"Clear Libertarian Right", -0.5, 0.5, "Libertarian Right"},
		{"Clear Authoritarian Right", -0.5, -0.5, "Authoritarian Right"},
		{"Exactly zero both", 0.0, 0.0, "Authoritarian Right"},
		{"Small negative values", -0.1, -0.1, "Authoritarian Right"},
		{"Small positive values", 0.1, 0.1, "Libertarian Left"},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			resetState()
			handlePoliticalCompass(PoliticalCompassArgs{Response: ""})

			// Set up for completion with specific scores
			totalEconomicScore = (tc.economic - 0.38) * 8.0
			totalSocialScore = (tc.social - 2.41) * 19.5
			questionCount = len(politicalcompass.AllQuestions)
			currentIndex = len(politicalcompass.AllQuestions)

			response, err := handlePoliticalCompass(PoliticalCompassArgs{Response: "Agree"})
			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}

			content := response.Content[0].TextContent.Text
			if !strings.Contains(content, tc.expectedQuad) {
				t.Errorf("expected quadrant '%s', content: %s", tc.expectedQuad, content)
			}
		})
	}
}

// Test main.go functions
func TestMainFunctions(t *testing.T) {
	t.Run("createServerTransport", func(t *testing.T) {
		transport := createServerTransport()
		if transport == nil {
			t.Fatal("Expected transport to be created, got nil")
		}
	})

	t.Run("setupServer", func(t *testing.T) {
		transport := createServerTransport()
		server, err := setupServer(transport)

		if err != nil {
			t.Fatalf("Expected no error setting up server, got %v", err)
		}

		if server == nil {
			t.Fatal("Expected server to be created, got nil")
		}
	})
}

// Test server setup with mock transport
func TestServerSetupIntegration(t *testing.T) {
	t.Run("server setup with all tools", func(t *testing.T) {
		transport := createServerTransport()
		server, err := setupServer(transport)

		if err != nil {
			t.Fatalf("Expected no error, got %v", err)
		}

		if server == nil {
			t.Fatal("Expected server to be created, got nil")
		}

		// Test that the server was properly configured
		// This indirectly tests that all tools were registered successfully
		// since setupServer would return an error if registration failed
	})
}

// TestDetailedOutputValidation validates the complete output format and checks for any text errors
func TestDetailedOutputValidation(t *testing.T) {
	resetState()

	// Start the quiz
	handlePoliticalCompass(PoliticalCompassArgs{Response: ""})

	// Set up for a specific completion scenario - Libertarian Left
	totalEconomicScore = (1.5 - 0.38) * 8.0 // Economic score: +1.5
	totalSocialScore = (1.2 - 2.41) * 19.5  // Social score: +1.2
	questionCount = len(politicalcompass.AllQuestions)
	currentIndex = len(politicalcompass.AllQuestions)

	response, err := handlePoliticalCompass(PoliticalCompassArgs{Response: "Agree"})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	content := response.Content[0].TextContent.Text

	// Print the output for manual inspection
	t.Logf("Complete Quiz Output:\n%s", content)

	// Validate that text explanations are correct
	if !strings.Contains(content, "Left: + | Right: -") {
		t.Error("Economic score explanation is incorrect or missing")
	}

	if !strings.Contains(content, "Libertarian: + | Authoritarian: -") {
		t.Error("Social score explanation is incorrect or missing")
	}

	if !strings.Contains(content, "Your Political Quadrant: Libertarian Left") {
		t.Error("Expected Libertarian Left quadrant for positive economic and social scores")
	}

	// Check that the scores are displayed correctly
	if !strings.Contains(content, "Final Economic Score: 1.50") {
		t.Error("Economic score not displayed correctly")
	}

	if !strings.Contains(content, "Final Social Score: 1.20") {
		t.Error("Social score not displayed correctly")
	}

	// Validate SVG is present and well-formed
	if !strings.Contains(content, "<svg width=\"400\" height=\"400\"") {
		t.Error("SVG dimensions not correct")
	}

	if !strings.Contains(content, "xmlns=\"http://www.w3.org/2000/svg\"") {
		t.Error("SVG namespace missing")
	}

	if !strings.Contains(content, "Position: (1.50, 1.20)") {
		t.Error("Position coordinates not displayed correctly in SVG")
	}
}

// TestQuizStatusTool tests the quiz status functionality
func TestQuizStatusTool(t *testing.T) {
	resetState()

	// Test empty status
	response, err := handleQuizStatus(QuizStatusArgs{})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	content := response.Content[0].TextContent.Text
	if !strings.Contains(content, "📊 **Political Compass Quiz Status**") {
		t.Error("should show status header")
	}

	if !strings.Contains(content, "Questions answered: 0/62") {
		t.Error("should show zero progress initially")
	}

	// Start quiz and answer some questions
	handlePoliticalCompass(PoliticalCompassArgs{Response: ""})
	handlePoliticalCompass(PoliticalCompassArgs{Response: "Agree"})
	handlePoliticalCompass(PoliticalCompassArgs{Response: "Disagree"})

	// Check status after answering questions
	response, err = handleQuizStatus(QuizStatusArgs{})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	content = response.Content[0].TextContent.Text
	if !strings.Contains(content, "Questions answered: 2/62") {
		t.Error("should show correct progress after answering questions")
	}

	if !strings.Contains(content, "Response Distribution:") {
		t.Error("should show response distribution")
	}
}
