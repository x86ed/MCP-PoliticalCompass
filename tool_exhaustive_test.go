package main

import (
	"context"
	"fmt"
	"strings"
	"testing"

	"github.com/x86ed/MCP-PoliticalCompass/v3/eightvalues"
	politicalcompass "github.com/x86ed/MCP-PoliticalCompass/v3/political-compass"
	"github.com/x86ed/MCP-PoliticalCompass/v3/politiscales"
)

// Exhaustive tests for handlePolitiscales to reach maximum coverage
func TestHandlePolitiscalesExhaustive(t *testing.T) {
	// Test 1: First question start with empty response
	t.Run("Start with empty response", func(t *testing.T) {
		resetState()

		response, err := handlePolitiscales(context.Background(), createMockRequest("politiscales", map[string]interface{}{"answer": ""}))
		if err != nil {
			t.Fatalf("Expected no error, got: %v", err)
		}

		content := extractTextContent(response)
		if !strings.Contains(content, "Politiscales Quiz Started!") {
			t.Error("Expected quiz start message")
		}
		if politiscalesQuestionCount != 1 {
			t.Errorf("Expected question count 1, got %d", politiscalesQuestionCount)
		}
		if politiscalesCurrentIndex != 1 {
			t.Errorf("Expected current index 1, got %d", politiscalesCurrentIndex)
		}
	})

	// Test 2: All possible response values in sequence
	responseValues := []string{"strongly_disagree", "disagree", "neutral", "agree", "strongly_agree"}
	for _, respValue := range responseValues {
		t.Run("Response: "+respValue, func(t *testing.T) {
			resetState()

			// Start quiz
			_, err := handlePolitiscales(context.Background(), createMockRequest("politiscales", map[string]interface{}{"answer": ""}))
			if err != nil {
				t.Fatalf("Error starting quiz: %v", err)
			}

			// Respond with specific value
			response, err := handlePolitiscales(context.Background(), createMockRequest("politiscales", map[string]interface{}{"answer": respValue}))
			if err != nil {
				t.Fatalf("Error with response %s: %v", respValue, err)
			}

			content := extractTextContent(response)
			if !strings.Contains(content, "Response recorded!") {
				t.Error("Expected response recorded message")
			}

			// Verify response was stored correctly
			if len(politiscalesQuizState.Responses) != 1 {
				t.Errorf("Expected 1 response stored, got %d", len(politiscalesQuizState.Responses))
			}

			// Verify the actual response value stored
			var expectedValue float64
			switch respValue {
			case "strongly_disagree":
				expectedValue = politiscales.StronglyDisagree
			case "disagree":
				expectedValue = politiscales.Disagree
			case "neutral":
				expectedValue = politiscales.Neutral
			case "agree":
				expectedValue = politiscales.Agree
			case "strongly_agree":
				expectedValue = politiscales.StronglyAgree
			}

			// Check if any stored response matches expected value
			found := false
			for _, stored := range politiscalesQuizState.Responses {
				if stored == expectedValue {
					found = true
					break
				}
			}
			if !found {
				t.Errorf("Expected response value %f not found in stored responses", expectedValue)
			}
		})
	}

	// Test 3: Multiple responses in sequence to test progression
	t.Run("Multiple response progression", func(t *testing.T) {
		resetState()

		responses := []string{"", "strongly_agree", "disagree", "neutral", "agree"}

		for i, resp := range responses {
			response, err := handlePolitiscales(context.Background(), createMockRequest("politiscales", map[string]interface{}{"answer": resp}))
			if err != nil {
				t.Fatalf("Error at step %d with response '%s': %v", i, resp, err)
			}

			content := extractTextContent(response)

			if i == 0 {
				// First call should start quiz
				if !strings.Contains(content, "Politiscales Quiz Started!") {
					t.Error("Expected quiz start message")
				}
			} else {
				// Subsequent calls should show progress
				if !strings.Contains(content, "Response recorded!") {
					t.Errorf("Expected response recorded message at step %d", i)
				}
			}

			// Verify state progression
			expectedQuestionCount := i + 1
			if politiscalesQuestionCount != expectedQuestionCount {
				t.Errorf("At step %d: expected question count %d, got %d", i, expectedQuestionCount, politiscalesQuestionCount)
			}

			if i > 0 {
				expectedResponseCount := i
				if len(politiscalesQuizState.Responses) != expectedResponseCount {
					t.Errorf("At step %d: expected %d responses, got %d", i, expectedResponseCount, len(politiscalesQuizState.Responses))
				}
			}
		}
	})

	// Test 4: Test with various scoring scenarios
	t.Run("Scoring scenarios", func(t *testing.T) {
		resetState()

		// Start quiz
		_, _ = handlePolitiscales(context.Background(), createMockRequest("politiscales", map[string]interface{}{"answer": ""}))

		// Test positive response (should add to YesWeights)
		_, err := handlePolitiscales(context.Background(), createMockRequest("politiscales", map[string]interface{}{"answer": "strongly_agree"}))
		if err != nil {
			t.Fatalf("Error with positive response: %v", err)
		}

		// Test negative response (should add to NoWeights)
		_, err = handlePolitiscales(context.Background(), createMockRequest("politiscales", map[string]interface{}{"answer": "strongly_disagree"}))
		if err != nil {
			t.Fatalf("Error with negative response: %v", err)
		}

		// Test neutral response (should not affect scores)
		_, err = handlePolitiscales(context.Background(), createMockRequest("politiscales", map[string]interface{}{"answer": "neutral"}))
		if err != nil {
			t.Fatalf("Error with neutral response: %v", err)
		}

		// Verify that scores were accumulated
		hasScores := false
		for _, score := range politiscalesAxesScores {
			if score != 0.0 {
				hasScores = true
				break
			}
		}
		if !hasScores {
			t.Error("Expected some axis scores to be non-zero after positive/negative responses")
		}
	})

	// Test 5: Test error cases for invalid responses
	invalidResponses := []string{"invalid", "maybe", "sometimes", "AGREE", "StronglyAgree", "1", "yes", "no"}
	for _, invalidResp := range invalidResponses {
		t.Run("Invalid response: "+invalidResp, func(t *testing.T) {
			resetState()

			// Start quiz first
			_, _ = handlePolitiscales(context.Background(), createMockRequest("politiscales", map[string]interface{}{"answer": ""}))

			// Try invalid response
			response, err := handlePolitiscales(context.Background(), createMockRequest("politiscales", map[string]interface{}{"answer": invalidResp}))
			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}

			if !isErrorResult(response) {
				t.Errorf("Expected error result for invalid response '%s'", invalidResp)
			}

			responseText := extractTextContent(response)
			if !strings.Contains(responseText, "invalid response") {
				t.Errorf("Expected 'invalid response' in error message, got: %s", responseText)
			}
		})
	}

	// Test 6: Test different languages during quiz execution
	t.Run("Language during quiz", func(t *testing.T) {
		resetState()

		// Set non-English language
		politiscalesLanguage = "fr"

		response, err := handlePolitiscales(context.Background(), createMockRequest("politiscales", map[string]interface{}{"answer": ""}))
		if err != nil {
			t.Fatalf("Error starting quiz in French: %v", err)
		}

		content := extractTextContent(response)
		if !strings.Contains(content, "Language: fr") {
			t.Error("Expected French language indication")
		}

		// Reset language
		politiscalesLanguage = "en"
	})

	// Test 7: Test completion with minimal questions (complete a small subset)
	t.Run("Rapid completion test", func(t *testing.T) {
		resetState()

		// Start the quiz properly
		_, err := handlePolitiscales(context.Background(), createMockRequest("politiscales", map[string]interface{}{"answer": ""}))
		if err != nil {
			t.Fatalf("Error starting quiz: %v", err)
		}

		// Answer all questions to complete the quiz
		for i := 0; i < len(politiscales.Questions); i++ {
			response, err := handlePolitiscales(context.Background(), createMockRequest("politiscales", map[string]interface{}{"answer": "agree"}))
			if err != nil {
				t.Fatalf("Error on question %d: %v", i+1, err)
			}

			// Check if quiz is complete on the last question
			if i == len(politiscales.Questions)-1 {
				content := extractTextContent(response)
				if !strings.Contains(content, "Politiscales Quiz Complete!") {
					t.Error("Expected quiz completion message")
				}
				if !strings.Contains(content, "Your Political Profile:") {
					t.Error("Expected political profile in completion")
				}
				if !strings.Contains(content, "<svg") {
					t.Error("Expected SVG chart in politiscales completion")
				}
				if !strings.Contains(content, "Render the SVG chart above") {
					t.Error("Expected SVG rendering instructions in politiscales completion")
				}
			}
		}
	})
}

