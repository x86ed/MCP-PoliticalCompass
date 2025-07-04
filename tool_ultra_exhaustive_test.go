package main

import (
	"context"
	"fmt"
	"strings"
	"sync"
	"testing"

	"github.com/x86ed/MCP-PoliticalCompass/v3/eightvalues"
	"github.com/x86ed/MCP-PoliticalCompass/v3/politiscales"
)

// Ultra-exhaustive tests targeting the remaining uncovered code paths

// Test all the uncovered paths in handleEightValuesStatus
func TestHandleEightValuesStatusUncoveredPaths(t *testing.T) {
	// Test when all scores are exactly zero
	t.Run("All scores exactly zero", func(t *testing.T) {
		resetState()

		eightValuesEconScore = 0.0
		eightValuesDiplScore = 0.0
		eightValuesGovtScore = 0.0
		eightValuesSctyScore = 0.0
		eightValuesQuestionCount = len(eightvalues.Questions)
		eightValuesCurrentIndex = len(eightvalues.Questions)

		// Add responses to match question count
		for i := 0; i < len(eightvalues.Questions); i++ {
			eightValuesQuizState.Responses = append(eightValuesQuizState.Responses, eightvalues.Neutral)
		}

		response, err := handleEightValuesStatus(context.Background(), createMockRequest("eight_values_status", map[string]interface{}{}))
		if err != nil {
			t.Fatalf("Expected no error, got: %v", err)
		}

		content := extractTextContent(response)
		if !strings.Contains(content, "Final Scores:") {
			t.Error("Expected final scores section with all zero scores")
		}

		// Should contain Centrist labels for all axes
		if !strings.Contains(content, "Centrist") {
			t.Error("Expected Centrist labels for zero scores")
		}
	})

	// Test the response distribution section which should improve coverage
	t.Run("Response distribution edge cases", func(t *testing.T) {
		resetState()

		// Create responses with known values to complete the quiz
		responses := []float64{}
		for i := 0; i < len(eightvalues.Questions); i++ {
			// Alternate between different response types
			switch i % 5 {
			case 0:
				responses = append(responses, eightvalues.StronglyDisagree)
			case 1:
				responses = append(responses, eightvalues.Disagree)
			case 2:
				responses = append(responses, eightvalues.Neutral)
			case 3:
				responses = append(responses, eightvalues.Agree)
			case 4:
				responses = append(responses, eightvalues.StronglyAgree)
			}
		}

		eightValuesQuizState.Responses = responses
		eightValuesQuestionCount = len(responses)
		eightValuesCurrentIndex = len(responses)

		response, err := handleEightValuesStatus(context.Background(), createMockRequest("eight_values_status", map[string]interface{}{}))
		if err != nil {
			t.Fatalf("Expected no error, got: %v", err)
		}

		content := extractTextContent(response)

		// Should contain response distribution
		if !strings.Contains(content, "Response Distribution:") {
			t.Error("Expected response distribution section")
		}

		// Should show final scores since quiz is complete
		if !strings.Contains(content, "Final Scores:") {
			t.Error("Expected final scores section for completed quiz")
		}
	})

	// Test the uncovered state in score calculation
	t.Run("Score calculation without full quiz completion", func(t *testing.T) {
		resetState()

		// Add only a few responses but don't complete the quiz
		eightValuesQuizState.Responses = []float64{
			eightvalues.StronglyAgree,
			eightvalues.StronglyDisagree,
			eightvalues.Neutral,
		}

		// Set scores manually as if partially calculated
		eightValuesEconScore = 25.5
		eightValuesDiplScore = -15.2
		eightValuesGovtScore = 5.1
		eightValuesSctyScore = -8.7

		eightValuesQuestionCount = 3
		eightValuesCurrentIndex = 3

		response, err := handleEightValuesStatus(context.Background(), createMockRequest("eight_values_status", map[string]interface{}{}))
		if err != nil {
			t.Fatalf("Expected no error, got: %v", err)
		}

		content := extractTextContent(response)

		// Should NOT contain "Final Scores" since quiz is not complete
		if strings.Contains(content, "Final Scores:") {
			t.Error("Should not contain final scores for incomplete quiz")
		}

		// Should contain progress information
		if !strings.Contains(content, "Questions answered: 3") {
			t.Error("Expected progress information")
		}
	})
}

