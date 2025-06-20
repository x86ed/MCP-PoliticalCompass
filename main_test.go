package main

import (
	"context"
	"fmt"
	"strings"
	"testing"

	politicalcompass "github.com/x86ed/MCP-PoliticalCompass/v3/political-compass"
)

func TestPoliticalCompassToolStart(t *testing.T) {
	resetState()

	response, err := handlePoliticalCompass(context.Background(), createMockRequest("political_compass", map[string]interface{}{"answer": ""}))
	if err != nil {
		t.Fatalf("unexpected error starting quiz: %v", err)
	}

	if response == nil {
		t.Fatal("response is nil")
	}

	if len(response.Content) == 0 {
		t.Fatal("response content is empty")
	}

	responseText := extractTextContent(response)
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
	handlePoliticalCompass(context.Background(), createMockRequest("political_compass", map[string]interface{}{"answer": ""}))

	// Try invalid response
	response, err := handlePoliticalCompass(context.Background(), createMockRequest("political_compass", map[string]interface{}{"answer": "Invalid Response"}))
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if !isErrorResult(response) {
		t.Error("expected error result for invalid response")
	}

	responseText := extractTextContent(response)
	expectedError := "invalid response: Invalid Response. Please use one of: strongly_disagree, disagree, agree, strongly_agree"
	if responseText != expectedError {
		t.Errorf("expected error message '%s', got '%s'", expectedError, responseText)
	}
}

func TestPoliticalCompassAllResponseTypes(t *testing.T) {
	resetState()

	responses := []string{"Strongly Disagree", "Disagree", "Agree", "Strongly Agree"}

	for _, resp := range responses {
		t.Run(resp, func(t *testing.T) {
			resetState()

			// Start quiz
			handlePoliticalCompass(context.Background(), createMockRequest("political_compass", map[string]interface{}{"answer": ""}))

			// Test response
			response, err := handlePoliticalCompass(context.Background(), createMockRequest("political_compass", map[string]interface{}{"answer": resp}))
			if err != nil {
				t.Fatalf("unexpected error for response '%s': %v", resp, err)
			}

			if response == nil {
				t.Fatal("response is nil")
			}

			content := extractTextContent(response)
			if !strings.Contains(content, "Response recorded!") {
				t.Errorf("response should contain 'Response recorded!' for '%s'", resp)
			}
		})
	}
}

func TestPoliticalCompassProgressDisplay(t *testing.T) {
	resetState()

	// Start quiz
	handlePoliticalCompass(context.Background(), createMockRequest("political_compass", map[string]interface{}{"answer": ""}))

	// Answer first question
	response, err := handlePoliticalCompass(context.Background(), createMockRequest("political_compass", map[string]interface{}{"answer": "Agree"}))
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	content := extractTextContent(response)

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
	handlePoliticalCompass(context.Background(), createMockRequest("political_compass", map[string]interface{}{"answer": ""})) // Start

	// Answer all questions except the last one
	for i := 0; i < len(politicalcompass.AllQuestions)-1; i++ {
		_, err := handlePoliticalCompass(context.Background(), createMockRequest("political_compass", map[string]interface{}{"answer": "Agree"}))
		if err != nil {
			t.Fatalf("unexpected error on question %d: %v", i+1, err)
		}
	}

	// Answer the final question
	response, err := handlePoliticalCompass(context.Background(), createMockRequest("political_compass", map[string]interface{}{"answer": "Agree"}))
	if err != nil {
		t.Fatalf("unexpected error on final question: %v", err)
	}

	content := extractTextContent(response)

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
	handlePoliticalCompass(context.Background(), createMockRequest("political_compass", map[string]interface{}{"answer": ""})) // Start

	// Answer all questions
	for i := 0; i < len(politicalcompass.AllQuestions); i++ {
		_, err := handlePoliticalCompass(context.Background(), createMockRequest("political_compass", map[string]interface{}{"answer": "Agree"}))
		if err != nil {
			t.Fatalf("unexpected error on question %d: %v", i+1, err)
		}
	}

	// Get the last response which should contain the SVG
	response, err := handlePoliticalCompass(context.Background(), createMockRequest("political_compass", map[string]interface{}{"answer": "Agree"}))
	if err != nil {
		t.Fatalf("unexpected error on completion: %v", err)
	}

	content := extractTextContent(response)

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
			handlePoliticalCompass(context.Background(), createMockRequest("political_compass", map[string]interface{}{"answer": ""}))

			// Manually set scores to simulate reaching the desired final scores
			totalEconomicScore = (tc.economicScore - 0.38) * 8.0
			totalSocialScore = (tc.socialScore - 2.41) * 19.5
			questionCount = len(politicalcompass.AllQuestions)
			currentIndex = len(politicalcompass.AllQuestions) // Set to completion point

			// Call with a valid response to trigger completion logic
			response, err := handlePoliticalCompass(context.Background(), createMockRequest("political_compass", map[string]interface{}{"answer": "Agree"}))
			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}

			content := extractTextContent(response)
			if !strings.Contains(content, tc.expectedQuad) {
				t.Errorf("expected quadrant '%s' but content was: %s", tc.expectedQuad, content)
			}
		})
	}
}