// Exhaustive tests for handleEightValuesStatus edge cases
func TestHandleEightValuesStatusExhaustive(t *testing.T) {
	// Test with partial responses having mixed values
	t.Run("Mixed response values", func(t *testing.T) {
		resetState()

		// Add varied responses to test all distribution paths
		responses := []float64{
			eightvalues.StronglyDisagree,
			eightvalues.StronglyDisagree,
			eightvalues.Disagree,
			eightvalues.Neutral,
			eightvalues.Neutral,
			eightvalues.Neutral,
			eightvalues.Agree,
			eightvalues.Agree,
			eightvalues.Agree,
			eightvalues.Agree,
			eightvalues.StronglyAgree,
		}

		eightValuesQuizState.Responses = responses

		response, err := handleEightValuesStatus(context.Background(), createMockRequest("eight_values_status", map[string]interface{}{}))
		if err != nil {
			t.Fatalf("Expected no error, got: %v", err)
		}

		content := extractTextContent(response)

		// Verify all response types appear in distribution
		expectedCounts := map[string]int{
			"Strongly Disagree": 2,
			"Disagree":          1,
			"Neutral":           3,
			"Agree":             4,
			"Strongly Agree":    1,
		}

		for respType, expectedCount := range expectedCounts {
			expectedText := fmt.Sprintf("%s: %d", respType, expectedCount)
			if !strings.Contains(content, expectedText) {
				t.Errorf("Expected '%s' in response distribution", expectedText)
			}
		}
	})

	// Test with zero scores (edge case for division)
	t.Run("Zero scores edge case", func(t *testing.T) {
		resetState()

		// Set all scores to zero
		eightValuesEconScore = 0.0
		eightValuesDiplScore = 0.0
		eightValuesGovtScore = 0.0
		eightValuesSctyScore = 0.0
		eightValuesQuestionCount = len(eightvalues.Questions)
		eightValuesCurrentIndex = len(eightvalues.Questions)

		// Add neutral responses only
		for i := 0; i < len(eightvalues.Questions); i++ {
			eightValuesQuizState.Responses = append(eightValuesQuizState.Responses, eightvalues.Neutral)
		}

		response, err := handleEightValuesStatus(context.Background(), createMockRequest("eight_values_status", map[string]interface{}{}))
		if err != nil {
			t.Fatalf("Expected no error, got: %v", err)
		}

		content := extractTextContent(response)
		if !strings.Contains(content, "50.0%") {
			t.Error("Expected neutral 50% scores for zero axis values")
		}
	})

	// Test with very small scores (near-zero)
	t.Run("Near-zero scores", func(t *testing.T) {
		resetState()

		eightValuesEconScore = 0.1
		eightValuesDiplScore = -0.1
		eightValuesGovtScore = 0.01
		eightValuesSctyScore = -0.01
		eightValuesQuestionCount = len(eightvalues.Questions)
		eightValuesCurrentIndex = len(eightvalues.Questions)

		// Add responses for all questions to mark quiz as complete
		for i := 0; i < len(eightvalues.Questions); i++ {
			eightValuesQuizState.Responses = append(eightValuesQuizState.Responses, eightvalues.Neutral)
		}

		response, err := handleEightValuesStatus(context.Background(), createMockRequest("eight_values_status", map[string]interface{}{}))
		if err != nil {
			t.Fatalf("Expected no error, got: %v", err)
		}

		content := extractTextContent(response)
		if !strings.Contains(content, "Final Scores:") {
			t.Error("Expected final scores section for completed quiz")
		}
	})

	// Test maximum boundary values
	t.Run("Maximum boundary values", func(t *testing.T) {
		resetState()

		eightValuesEconScore = 100.0
		eightValuesDiplScore = -100.0
		eightValuesGovtScore = 100.0
		eightValuesSctyScore = -100.0
		eightValuesQuestionCount = len(eightvalues.Questions)
		eightValuesCurrentIndex = len(eightvalues.Questions)

		for i := 0; i < len(eightvalues.Questions); i++ {
			eightValuesQuizState.Responses = append(eightValuesQuizState.Responses, eightvalues.StronglyAgree)
		}

		response, err := handleEightValuesStatus(context.Background(), createMockRequest("eight_values_status", map[string]interface{}{}))
		if err != nil {
			t.Fatalf("Expected no error, got: %v", err)
		}

		content := extractTextContent(response)

		// Should handle extreme values gracefully
		if !strings.Contains(content, "100.0%") {
			t.Error("Expected maximum percentage values")
		}
	})

	// Test with incomplete responses (some missing)
	t.Run("Sparse responses", func(t *testing.T) {
		resetState()

		// Add responses with gaps
		eightValuesQuizState.Responses = []float64{
			eightvalues.Agree,
			// gap here
			eightvalues.Disagree,
			// gap here
			eightvalues.StronglyAgree,
		}

		response, err := handleEightValuesStatus(context.Background(), createMockRequest("eight_values_status", map[string]interface{}{}))
		if err != nil {
			t.Fatalf("Expected no error, got: %v", err)
		}

		content := extractTextContent(response)
		if !strings.Contains(content, "Questions answered: 3") {
			t.Error("Expected 3 questions answered for sparse responses")
		}
	})
}