// Test all edge cases in handlePolitiscales that might not be covered
func TestHandlePolitiscalesUncoveredPaths(t *testing.T) {
	// Test the exact moment of quiz completion
	t.Run("Quiz completion edge case", func(t *testing.T) {
		resetState()

		// Start quiz
		_, err := handlePolitiscales(context.Background(), createMockRequest("politiscales", map[string]interface{}{"answer": ""}))
		if err != nil {
			t.Fatalf("Error starting quiz: %v", err)
		}

		// Answer all questions except the last one
		for i := 0; i < len(politiscales.Questions)-1; i++ {
			_, err := handlePolitiscales(context.Background(), createMockRequest("politiscales", map[string]interface{}{"answer": "neutral"}))
			if err != nil {
				t.Fatalf("Error on question %d: %v", i+1, err)
			}
		}

		// Now answer the final question and ensure completion triggers
		response, err := handlePolitiscales(context.Background(), createMockRequest("politiscales", map[string]interface{}{"answer": "strongly_agree"}))
		if err != nil {
			t.Fatalf("Error on final question: %v", err)
		}

		content := extractTextContent(response)

		// Must contain completion message
		if !strings.Contains(content, "Politiscales Quiz Complete!") {
			t.Error("Expected completion message on final question")
		}

		// Must contain political profile
		if !strings.Contains(content, "Your Political Profile:") {
			t.Error("Expected political profile on completion")
		}
	})

	// Test state consistency during language changes
	t.Run("Language change consistency", func(t *testing.T) {
		resetState()

		// Start quiz in English
		politiscalesLanguage = "en"
		_, err := handlePolitiscales(context.Background(), createMockRequest("politiscales", map[string]interface{}{"answer": ""}))
		if err != nil {
			t.Fatalf("Error starting quiz: %v", err)
		}

		// Answer a few questions
		for i := 0; i < 3; i++ {
			_, err := handlePolitiscales(context.Background(), createMockRequest("politiscales", map[string]interface{}{"answer": "agree"}))
			if err != nil {
				t.Fatalf("Error on question %d: %v", i+1, err)
			}
		}

		// Change language mid-quiz
		_, err = handleSetPolitiscalesLanguage(context.Background(), createMockRequest("set_politiscales_language", map[string]interface{}{"language": "fr"}))
		if err != nil {
			t.Fatalf("Error changing language: %v", err)
		}

		// Continue with more questions
		response, err := handlePolitiscales(context.Background(), createMockRequest("politiscales", map[string]interface{}{"answer": "disagree"}))
		if err != nil {
			t.Fatalf("Error after language change: %v", err)
		}

		content := extractTextContent(response)

		// Should continue working (language will be French for subsequent questions)
		if !strings.Contains(content, "Response recorded!") {
			t.Error("Expected quiz to continue working after language change")
		}
	})
}

// Test concurrent access scenarios for race conditions
func TestConcurrentAccessUltraExhaustive(t *testing.T) {
	t.Run("Heavy concurrent operations", func(t *testing.T) {
		resetState()

		var wg sync.WaitGroup
		numGoroutines := 50

		// Concurrent quiz starts
		for i := 0; i < numGoroutines; i++ {
			wg.Add(1)
			go func(id int) {
				defer wg.Done()

				// Each goroutine performs different operations
				switch id % 4 {
				case 0:
					// Start political compass quiz
					handlePoliticalCompass(context.Background(), createMockRequest("political_compass", map[string]interface{}{"answer": ""}))
				case 1:
					// Start 8values quiz
					handleEightValues(context.Background(), createMockRequest("eight_values", map[string]interface{}{"answer": ""}))
				case 2:
					// Start politiscales quiz
					handlePolitiscales(context.Background(), createMockRequest("politiscales", map[string]interface{}{"answer": ""}))
				case 3:
					// Check status
					handleQuizStatus(context.Background(), createMockRequest("quiz_status", map[string]interface{}{}))
				}
			}(i)
		}

		wg.Wait()

		// Verify that the system is still in a consistent state
		// At least one quiz should have been started
		if quizState == nil && eightValuesQuizState == nil && politiscalesQuizState == nil {
			t.Error("Expected at least one quiz to be started")
		}
	})

	// Test concurrent resets and status checks
	t.Run("Concurrent resets and operations", func(t *testing.T) {
		resetState()

		var wg sync.WaitGroup
		numOps := 30

		for i := 0; i < numOps; i++ {
			wg.Add(1)
			go func(id int) {
				defer wg.Done()

				switch id % 6 {
				case 0:
					handleResetQuiz(context.Background(), createMockRequest("reset_quiz", map[string]interface{}{}))
				case 1:
					handleResetEightValues(context.Background(), createMockRequest("reset_eight_values", map[string]interface{}{}))
				case 2:
					handleResetPolitiscales(context.Background(), createMockRequest("reset_politiscales", map[string]interface{}{}))
				case 3:
					handleQuizStatus(context.Background(), createMockRequest("quiz_status", map[string]interface{}{}))
				case 4:
					handleEightValuesStatus(context.Background(), createMockRequest("eight_values_status", map[string]interface{}{}))
				case 5:
					handlePolitiscalesStatus(context.Background(), createMockRequest("politiscales_status", map[string]interface{}{}))
				}
			}(i)
		}

		wg.Wait()

		// System should still be functional
		_, err := handleQuizStatus(context.Background(), createMockRequest("quiz_status", map[string]interface{}{}))
		if err != nil {
			t.Errorf("System not functional after concurrent operations: %v", err)
		}
	})
}