func TestPoliticalCompassScoreAccumulation(t *testing.T) {
	resetState()

	// Start quiz
	handlePoliticalCompass(context.Background(), createMockRequest("political_compass", map[string]interface{}{"answer": ""}))

	// Test basic score accumulation by checking that after answering enough questions,
	// at least one of the scores changes from 0.0
	initialEconomic := totalEconomicScore
	initialSocial := totalSocialScore

	// Answer a few questions with different responses
	responses := []string{"Strongly Agree", "Agree", "Disagree", "Strongly Disagree"}

	for i := 0; i < 4 && currentIndex < len(shuffledQuestions); i++ {
		response := responses[i%len(responses)]
		_, err := handlePoliticalCompass(context.Background(), createMockRequest("political_compass", map[string]interface{}{"answer": response}))
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
	}

	// After answering multiple questions, at least one score should have changed
	// This tests that the scoring mechanism is working
	if totalEconomicScore == initialEconomic && totalSocialScore == initialSocial {
		// Check if we can find any question with non-zero effects to verify it's not just bad luck
		hasNonZeroEffects := false
		for _, q := range politicalcompass.AllQuestions {
			for _, econ := range q.Economic {
				if econ != 0 {
					hasNonZeroEffects = true
					break
				}
			}
			if hasNonZeroEffects {
				break
			}
			for _, social := range q.Social {
				if social != 0 {
					hasNonZeroEffects = true
					break
				}
			}
			if hasNonZeroEffects {
				break
			}
		}

		if hasNonZeroEffects {
			t.Error("scores should have changed after answering questions, but both remained at initial values")
		} else {
			t.Skip("All questions have zero effects - data issue")
		}
	}

	// Verify that scores are reasonable (not extreme values that would indicate a bug)
	if totalEconomicScore < -1000 || totalEconomicScore > 1000 {
		t.Errorf("economic score %f is unreasonably extreme", totalEconomicScore)
	}

	if totalSocialScore < -1000 || totalSocialScore > 1000 {
		t.Errorf("social score %f is unreasonably extreme", totalSocialScore)
	}
}