// Exhaustive tests for getPolitiscalesQuestionText remaining edge cases
func TestGetPolitiscalesQuestionTextExhaustive(t *testing.T) {
	// Test with all edge case languages and missing translations
	t.Run("Missing translations fallback", func(t *testing.T) {
		// Create a fake question key that might not exist in all languages
		testKey := "non_existent_question_key_12345"

		languages := []string{"en", "fr", "es", "it", "ar", "ru", "zh", "invalid"}

		for _, lang := range languages {
			politiscalesLanguage = lang
			result := getPolitiscalesQuestionText(testKey)

			// Should fall back to the key itself for missing translations
			if result != testKey {
				// It might have fallen back to English, which is also valid
				if result == "" {
					t.Errorf("Got empty result for language %s with missing key", lang)
				}
			}
		}
	})

	// Test with empty and nil-like inputs
	t.Run("Edge case inputs", func(t *testing.T) {
		politiscalesLanguage = "en"

		edgeCases := []string{"", " ", "\n", "\t", "null", "undefined"}

		for _, testCase := range edgeCases {
			result := getPolitiscalesQuestionText(testCase)
			// For empty string, the function should return empty string (which is correct behavior)
			// For other edge cases, it should fallback to the input itself
			if testCase == "" {
				if result != "" {
					t.Errorf("Expected empty result for empty input, got: '%s'", result)
				}
			} else {
				if result == "" {
					t.Errorf("Got empty result for edge case input: '%s'", testCase)
				}
			}
		}
	})

	// Test language switching during function calls
	t.Run("Language switching during calls", func(t *testing.T) {
		testKey := "constructivism_becoming_woman"

		// Test rapid language switching
		politiscalesLanguage = "en"
		result1 := getPolitiscalesQuestionText(testKey)

		politiscalesLanguage = "fr"
		result2 := getPolitiscalesQuestionText(testKey)

		politiscalesLanguage = "zh"
		result3 := getPolitiscalesQuestionText(testKey)

		// All should return non-empty results
		if result1 == "" || result2 == "" || result3 == "" {
			t.Error("Expected non-empty results for all language switches")
		}

		// Results might be different (different languages) or same (fallbacks)
		// Both scenarios are valid
	})
}

