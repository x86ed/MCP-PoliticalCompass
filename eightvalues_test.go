package main

import (
	"fmt"
	"strings"
	"testing"

	"github.com/x86ed/MCP-PoliticalCompass/v2/eightvalues"
)

func TestEightValuesToolStart(t *testing.T) {
	// Reset state first
	resetState()

	args := EightValuesArgs{
		Response: "",
	}

	response, err := handleEightValues(args)
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	content := response.Content[0]
	if content.TextContent == nil {
		t.Fatal("response content is not TextContent")
	}

	text := content.TextContent.Text

	if !strings.Contains(text, "8values Political Quiz Started!") {
		t.Errorf("Expected quiz start message, got: %s", text)
	}

	if !strings.Contains(text, "Question 1 of 70") {
		t.Errorf("Expected question count, got: %s", text)
	}

	if !strings.Contains(text, "strongly_disagree, disagree, neutral, agree, or strongly_agree") {
		t.Errorf("Expected response options, got: %s", text)
	}
}

func TestEightValuesInvalidResponse(t *testing.T) {
	// Reset state and start quiz
	resetState()

	// First call to start quiz
	startArgs := EightValuesArgs{Response: ""}
	_, err := handleEightValues(startArgs)
	if err != nil {
		t.Fatalf("Expected no error starting quiz, got %v", err)
	}

	// Test invalid response
	invalidArgs := EightValuesArgs{
		Response: "invalid_response",
	}

	_, err = handleEightValues(invalidArgs)
	if err == nil {
		t.Errorf("Expected error for invalid response, got nil")
	}

	expectedError := "invalid response: invalid_response"
	if !strings.Contains(err.Error(), expectedError) {
		t.Errorf("Expected error containing '%s', got: %s", expectedError, err.Error())
	}
}

func TestEightValuesAllResponseTypes(t *testing.T) {
	responses := []string{"strongly_disagree", "disagree", "neutral", "agree", "strongly_agree"}

	for _, response := range responses {
		t.Run(response, func(t *testing.T) {
			// Reset state
			resetState()

			// Start quiz
			startArgs := EightValuesArgs{Response: ""}
			_, err := handleEightValues(startArgs)
			if err != nil {
				t.Fatalf("Expected no error starting quiz, got %v", err)
			}

			// Test response
			args := EightValuesArgs{Response: response}
			result, err := handleEightValues(args)
			if err != nil {
				t.Fatalf("Expected no error for response %s, got %v", response, err)
			}

			content := result.Content[0]
			if content.TextContent == nil {
				t.Fatal("response content is not TextContent")
			}

			text := content.TextContent.Text

			if !strings.Contains(text, "Response recorded!") {
				t.Errorf("Expected response recorded message for %s, got: %s", response, text)
			}
		})
	}
}

func TestResetEightValuesTool(t *testing.T) {
	// Start a quiz and answer some questions
	resetState()

	startArgs := EightValuesArgs{Response: ""}
	_, err := handleEightValues(startArgs)
	if err != nil {
		t.Fatalf("Expected no error starting quiz, got %v", err)
	}

	// Answer a few questions
	for i := 0; i < 3; i++ {
		args := EightValuesArgs{Response: "agree"}
		_, err := handleEightValues(args)
		if err != nil {
			t.Fatalf("Expected no error on question %d, got %v", i+1, err)
		}
	}

	// Reset the quiz
	resetArgs := ResetEightValuesArgs{}
	response, err := handleResetEightValues(resetArgs)
	if err != nil {
		t.Fatalf("Expected no error resetting quiz, got %v", err)
	}

	content := response.Content[0]
	if content.TextContent == nil {
		t.Fatal("response content is not TextContent")
	}

	text := content.TextContent.Text

	if !strings.Contains(text, "8values Quiz Reset!") {
		t.Errorf("Expected reset message, got: %s", text)
	}

	// Verify state is reset
	if eightValuesQuestionCount != 0 {
		t.Errorf("Expected question count to be 0 after reset, got %d", eightValuesQuestionCount)
	}

	if len(eightValuesQuizState.Responses) != 0 {
		t.Errorf("Expected no responses after reset, got %d", len(eightValuesQuizState.Responses))
	}
}