func TestResetQuizTool(t *testing.T) {
	resetState()
	handlePoliticalCompass(context.Background(), createMockRequest("political_compass", map[string]interface{}{"answer": ""}))
	handlePoliticalCompass(context.Background(), createMockRequest("political_compass", map[string]interface{}{"answer": "Agree"}))

	if questionCount == 0 {
		t.Fatal("quiz should have started")
	}

	response, err := handleResetQuiz(context.Background(), createMockRequest("reset_quiz", map[string]interface{}{}))
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

		response, err := handlePoliticalCompass(context.Background(), createMockRequest("political_compass", map[string]interface{}{"answer": ""}))
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}

		content := extractTextContent(response)
		if !strings.Contains(content, "Political Compass Quiz Started!") {
			t.Error("should start quiz with empty response")
		}
	})

	t.Run("Shuffled questions are different each time", func(t *testing.T) {
		resetState()
		handlePoliticalCompass(context.Background(), createMockRequest("political_compass", map[string]interface{}{"answer": ""}))
		firstShuffle := make([]int, len(shuffledQuestions))
		copy(firstShuffle, shuffledQuestions)

		resetState()
		handlePoliticalCompass(context.Background(), createMockRequest("political_compass", map[string]interface{}{"answer": ""}))
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
		handlePoliticalCompass(context.Background(), createMockRequest("political_compass", map[string]interface{}{"answer": ""}))
		firstShuffle := make([]int, len(shuffledQuestions))
		copy(firstShuffle, shuffledQuestions)

		// Call again without reset - should use same shuffled order
		handlePoliticalCompass(context.Background(), createMockRequest("political_compass", map[string]interface{}{"answer": "Agree"}))

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
	handlePoliticalCompass(context.Background(), createMockRequest("political_compass", map[string]interface{}{"answer": ""}))

	// Get the first question
	firstQuestionIndex := shuffledQuestions[0]
	firstQuestion := politicalcompass.AllQuestions[firstQuestionIndex]

	// Answer with "Agree" (index 2)
	expectedEconomic := firstQuestion.Economic[2]
	expectedSocial := firstQuestion.Social[2]

	handlePoliticalCompass(context.Background(), createMockRequest("political_compass", map[string]interface{}{"answer": "Agree"}))

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
			handlePoliticalCompass(context.Background(), createMockRequest("political_compass", map[string]interface{}{"answer": ""}))

			// Set up for completion with specific scores
			totalEconomicScore = (tc.economic - 0.38) * 8.0
			totalSocialScore = (tc.social - 2.41) * 19.5
			questionCount = len(politicalcompass.AllQuestions)
			currentIndex = len(politicalcompass.AllQuestions)

			response, err := handlePoliticalCompass(context.Background(), createMockRequest("political_compass", map[string]interface{}{"answer": "Agree"}))
			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}

			content := extractTextContent(response)
			if !strings.Contains(content, tc.expectedQuad) {
				t.Errorf("expected quadrant '%s', content: %s", tc.expectedQuad, content)
			}
		})
	}
}

// TestDetailedOutputValidation validates the complete output format and checks for any text errors
func TestDetailedOutputValidation(t *testing.T) {
	resetState()

	// Start the quiz
	handlePoliticalCompass(context.Background(), createMockRequest("political_compass", map[string]interface{}{"answer": ""}))

	// Set up for a specific completion scenario - Libertarian Left
	totalEconomicScore = (1.5 - 0.38) * 8.0 // Economic score: +1.5
	totalSocialScore = (1.2 - 2.41) * 19.5  // Social score: +1.2
	questionCount = len(politicalcompass.AllQuestions)
	currentIndex = len(politicalcompass.AllQuestions)

	response, err := handlePoliticalCompass(context.Background(), createMockRequest("political_compass", map[string]interface{}{"answer": "Agree"}))
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	content := extractTextContent(response)

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
	response, err := handleQuizStatus(context.Background(), createMockRequest("quiz_status", map[string]interface{}{}))
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	content := extractTextContent(response)
	if !strings.Contains(content, "📊 **Political Compass Quiz Status**") {
		t.Error("should show status header")
	}

	if !strings.Contains(content, "Questions answered: 0/62") {
		t.Error("should show zero progress initially")
	}

	// Verify scores are not shown when no questions answered
	if strings.Contains(content, "Current Scores:") || strings.Contains(content, "Final Scores:") {
		t.Error("should not show scores when no questions answered")
	}
	if strings.Contains(content, "Current Quadrant:") || strings.Contains(content, "Your Quadrant:") {
		t.Error("should not show quadrant when no questions answered")
	}

	// Start quiz and answer some questions (but not all)
	handlePoliticalCompass(context.Background(), createMockRequest("political_compass", map[string]interface{}{"answer": ""}))
	handlePoliticalCompass(context.Background(), createMockRequest("political_compass", map[string]interface{}{"answer": "Agree"}))
	handlePoliticalCompass(context.Background(), createMockRequest("political_compass", map[string]interface{}{"answer": "Disagree"}))

	// Check status after answering questions (incomplete quiz)
	response, err = handleQuizStatus(context.Background(), createMockRequest("quiz_status", map[string]interface{}{}))
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	content = extractTextContent(response)
	if !strings.Contains(content, "Questions answered: 2/62") {
		t.Error("should show correct progress after answering questions")
	}

	if !strings.Contains(content, "Response Distribution:") {
		t.Error("should show response distribution")
	}

	// Verify scores and quadrant are NOT shown for incomplete quiz
	if strings.Contains(content, "Current Scores:") || strings.Contains(content, "Final Scores:") {
		t.Error("should not show scores for incomplete quiz")
	}
	if strings.Contains(content, "Current Quadrant:") || strings.Contains(content, "Your Quadrant:") {
		t.Error("should not show quadrant for incomplete quiz")
	}
	if strings.Contains(content, "Economic axis:") {
		t.Error("should not show economic axis scores for incomplete quiz")
	}
	if strings.Contains(content, "Social axis:") {
		t.Error("should not show social axis scores for incomplete quiz")
	}

	// Complete the entire quiz to test final status
	// Answer all remaining questions with "Agree"
	for i := 2; i < 62; i++ {
		handlePoliticalCompass(context.Background(), createMockRequest("political_compass", map[string]interface{}{"answer": "Agree"}))
	}

	// Check status after completing quiz
	response, err = handleQuizStatus(context.Background(), createMockRequest("quiz_status", map[string]interface{}{}))
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	content = extractTextContent(response)
	if !strings.Contains(content, "Questions answered: 62/62") {
		t.Error("should show complete progress")
	}

	if !strings.Contains(content, "Questions remaining: 0") {
		t.Error("should show zero remaining questions")
	}

	// Verify scores and quadrant ARE shown for complete quiz
	if !strings.Contains(content, "Final Scores:") {
		t.Error("should show final scores for complete quiz")
	}
	if !strings.Contains(content, "Your Quadrant:") {
		t.Error("should show quadrant for complete quiz")
	}
	if !strings.Contains(content, "Economic axis:") {
		t.Error("should show economic axis scores for complete quiz")
	}
	if !strings.Contains(content, "Social axis:") {
		t.Error("should show social axis scores for complete quiz")
	}

	// Should not show "Current" labels anymore
	if strings.Contains(content, "Current Scores:") {
		t.Error("should not show 'Current Scores' label for complete quiz")
	}
	if strings.Contains(content, "Current Quadrant:") {
		t.Error("should not show 'Current Quadrant' label for complete quiz")
	}
}

