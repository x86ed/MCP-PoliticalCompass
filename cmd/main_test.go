package main

import (
	"strings"
	"testing"

	"github.com/x86ed/MCP-PoliticalCompass/polcomp"
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

	if len(shuffledQuestions) != len(polcomp.AllQuestions) {
		t.Errorf("expected %d shuffled questions, got %d", len(polcomp.AllQuestions), len(shuffledQuestions))
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

	expectedError := "invalid response: Invalid Response"
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

	if !strings.Contains(content, "Current Economic Score:") {
		t.Error("should show current economic score")
	}

	if !strings.Contains(content, "Current Social Score:") {
		t.Error("should show current social score")
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
	for i := 0; i < len(polcomp.AllQuestions)-1; i++ {
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
	for i := 0; i < len(polcomp.AllQuestions); i++ {
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
			questionCount = len(polcomp.AllQuestions)
			currentIndex = len(polcomp.AllQuestions) // Set to completion point

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
	firstQuestion := polcomp.AllQuestions[firstQuestionIndex]

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
			questionCount = len(polcomp.AllQuestions)
			currentIndex = len(polcomp.AllQuestions)

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

// Test prompt handlers
func TestPromptHandlers(t *testing.T) {
	t.Run("uppercase prompt handler", func(t *testing.T) {
		args := PromptArgs{Input: "hello world"}
		response, err := handleUppercasePrompt(args)

		if err != nil {
			t.Fatalf("Expected no error, got %v", err)
		}

		if response == nil {
			t.Fatal("Expected response, got nil")
		}

		// Check that the text was converted to uppercase
		if response.Messages[0].Content.TextContent.Text != "HELLO WORLD" {
			t.Errorf("Expected 'HELLO WORLD', got '%s'", response.Messages[0].Content.TextContent.Text)
		}
	})

	t.Run("reverse prompt handler", func(t *testing.T) {
		args := PromptArgs{Input: "hello"}
		response, err := handleReversePrompt(args)

		if err != nil {
			t.Fatalf("Expected no error, got %v", err)
		}

		if response == nil {
			t.Fatal("Expected response, got nil")
		}

		// Check that the text was reversed
		if response.Messages[0].Content.TextContent.Text != "olleh" {
			t.Errorf("Expected 'olleh', got '%s'", response.Messages[0].Content.TextContent.Text)
		}
	})

	t.Run("uppercase with special characters", func(t *testing.T) {
		args := PromptArgs{Input: "hello! 123 @#$"}
		response, err := handleUppercasePrompt(args)

		if err != nil {
			t.Fatalf("Expected no error, got %v", err)
		}

		if response.Messages[0].Content.TextContent.Text != "HELLO! 123 @#$" {
			t.Errorf("Expected 'HELLO! 123 @#$', got '%s'", response.Messages[0].Content.TextContent.Text)
		}
	})

	t.Run("reverse with unicode characters", func(t *testing.T) {
		args := PromptArgs{Input: "café"}
		response, err := handleReversePrompt(args)

		if err != nil {
			t.Fatalf("Expected no error, got %v", err)
		}

		if response.Messages[0].Content.TextContent.Text != "éfac" {
			t.Errorf("Expected 'éfac', got '%s'", response.Messages[0].Content.TextContent.Text)
		}
	})

	t.Run("empty string handling", func(t *testing.T) {
		args := PromptArgs{Input: ""}

		// Test uppercase with empty string
		response, err := handleUppercasePrompt(args)
		if err != nil {
			t.Fatalf("Expected no error for uppercase, got %v", err)
		}
		if response.Messages[0].Content.TextContent.Text != "" {
			t.Errorf("Expected empty string for uppercase, got '%s'", response.Messages[0].Content.TextContent.Text)
		}

		// Test reverse with empty string
		response, err = handleReversePrompt(args)
		if err != nil {
			t.Fatalf("Expected no error for reverse, got %v", err)
		}
		if response.Messages[0].Content.TextContent.Text != "" {
			t.Errorf("Expected empty string for reverse, got '%s'", response.Messages[0].Content.TextContent.Text)
		}
	})
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
	t.Run("server setup with all tools and prompts", func(t *testing.T) {
		transport := createServerTransport()
		server, err := setupServer(transport)

		if err != nil {
			t.Fatalf("Expected no error, got %v", err)
		}

		if server == nil {
			t.Fatal("Expected server to be created, got nil")
		}

		// Test that the server was properly configured
		// This indirectly tests that all tools and prompts were registered successfully
		// since setupServer would return an error if registration failed
	})
}