// Exhaustive tests for remaining handleQuizStatus edge cases
func TestHandleQuizStatusExhaustive(t *testing.T) {
	// Test with corrupted state scenarios
	t.Run("Corrupted state handling", func(t *testing.T) {
		resetState()

		// Simulate corrupted state: responses without proper initialization
		quizState.Responses = []politicalcompass.Response{
			politicalcompass.Agree,
			politicalcompass.Disagree,
		}
		// But no shuffled questions
		shuffledQuestions = nil

		response, err := handleQuizStatus(context.Background(), createMockRequest("quiz_status", map[string]interface{}{}))
		if err != nil {
			t.Fatalf("Expected no error even with corrupted state, got: %v", err)
		}

		content := extractTextContent(response)
		if !strings.Contains(content, "Questions answered: 2") {
			t.Error("Expected to handle responses even without proper initialization")
		}
	})

	// Test with mismatched state
	t.Run("Mismatched state", func(t *testing.T) {
		resetState()

		// Initialize questions properly
		initializeQuestions()

		// Add more responses than questions asked
		questionCount = 5
		for i := 0; i < 10; i++ {
			quizState.Responses = append(quizState.Responses, politicalcompass.Agree)
		}

		response, err := handleQuizStatus(context.Background(), createMockRequest("quiz_status", map[string]interface{}{}))
		if err != nil {
			t.Fatalf("Expected no error with mismatched state, got: %v", err)
		}

		content := extractTextContent(response)
		// Should handle the mismatch gracefully
		if !strings.Contains(content, "Response Distribution:") {
			t.Error("Expected response distribution section")
		}
	})

	// Test with extreme score values
	t.Run("Extreme score values", func(t *testing.T) {
		resetState()

		// Set extreme scores
		totalEconomicScore = 1000.0
		totalSocialScore = -1000.0
		questionCount = len(politicalcompass.AllQuestions)
		currentIndex = len(politicalcompass.AllQuestions)

		// Fill with responses
		for i := 0; i < len(politicalcompass.AllQuestions); i++ {
			quizState.Responses = append(quizState.Responses, politicalcompass.StronglyAgree)
		}

		response, err := handleQuizStatus(context.Background(), createMockRequest("quiz_status", map[string]interface{}{}))
		if err != nil {
			t.Fatalf("Expected no error with extreme scores, got: %v", err)
		}

		content := extractTextContent(response)
		if !strings.Contains(content, "Your Quadrant:") {
			t.Error("Expected quadrant information with extreme scores")
		}
	})
}