func TestFirstCallBehaviorAndWeightAlignment(t *testing.T) {
	resetState()

	// First call should NOT process any answer, only show the first question
	response, err := handlePoliticalCompass(context.Background(), createMockRequest("political_compass", map[string]interface{}{"answer": ""}))
	if err != nil {
		t.Fatalf("unexpected error on first call: %v", err)
	}

	content := extractTextContent(response)

	// Verify first call shows start message
	if !strings.Contains(content, "Political Compass Quiz Started!") {
		t.Error("first call should show quiz started message")
	}

	// Verify first call shows question 1
	if !strings.Contains(content, "Question 1 of") {
		t.Error("first call should show Question 1")
	}

	// Verify we have exactly 1 question count but no responses yet
	if questionCount != 1 {
		t.Errorf("expected questionCount to be 1 after first call, got %d", questionCount)
	}

	if len(quizState.Responses) != 0 {
		t.Errorf("expected 0 responses after first call, got %d", len(quizState.Responses))
	}

	// Verify currentIndex is 1 (pointing to next question to be shown)
	if currentIndex != 1 {
		t.Errorf("expected currentIndex to be 1 after first call, got %d", currentIndex)
	}

	// Get the question that was shown (index 0 in shuffled questions)
	firstQuestionIndex := shuffledQuestions[0]
	firstQuestion := politicalcompass.AllQuestions[firstQuestionIndex]

	// Answer the first question
	response2, err := handlePoliticalCompass(context.Background(), createMockRequest("political_compass", map[string]interface{}{"answer": "Agree"}))
	if err != nil {
		t.Fatalf("unexpected error answering first question: %v", err)
	}

	// Verify the response was recorded and scores calculated
	if len(quizState.Responses) != 1 {
		t.Errorf("expected 1 response after answering first question, got %d", len(quizState.Responses))
	}

	if quizState.Responses[0] != politicalcompass.Agree {
		t.Errorf("expected first response to be Agree, got %v", quizState.Responses[0])
	}

	// Verify that the scores match the weights for the first question with "Agree" response
	expectedEconomicScore := firstQuestion.Economic[int(politicalcompass.Agree)]
	expectedSocialScore := firstQuestion.Social[int(politicalcompass.Agree)]

	if totalEconomicScore != expectedEconomicScore {
		t.Errorf("expected economic score %f, got %f", expectedEconomicScore, totalEconomicScore)
	}

	if totalSocialScore != expectedSocialScore {
		t.Errorf("expected social score %f, got %f", expectedSocialScore, totalSocialScore)
	}

	content2 := extractTextContent(response2)

	// Verify second call shows response recorded
	if !strings.Contains(content2, "Response recorded!") {
		t.Error("second call should show response recorded message")
	}

	// Verify second call shows question 2
	if !strings.Contains(content2, "Question 2 of") {
		t.Error("second call should show Question 2")
	}

	// Verify progress shows 1 completed
	if !strings.Contains(content2, "Progress: 1 of") {
		t.Error("second call should show progress of 1 completed")
	}
}