func TestEightValuesStatusTool(t *testing.T) {
	// Reset state
	resetState()

	// Test status with no progress
	statusArgs := EightValuesStatusArgs{}
	response, err := handleEightValuesStatus(statusArgs)
	if err != nil {
		t.Fatalf("Expected no error getting status, got %v", err)
	}

	content := response.Content[0]
	if content.TextContent == nil {
		t.Fatal("response content is not TextContent")
	}

	text := content.TextContent.Text

	if !strings.Contains(text, "8values Quiz Status") {
		t.Errorf("Expected status title, got: %s", text)
	}

	if !strings.Contains(text, "Questions answered: 0/70") {
		t.Errorf("Expected initial progress, got: %s", text)
	}

	// Start quiz and answer some questions
	startArgs := EightValuesArgs{Response: ""}
	_, err = handleEightValues(startArgs)
	if err != nil {
		t.Fatalf("Expected no error starting quiz, got %v", err)
	}

	// Answer 3 questions
	for i := 0; i < 3; i++ {
		args := EightValuesArgs{Response: "agree"}
		_, err := handleEightValues(args)
		if err != nil {
			t.Fatalf("Expected no error on question %d, got %v", i+1, err)
		}
	}

	// Check status again
	response, err = handleEightValuesStatus(statusArgs)
	if err != nil {
		t.Fatalf("Expected no error getting status, got %v", err)
	}

	content = response.Content[0]
	if content.TextContent == nil {
		t.Fatal("response content is not TextContent")
	}

	text = content.TextContent.Text

	if !strings.Contains(text, "Questions answered: 3/70") {
		t.Errorf("Expected progress after 3 questions, got: %s", text)
	}

	if !strings.Contains(text, "Response Distribution:") {
		t.Errorf("Expected response distribution, got: %s", text)
	}
}

func TestEightValuesSVGGeneration(t *testing.T) {
	testCases := []struct {
		name         string
		econPercent  float64
		diplPercent  float64
		govtPercent  float64
		sctyPercent  float64
		expectedTags []string
	}{
		{
			name:         "Center",
			econPercent:  50.0,
			diplPercent:  50.0,
			govtPercent:  50.0,
			sctyPercent:  50.0,
			expectedTags: []string{"<svg", "Economic Axis", "Diplomatic Axis", "Government Axis", "Society Axis"},
		},
		{
			name:         "Socialist Progressive",
			econPercent:  75.0,
			diplPercent:  60.0,
			govtPercent:  70.0,
			sctyPercent:  80.0,
			expectedTags: []string{"<svg", "75.0%", "60.0%", "70.0%", "80.0%"},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			svg := generateEightValuesSVG(tc.econPercent, tc.diplPercent, tc.govtPercent, tc.sctyPercent)

			for _, expectedTag := range tc.expectedTags {
				if !strings.Contains(svg, expectedTag) {
					t.Errorf("Expected SVG to contain '%s', but it didn't. SVG: %s", expectedTag, svg)
				}
			}

			// Verify it's valid XML structure
			if !strings.HasPrefix(svg, "<svg") {
				t.Errorf("Expected SVG to start with '<svg', got: %s", svg[:10])
			}

			if !strings.HasSuffix(svg, "</svg>") {
				t.Errorf("Expected SVG to end with '</svg>', got: %s", svg[len(svg)-10:])
			}
		})
	}
}

func TestEightValuesResponseMapping(t *testing.T) {
	// Test that response strings map to correct multiplier values
	testCases := []struct {
		response string
		expected float64
	}{
		{"strongly_disagree", eightvalues.StronglyDisagree},
		{"disagree", eightvalues.Disagree},
		{"neutral", eightvalues.Neutral},
		{"agree", eightvalues.Agree},
		{"strongly_agree", eightvalues.StronglyAgree},
	}

	for _, tc := range testCases {
		t.Run(tc.response, func(t *testing.T) {
			// Reset state for clean test
			resetState()

			// Start quiz
			startArgs := EightValuesArgs{Response: ""}
			_, err := handleEightValues(startArgs)
			if err != nil {
				t.Fatalf("Expected no error starting quiz, got %v", err)
			}

			args := EightValuesArgs{Response: tc.response}
			_, err = handleEightValues(args)
			if err != nil {
				t.Fatalf("Expected no error for response %s, got %v", tc.response, err)
			}

			// Check that the response was recorded with the correct value
			if len(eightValuesQuizState.Responses) != 1 {
				t.Fatalf("Expected 1 response recorded, got %d", len(eightValuesQuizState.Responses))
			}

			if eightValuesQuizState.Responses[0] != tc.expected {
				t.Errorf("Expected response value %f, got %f", tc.expected, eightValuesQuizState.Responses[0])
			}
		})
	}
}