// Exhaustive tests for handlePolitiscalesStatus remaining cases
func TestHandlePolitiscalesStatusExhaustive(t *testing.T) {
	// Test with special threshold indicators
	t.Run("Special threshold indicators", func(t *testing.T) {
		resetState()

		// Initialize the quiz state
		politiscalesQuizState = &PolitiscalesQuizState{
			Responses: make(map[int32]float64),
		}

		// Set up responses that would trigger special indicators
		for i := 0; i < len(politiscales.Questions); i++ {
			politiscalesQuizState.Responses[int32(i)] = 1.0 // StronglyAgree equivalent
		}

		response, err := handlePolitiscalesStatus(context.Background(), createMockRequest("politiscales_status", map[string]interface{}{}))
		if err != nil {
			t.Fatalf("Expected no error, got: %v", err)
		}

		content := extractTextContent(response)
		expectedAnswered := fmt.Sprintf("Questions answered: %d", len(politiscales.Questions))
		if !strings.Contains(content, expectedAnswered) {
			t.Errorf("Expected all questions to be answered. Content: %s", content)
		}
		if !strings.Contains(content, "Final Results:") {
			t.Error("Expected final results section")
		}
	})

	// Test with mixed language and completion states
	t.Run("Mixed language completion", func(t *testing.T) {
		resetState()

		// Set different language
		politiscalesLanguage = "zh"

		// Add partial responses
		for i := 0; i < 10; i++ {
			politiscalesQuizState.Responses[int32(i)] = politiscales.Neutral
		}

		response, err := handlePolitiscalesStatus(context.Background(), createMockRequest("politiscales_status", map[string]interface{}{}))
		if err != nil {
			t.Fatalf("Expected no error, got: %v", err)
		}

		content := extractTextContent(response)
		if !strings.Contains(content, "Language: zh") {
			t.Error("Expected Chinese language indication")
		}
		if !strings.Contains(content, "Questions answered: 10") {
			t.Error("Expected 10 questions answered")
		}

		// Reset language
		politiscalesLanguage = "en"
	})

	// Test edge case of exactly zero responses but non-zero state
	t.Run("Zero responses with state", func(t *testing.T) {
		resetState()

		// Set some state but no responses
		politiscalesQuestionCount = 5
		politiscalesCurrentIndex = 3
		// But leave responses empty

		response, err := handlePolitiscalesStatus(context.Background(), createMockRequest("politiscales_status", map[string]interface{}{}))
		if err != nil {
			t.Fatalf("Expected no error, got: %v", err)
		}

		content := extractTextContent(response)
		if !strings.Contains(content, "Questions answered: 0") {
			t.Error("Expected 0 questions answered despite other state")
		}
	})
}

// Test race conditions and concurrent access (if applicable)
func TestConcurrentAccess(t *testing.T) {
	t.Run("Concurrent reset and quiz operations", func(t *testing.T) {
		resetState()

		// This tests the mutex protection
		done := make(chan bool, 2)

		// Start a quiz in one goroutine
		go func() {
			defer func() { done <- true }()
			handlePolitiscales(context.Background(), createMockRequest("politiscales", map[string]interface{}{"answer": ""}))
		}()

		// Reset in another goroutine
		go func() {
			defer func() { done <- true }()
			handleResetPolitiscales(context.Background(), createMockRequest("reset_politiscales", map[string]interface{}{}))
		}()

		// Wait for both to complete
		<-done
		<-done

		// Should not crash or corrupt state
		response, err := handlePolitiscalesStatus(context.Background(), createMockRequest("politiscales_status", map[string]interface{}{}))
		if err != nil {
			t.Fatalf("Expected no error after concurrent operations, got: %v", err)
		}

		if response == nil {
			t.Error("Expected valid response after concurrent operations")
		}
	})
}