// Test memory and state management edge cases
func TestMemoryAndStateEdgeCases(t *testing.T) {
	t.Run("Large response arrays", func(t *testing.T) {
		resetState()

		// Create artificially large response arrays
		largeResponseCount := 1000

		// Test 8values with large response array
		for i := 0; i < largeResponseCount; i++ {
			eightValuesQuizState.Responses = append(eightValuesQuizState.Responses,
				float64(i%5-2)) // Vary between -2 and 2
		}

		eightValuesQuestionCount = largeResponseCount
		eightValuesCurrentIndex = largeResponseCount

		response, err := handleEightValuesStatus(context.Background(), createMockRequest("eight_values_status", map[string]interface{}{}))
		if err != nil {
			t.Fatalf("Error with large response array: %v", err)
		}

		content := extractTextContent(response)
		expectedCount := fmt.Sprintf("Questions answered: %d", largeResponseCount)
		if !strings.Contains(content, expectedCount) {
			t.Error("Expected correct count with large response array")
		}
	})

	t.Run("State reset during operations", func(t *testing.T) {
		resetState()

		// Start multiple quizzes
		handlePoliticalCompass(context.Background(), createMockRequest("political_compass", map[string]interface{}{"answer": ""}))
		handleEightValues(context.Background(), createMockRequest("eight_values", map[string]interface{}{"answer": ""}))
		handlePolitiscales(context.Background(), createMockRequest("politiscales", map[string]interface{}{"answer": ""}))

		// Answer some questions
		handlePoliticalCompass(context.Background(), createMockRequest("political_compass", map[string]interface{}{"answer": "agree"}))
		handleEightValues(context.Background(), createMockRequest("eight_values", map[string]interface{}{"answer": "neutral"}))
		handlePolitiscales(context.Background(), createMockRequest("politiscales", map[string]interface{}{"answer": "disagree"}))

		// Reset one quiz
		handleResetQuiz(context.Background(), createMockRequest("reset_quiz", map[string]interface{}{}))

		// Check that other quizzes are still intact
		statusResponse, err := handleEightValuesStatus(context.Background(), createMockRequest("eight_values_status", map[string]interface{}{}))
		if err != nil {
			t.Fatalf("Error checking 8values status after political compass reset: %v", err)
		}

		content := extractTextContent(statusResponse)
		if !strings.Contains(content, "Questions answered: 1") {
			t.Error("8values quiz should still have 1 question answered after political compass reset")
		}
	})
}

// Test specific uncovered error conditions
func TestSpecificErrorConditions(t *testing.T) {
	t.Run("Invalid state transitions", func(t *testing.T) {
		resetState()

		// Try to use a response before starting quiz (should handle gracefully)
		response, err := handlePoliticalCompass(context.Background(), createMockRequest("political_compass", map[string]interface{}{"answer": "agree"}))
		if err != nil {
			t.Fatalf("Unexpected error: %v", err)
		}

		// Should start the quiz instead
		content := extractTextContent(response)
		if !strings.Contains(content, "Political Compass Quiz Started!") {
			t.Error("Expected quiz to start when response given before initialization")
		}
	})

	t.Run("Boundary response values", func(t *testing.T) {
		resetState()

		// Test with responses that might cause calculation issues
		testResponses := []string{
			"strongly_disagree", "disagree", "agree", "strongly_agree",
		}

		for _, resp := range testResponses {
			// Test each response type for each quiz
			for i := 0; i < 3; i++ { // Test with multiple questions per response type
				resetState()

				// Political compass (only supports 4 response types)
				_, err := handlePoliticalCompass(context.Background(), createMockRequest("political_compass", map[string]interface{}{"answer": ""}))
				if err != nil {
					t.Fatalf("Error starting political compass: %v", err)
				}

				_, err = handlePoliticalCompass(context.Background(), createMockRequest("political_compass", map[string]interface{}{"answer": resp}))
				if err != nil {
					t.Fatalf("Error with political compass response %s: %v", resp, err)
				}

				// 8values (supports 5 response types including neutral)
				resetState()
				_, err = handleEightValues(context.Background(), createMockRequest("eight_values", map[string]interface{}{"answer": ""}))
				if err != nil {
					t.Fatalf("Error starting 8values: %v", err)
				}

				eightValuesResponse := resp
				if i == 0 && resp == "strongly_disagree" {
					// Test neutral for 8values
					eightValuesResponse = "neutral"
				}

				_, err = handleEightValues(context.Background(), createMockRequest("eight_values", map[string]interface{}{"answer": eightValuesResponse}))
				if err != nil {
					t.Fatalf("Error with 8values response %s: %v", eightValuesResponse, err)
				}

				// Politiscales (supports 5 response types including neutral)
				resetState()
				_, err = handlePolitiscales(context.Background(), createMockRequest("politiscales", map[string]interface{}{"answer": ""}))
				if err != nil {
					t.Fatalf("Error starting politiscales: %v", err)
				}

				politiscalesResponse := resp
				if i == 1 && resp == "disagree" {
					// Test neutral for politiscales
					politiscalesResponse = "neutral"
				}

				_, err = handlePolitiscales(context.Background(), createMockRequest("politiscales", map[string]interface{}{"answer": politiscalesResponse}))
				if err != nil {
					t.Fatalf("Error with politiscales response %s: %v", politiscalesResponse, err)
				}
			}
		}
	})
}