func TestWeightAssignmentConsistency(t *testing.T) {
	resetState()

	// Start quiz and answer a few questions with different responses
	handlePoliticalCompass(context.Background(), createMockRequest("political_compass", map[string]interface{}{"answer": ""})) // Start

	responses := []politicalcompass.Response{
		politicalcompass.StronglyDisagree,
		politicalcompass.Disagree,
		politicalcompass.Agree,
		politicalcompass.StronglyAgree,
	}

	responseStrings := []string{
		"Strongly Disagree",
		"Disagree",
		"Agree",
		"Strongly Agree",
	}

	expectedEconomicTotal := 0.0
	expectedSocialTotal := 0.0

	// Answer first 4 questions and track expected scores
	for i, respStr := range responseStrings {
		questionIndex := shuffledQuestions[i]
		question := politicalcompass.AllQuestions[questionIndex]
		response := responses[i]

		// Calculate expected scores before answering
		expectedEconomicScore := question.Economic[int(response)]
		expectedSocialScore := question.Social[int(response)]
		expectedEconomicTotal += expectedEconomicScore
		expectedSocialTotal += expectedSocialScore

		// Answer the question
		_, err := handlePoliticalCompass(context.Background(), createMockRequest("political_compass", map[string]interface{}{"answer": respStr}))
		if err != nil {
			t.Fatalf("unexpected error on question %d: %v", i+1, err)
		}

		// Verify accumulated scores match expectations
		if totalEconomicScore != expectedEconomicTotal {
			t.Errorf("after question %d: expected economic total %f, got %f",
				i+1, expectedEconomicTotal, totalEconomicScore)
		}

		if totalSocialScore != expectedSocialTotal {
			t.Errorf("after question %d: expected social total %f, got %f",
				i+1, expectedSocialTotal, totalSocialScore)
		}
	}
}

func TestNoFencepostErrorsInQuestionIndexing(t *testing.T) {
	resetState()

	// Start quiz
	handlePoliticalCompass(context.Background(), createMockRequest("political_compass", map[string]interface{}{"answer": ""}))

	// Verify we can answer exactly 62 questions (no more, no less)
	for i := 0; i < len(politicalcompass.AllQuestions); i++ {
		_, err := handlePoliticalCompass(context.Background(), createMockRequest("political_compass", map[string]interface{}{"answer": "Agree"}))
		if err != nil {
			t.Fatalf("unexpected error on question %d: %v", i+1, err)
		}

		// Check that we don't exceed the bounds
		if currentIndex > len(shuffledQuestions) {
			t.Errorf("currentIndex %d exceeds shuffled questions length %d after question %d",
				currentIndex, len(shuffledQuestions), i+1)
		}
	}

	// Verify we have answered exactly 62 questions
	if questionCount != len(politicalcompass.AllQuestions) {
		t.Errorf("expected questionCount to be %d, got %d",
			len(politicalcompass.AllQuestions), questionCount)
	}

	// Verify we have exactly 62 responses
	if len(quizState.Responses) != len(politicalcompass.AllQuestions) {
		t.Errorf("expected %d responses, got %d",
			len(politicalcompass.AllQuestions), len(quizState.Responses))
	}

	// Verify currentIndex equals length (pointing past the end, indicating completion)
	if currentIndex != len(shuffledQuestions) {
		t.Errorf("expected currentIndex to be %d (length), got %d",
			len(shuffledQuestions), currentIndex)
	}

	// Trying to answer another question should return completion message
	response, err := handlePoliticalCompass(context.Background(), createMockRequest("political_compass", map[string]interface{}{"answer": "Agree"}))
	if err != nil {
		t.Fatalf("unexpected error when trying to answer beyond 62 questions: %v", err)
	}

	content := extractTextContent(response)
	if !strings.Contains(content, "Political Compass Quiz Complete!") {
		t.Error("should show completion message when trying to answer beyond 62 questions")
	}
}