// TestGetQuadrant tests all quadrant possibilities to improve coverage
func TestGetQuadrant(t *testing.T) {
	testCases := []struct {
		name     string
		economic float64
		social   float64
		expected string
	}{
		{"Libertarian Right", 5.0, 3.0, "Libertarian Right"},
		{"Libertarian Left", -3.0, 2.0, "Libertarian Left"},
		{"Authoritarian Right", 4.0, -2.0, "Authoritarian Right"},
		{"Authoritarian Left", -1.0, -4.0, "Authoritarian Left"},
		{"Center Right Libertarian", 0.1, 0.1, "Libertarian Right"},
		{"Center Left Libertarian", -0.1, 0.1, "Libertarian Left"},
		{"Center Right Authoritarian", 0.1, -0.1, "Authoritarian Right"},
		{"Exact Center", 0.0, 0.0, "Authoritarian Left"}, // 0,0 falls to else case
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := getQuadrant(tc.economic, tc.social)
			if result != tc.expected {
				t.Errorf("getQuadrant(%f, %f) = %s, expected %s", tc.economic, tc.social, result, tc.expected)
			}
		})
	}
}

// TestGeneratePoliticalCompassSVGBoundaries tests edge cases to improve coverage
func TestGeneratePoliticalCompassSVGBoundaries(t *testing.T) {
	testCases := []struct {
		name      string
		economic  float64
		social    float64
		testCheck func(svg string) bool
	}{
		{
			name:     "Extreme Left Economic",
			economic: -15.0, // Beyond normal range to test clamping
			social:   0.0,
			testCheck: func(svg string) bool {
				return strings.Contains(svg, "Position: (-15.00, 0.00)")
			},
		},
		{
			name:     "Extreme Right Economic",
			economic: 15.0, // Beyond normal range to test clamping
			social:   0.0,
			testCheck: func(svg string) bool {
				return strings.Contains(svg, "Position: (15.00, 0.00)")
			},
		},
		{
			name:     "Extreme Libertarian Social",
			economic: 0.0,
			social:   15.0, // Beyond normal range to test clamping
			testCheck: func(svg string) bool {
				return strings.Contains(svg, "Position: (0.00, 15.00)")
			},
		},
		{
			name:     "Extreme Authoritarian Social",
			economic: 0.0,
			social:   -15.0, // Beyond normal range to test clamping
			testCheck: func(svg string) bool {
				return strings.Contains(svg, "Position: (0.00, -15.00)")
			},
		},
		{
			name:     "Normal Center Position",
			economic: 0.0,
			social:   0.0,
			testCheck: func(svg string) bool {
				return strings.Contains(svg, "Position: (0.00, 0.00)") &&
					strings.Contains(svg, "<circle cx=\"200\" cy=\"200\"") // Should be center
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			svg := generatePoliticalCompassSVG(tc.economic, tc.social)

			// Basic SVG structure checks
			if !strings.HasPrefix(svg, "<svg") {
				t.Errorf("Expected SVG to start with '<svg', got: %s", svg[:10])
			}
			if !strings.HasSuffix(svg, "</svg>") {
				t.Errorf("Expected SVG to end with '</svg>'")
			}

			// Custom test check
			if !tc.testCheck(svg) {
				t.Errorf("Custom test check failed for %s", tc.name)
			}
		})
	}
}

// TestEightValuesCompleteQuiz tests completing the entire quiz to improve coverage
func TestEightValuesCompleteQuiz(t *testing.T) {
	// Reset state
	resetState()

	// Start quiz
	startArgs := EightValuesArgs{Response: ""}
	response, err := handleEightValues(startArgs)
	if err != nil {
		t.Fatalf("Expected no error starting quiz, got %v", err)
	}

	// Answer all questions to complete the quiz
	totalQuestions := len(eightvalues.Questions)
	responses := []string{"strongly_agree", "agree", "neutral", "disagree", "strongly_disagree"}

	for i := 0; i < totalQuestions; i++ {
		// Use different responses to test scoring
		responseType := responses[i%len(responses)]
		args := EightValuesArgs{Response: responseType}
		response, err = handleEightValues(args)
		if err != nil {
			t.Fatalf("Expected no error on question %d, got %v", i+1, err)
		}
	}

	// Verify completion
	content := response.Content[0]
	if content.TextContent == nil {
		t.Fatal("response content is not TextContent")
	}

	text := content.TextContent.Text

	if !strings.Contains(text, "8values Political Quiz Complete!") {
		t.Errorf("Expected completion message, got: %s", text)
	}

	if !strings.Contains(text, "Questions answered: 70") {
		t.Errorf("Expected final question count, got: %s", text)
	}

	// Should contain all four axis results
	expectedLabels := []string{"Economic Axis:", "Diplomatic Axis:", "Government Axis:", "Society Axis:"}
	for _, label := range expectedLabels {
		if !strings.Contains(text, label) {
			t.Errorf("Expected label '%s' in final results, got: %s", label, text)
		}
	}

	// Should contain SVG
	if !strings.Contains(text, "<svg") {
		t.Errorf("Expected SVG in final results, got: %s", text)
	}

	// Should contain ideological classifications
	ideologyLabels := []string{"Socialist", "Capitalist", "Internationalist", "Nationalist",
		"Libertarian", "Authoritarian", "Progressive", "Traditional"}
	foundIdeology := false
	for _, label := range ideologyLabels {
		if strings.Contains(text, label) {
			foundIdeology = true
			break
		}
	}
	if !foundIdeology {
		t.Errorf("Expected at least one ideological classification in results")
	}
}