// TestQuizStatusEdgeCases tests specific edge cases to improve coverage
func TestQuizStatusEdgeCases(t *testing.T) {
	t.Run("StatusWithPartialProgress", func(t *testing.T) {
		resetState()

		// Start quiz and answer exactly half the questions
		handlePoliticalCompass(context.Background(), createMockRequest("political_compass", map[string]interface{}{"answer": ""}))

		halfQuestions := len(politicalcompass.AllQuestions) / 2
		for i := 0; i < halfQuestions; i++ {
			_, err := handlePoliticalCompass(context.Background(), createMockRequest("political_compass", map[string]interface{}{"answer": "agree"}))
			if err != nil {
				t.Fatalf("Unexpected error on question %d: %v", i+1, err)
			}
		}

		response, err := handleQuizStatus(context.Background(), createMockRequest("quiz_status", map[string]interface{}{}))
		if err != nil {
			t.Fatalf("Expected no error getting status, got %v", err)
		}

		text := extractTextContent(response)

		// Should show partial progress
		expectedAnswered := fmt.Sprintf("Questions answered: %d/%d", halfQuestions, len(politicalcompass.AllQuestions))
		if !strings.Contains(text, expectedAnswered) {
			t.Errorf("Expected '%s' in status, got: %s", expectedAnswered, text)
		}

		expectedRemaining := fmt.Sprintf("Questions remaining: %d", len(politicalcompass.AllQuestions)-halfQuestions)
		if !strings.Contains(text, expectedRemaining) {
			t.Errorf("Expected '%s' in status, got: %s", expectedRemaining, text)
		}

		// Should contain response distribution
		if !strings.Contains(text, "Response Distribution:") {
			t.Errorf("Expected response distribution section")
		}

		if !strings.Contains(text, "Agree:") {
			t.Errorf("Expected 'Agree' responses in distribution")
		}
	})

	t.Run("StatusWithMixedResponses", func(t *testing.T) {
		resetState()

		// Start quiz and give mixed responses to test all response types
		handlePoliticalCompass(context.Background(), createMockRequest("political_compass", map[string]interface{}{"answer": ""}))

		responses := []string{"strongly_disagree", "disagree", "agree", "strongly_agree"}
		for i, resp := range responses {
			_, err := handlePoliticalCompass(context.Background(), createMockRequest("political_compass", map[string]interface{}{"answer": resp}))
			if err != nil {
				t.Fatalf("Unexpected error on question %d with response %s: %v", i+1, resp, err)
			}
		}

		response, err := handleQuizStatus(context.Background(), createMockRequest("quiz_status", map[string]interface{}{}))
		if err != nil {
			t.Fatalf("Expected no error getting status, got %v", err)
		}

		text := extractTextContent(response)

		// Should show all response types in distribution
		responseTypes := []string{"Strongly Disagree", "Disagree", "Agree", "Strongly Agree"}
		for _, respType := range responseTypes {
			if !strings.Contains(text, respType+": 1 (25.0%)") {
				t.Errorf("Expected response type '%s' with count in distribution", respType)
			}
		}
	})

	t.Run("StatusWithNoQuestions", func(t *testing.T) {
		resetState()

		// Test status with completely empty state (no quiz started)
		response, err := handleQuizStatus(context.Background(), createMockRequest("quiz_status", map[string]interface{}{}))
		if err != nil {
			t.Fatalf("Expected no error, got %v", err)
		}

		text := extractTextContent(response)

		if !strings.Contains(text, "Questions answered: 0/") {
			t.Errorf("Expected zero answered questions in status")
		}

		if !strings.Contains(text, "No questions answered yet") {
			t.Errorf("Expected message about no questions answered")
		}
	})

	t.Run("StatusWithCompletedQuizScores", func(t *testing.T) {
		resetState()

		// Complete entire quiz with known responses
		handlePoliticalCompass(context.Background(), createMockRequest("political_compass", map[string]interface{}{"answer": ""}))

		totalQuestions := len(politicalcompass.AllQuestions)
		for i := 0; i < totalQuestions; i++ {
			_, err := handlePoliticalCompass(context.Background(), createMockRequest("political_compass", map[string]interface{}{"answer": "agree"}))
			if err != nil {
				t.Fatalf("Unexpected error completing quiz: %v", err)
			}
		}

		response, err := handleQuizStatus(context.Background(), createMockRequest("quiz_status", map[string]interface{}{}))
		if err != nil {
			t.Fatalf("Expected no error getting status, got %v", err)
		}

		text := extractTextContent(response)

		// Should show completion
		if !strings.Contains(text, "Questions answered: "+fmt.Sprintf("%d/%d", totalQuestions, totalQuestions)) {
			t.Errorf("Expected full completion in status")
		}

		if !strings.Contains(text, "Questions remaining: 0") {
			t.Errorf("Expected no remaining questions")
		}

		if !strings.Contains(text, "Completion: 100.0%") {
			t.Errorf("Expected 100%% completion")
		}

		// Should show final scores
		if !strings.Contains(text, "Final Scores:") {
			t.Errorf("Expected final scores section")
		}

		// Should show quadrant
		quadrants := []string{"Libertarian Left", "Libertarian Right", "Authoritarian Left", "Authoritarian Right"}
		foundQuadrant := false
		for _, quadrant := range quadrants {
			if strings.Contains(text, quadrant) {
				foundQuadrant = true
				break
			}
		}
		if !foundQuadrant {
			t.Errorf("Expected one of the quadrants to be mentioned")
		}
	})
}