// TestEightValuesStatusWithCompleteQuiz tests status after quiz completion
func TestEightValuesStatusWithCompleteQuiz(t *testing.T) {
	// Reset state
	resetState()

	// Complete the entire quiz first
	startArgs := EightValuesArgs{Response: ""}
	_, err := handleEightValues(startArgs)
	if err != nil {
		t.Fatalf("Expected no error starting quiz, got %v", err)
	}

	// Answer all questions
	totalQuestions := len(eightvalues.Questions)
	for i := 0; i < totalQuestions; i++ {
		args := EightValuesArgs{Response: "agree"}
		_, err = handleEightValues(args)
		if err != nil {
			t.Fatalf("Expected no error on question %d, got %v", i+1, err)
		}
	}

	// Now test status
	statusArgs := EightValuesStatusArgs{}
	response, err := handleEightValuesStatus(statusArgs)
	if err != nil {
		t.Fatalf("Expected no error getting status, got %v", err)
	}

	content := response.Content[0]
	if content.TextContent == nil {
		t.Fatal("response content is not TextContent")
	}

	text := content.TextContent.Text

	// Should show completion
	if !strings.Contains(text, "Questions answered: 70/70") {
		t.Errorf("Expected full completion in status, got: %s", text)
	}

	if !strings.Contains(text, "Questions remaining: 0") {
		t.Errorf("Expected no remaining questions, got: %s", text)
	}

	if !strings.Contains(text, "Completion: 100.0%") {
		t.Errorf("Expected 100%% completion, got: %s", text)
	}

	// Should show final scores and classifications
	if !strings.Contains(text, "**Final Scores:**") {
		t.Errorf("Expected final scores section, got: %s", text)
	}

	// Should contain axis percentages
	axisLabels := []string{"Economic Axis:", "Diplomatic Axis:", "Government Axis:", "Society Axis:"}
	for _, label := range axisLabels {
		if !strings.Contains(text, label) {
			t.Errorf("Expected axis label '%s' in status, got: %s", label, text)
		}
	}
}

// TestEightValuesEdgeCases tests edge cases to improve coverage
func TestEightValuesEdgeCases(t *testing.T) {
	t.Run("EmptyResponseAfterStart", func(t *testing.T) {
		resetState()

		// Start quiz
		startArgs := EightValuesArgs{Response: ""}
		_, err := handleEightValues(startArgs)
		if err != nil {
			t.Fatalf("Expected no error starting quiz, got %v", err)
		}

		// Try empty response
		emptyArgs := EightValuesArgs{Response: ""}
		_, err = handleEightValues(emptyArgs)
		if err == nil {
			t.Errorf("Expected error for empty response after start, got nil")
		}
	})

	t.Run("StatusWithPartialProgress", func(t *testing.T) {
		resetState()

		// Start quiz and answer exactly half the questions
		startArgs := EightValuesArgs{Response: ""}
		_, err := handleEightValues(startArgs)
		if err != nil {
			t.Fatalf("Expected no error starting quiz, got %v", err)
		}

		totalQuestions := len(eightvalues.Questions)
		halfQuestions := totalQuestions / 2

		for i := 0; i < halfQuestions; i++ {
			args := EightValuesArgs{Response: "neutral"}
			_, err = handleEightValues(args)
			if err != nil {
				t.Fatalf("Expected no error on question %d, got %v", i+1, err)
			}
		}

		// Check status
		statusArgs := EightValuesStatusArgs{}
		response, err := handleEightValuesStatus(statusArgs)
		if err != nil {
			t.Fatalf("Expected no error getting status, got %v", err)
		}

		content := response.Content[0]
		text := content.TextContent.Text

		// Should show partial progress
		expectedAnswered := fmt.Sprintf("Questions answered: %d/70", halfQuestions)
		if !strings.Contains(text, expectedAnswered) {
			t.Errorf("Expected '%s' in status, got: %s", expectedAnswered, text)
		}

		expectedRemaining := fmt.Sprintf("Questions remaining: %d", totalQuestions-halfQuestions)
		if !strings.Contains(text, expectedRemaining) {
			t.Errorf("Expected '%s' in status, got: %s", expectedRemaining, text)
		}
	})
}