// TestAbsAndGetQuadrantFunctions tests helper functions for complete coverage
func TestAbsAndGetQuadrantFunctions(t *testing.T) {
	t.Run("AbsFunction", func(t *testing.T) {
		testCases := []struct {
			input    float64
			expected float64
		}{
			{5.0, 5.0},
			{-5.0, 5.0},
			{0.0, 0.0},
			{-0.1, 0.1},
			{100.5, 100.5},
		}

		for _, tc := range testCases {
			result := abs(tc.input)
			if result != tc.expected {
				t.Errorf("abs(%.1f) = %.1f, expected %.1f", tc.input, result, tc.expected)
			}
		}
	})

	t.Run("GetQuadrantFunction", func(t *testing.T) {
		testCases := []struct {
			economic, social float64
			expected         string
		}{
			{1.0, 1.0, "Libertarian Right"},    // economic > 0 && social > 0
			{-1.0, 1.0, "Libertarian Left"},    // economic < 0 && social > 0
			{1.0, -1.0, "Authoritarian Right"}, // economic > 0 && social < 0
			{-1.0, -1.0, "Authoritarian Left"}, // economic < 0 && social < 0
			{0.0, 0.1, "Authoritarian Left"},   // economic == 0 (not > 0), social > 0 -> goes to else
			{0.1, 0.0, "Authoritarian Left"},   // economic > 0, social == 0 (not > 0) -> goes to else
			{0.0, 0.0, "Authoritarian Left"},   // both == 0 -> goes to else
		}

		for _, tc := range testCases {
			result := getQuadrant(tc.economic, tc.social)
			if result != tc.expected {
				t.Errorf("getQuadrant(%.1f, %.1f) = %s, expected %s",
					tc.economic, tc.social, result, tc.expected)
			}
		}
	})
}

// TestResponseVariations tests all response string variations
func TestResponseVariations(t *testing.T) {
	resetState()

	// Test all valid response variations
	responseVariations := [][]string{
		{"Strongly Disagree", "strongly_disagree"},
		{"Disagree", "disagree"},
		{"Agree", "agree"},
		{"Strongly Agree", "strongly_agree"},
	}

	for _, variations := range responseVariations {
		for _, variation := range variations {
			t.Run(variation, func(t *testing.T) {
				resetState()

				// Start quiz
				handlePoliticalCompass(context.Background(), createMockRequest("political_compass", map[string]interface{}{"answer": ""}))

				// Test response variation
				_, err := handlePoliticalCompass(context.Background(), createMockRequest("political_compass", map[string]interface{}{"answer": variation}))
				if err != nil {
					t.Errorf("Expected no error for response variation '%s', got: %v", variation, err)
				}
			})
		}
	}
}
