package main

import (
	"context"
	"fmt"
	"strings"
	"testing"

	"github.com/x86ed/MCP-PoliticalCompass/v3/eightvalues"
)

func TestToolHandlersDirectly(t *testing.T) {
	t.Run("political compass tool handler", func(t *testing.T) {
		resetState()

		response, err := handlePoliticalCompass(context.Background(), createMockRequest("political_compass", map[string]interface{}{"answer": ""}))
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}

		if response == nil {
			t.Fatal("political compass handler response is nil")
		}

		if len(response.Content) == 0 {
			t.Fatal("political compass handler response content is empty")
		}

		if response.Content[0] == nil {
			t.Fatal("political compass handler response content is not valid")
		}
	})

	t.Run("reset quiz tool handler", func(t *testing.T) {
		resetState()

		handlePoliticalCompass(context.Background(), createMockRequest("political_compass", map[string]interface{}{"answer": ""}))

		response, err := handleResetQuiz(context.Background(), createMockRequest("reset_quiz", map[string]interface{}{}))
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}

		if response == nil {
			t.Fatal("reset quiz handler response is nil")
		}

		if len(response.Content) == 0 {
			t.Fatal("reset quiz handler response content is empty")
		}

		if response.Content[0] == nil {
			t.Fatal("reset quiz handler response content is not valid")
		}
	})
}

// TestIntegrationEdgeCases adds more comprehensive integration test scenarios
func TestIntegrationEdgeCases(t *testing.T) {
	t.Run("MixedQuizWorkflow", func(t *testing.T) {
		resetState()

		// Test switching between quizzes
		// Start political compass
		_, err := handlePoliticalCompass(context.Background(), createMockRequest("political_compass", map[string]interface{}{"answer": ""}))
		if err != nil {
			t.Fatalf("Error starting political compass: %v", err)
		}

		// Answer a few questions
		for i := 0; i < 3; i++ {
			_, err = handlePoliticalCompass(context.Background(), createMockRequest("political_compass", map[string]interface{}{"answer": "agree"}))
			if err != nil {
				t.Fatalf("Error answering political compass question: %v", err)
			}
		}

		// Check political compass status
		_, err = handleQuizStatus(context.Background(), createMockRequest("quiz_status", map[string]interface{}{}))
		if err != nil {
			t.Fatalf("Error getting political compass status: %v", err)
		}

		// Start 8values quiz
		_, err = handleEightValues(context.Background(), createMockRequest("eight_values", map[string]interface{}{"answer": ""}))
		if err != nil {
			t.Fatalf("Error starting 8values quiz: %v", err)
		}

		// Answer a few 8values questions
		for i := 0; i < 3; i++ {
			_, err = handleEightValues(context.Background(), createMockRequest("eight_values", map[string]interface{}{"answer": "neutral"}))
			if err != nil {
				t.Fatalf("Error answering 8values question: %v", err)
			}
		}

		// Check 8values status
		_, err = handleEightValuesStatus(context.Background(), createMockRequest("eight_values_status", map[string]interface{}{}))
		if err != nil {
			t.Fatalf("Error getting 8values status: %v", err)
		}

		// Reset both quizzes
		_, err = handleResetQuiz(context.Background(), createMockRequest("reset_quiz", map[string]interface{}{}))
		if err != nil {
			t.Fatalf("Error resetting political compass: %v", err)
		}

		_, err = handleResetEightValues(context.Background(), createMockRequest("reset_eight_values", map[string]interface{}{}))
		if err != nil {
			t.Fatalf("Error resetting 8values: %v", err)
		}

		// Verify both are reset
		pcStatus, _ := handleQuizStatus(context.Background(), createMockRequest("quiz_status", map[string]interface{}{}))
		pcText := extractTextContent(pcStatus)
		if !strings.Contains(pcText, "Questions answered: 0/") {
			t.Error("Political compass should be reset")
		}

		evStatus, _ := handleEightValuesStatus(context.Background(), createMockRequest("eight_values_status", map[string]interface{}{}))
		evText := extractTextContent(evStatus)
		totalQuestions := len(eightvalues.Questions) // Dynamically derive total question count
		expectedText := fmt.Sprintf("Questions answered: 0/%d", totalQuestions)
		if !strings.Contains(evText, expectedText) {
			t.Error("8values should be reset")
		}
	})

	t.Run("ErrorHandlingWorkflow", func(t *testing.T) {
		resetState() // Test error handling in various scenarios

		// Start quiz first to enable response processing
		_, err := handlePoliticalCompass(context.Background(), createMockRequest("political_compass", map[string]interface{}{"answer": ""}))
		if err != nil {
			t.Fatalf("Error starting quiz: %v", err)
		}

		// Try invalid response after starting quiz
		response, err := handlePoliticalCompass(context.Background(), createMockRequest("political_compass", map[string]interface{}{"answer": "invalid"}))
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		if !isErrorResult(response) {
			t.Error("Expected error result for invalid response")
		}

		// Try various invalid responses
		invalidResponses := []string{"maybe", "sometimes", "invalid", "yes", "no", "1", "2"}
		for _, invalid := range invalidResponses {
			response, err := handlePoliticalCompass(context.Background(), createMockRequest("political_compass", map[string]interface{}{"answer": invalid}))
			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}
			if !isErrorResult(response) {
				t.Errorf("Expected error result for invalid response '%s'", invalid)
			}
		}

		// Similar test for 8values
		_, err = handleEightValues(context.Background(), createMockRequest("eight_values", map[string]interface{}{"answer": ""}))
		if err != nil {
			t.Fatalf("Error starting 8values: %v", err)
		}

		invalidEightValuesResponses := []string{"maybe", "sometimes", "invalid", "yes", "no", "1", "2", "kinda"}
		for _, invalid := range invalidEightValuesResponses {
			response, err := handleEightValues(context.Background(), createMockRequest("eight_values", map[string]interface{}{"answer": invalid}))
			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}
			if !isErrorResult(response) {
				t.Errorf("Expected error result for invalid 8values response '%s'", invalid)
			}
		}
	})

	t.Run("StateConsistencyWorkflow", func(t *testing.T) {
		resetState()

		// Test that state remains consistent across operations

		// Start and partially complete political compass
		handlePoliticalCompass(context.Background(), createMockRequest("political_compass", map[string]interface{}{"answer": ""}))
		for i := 0; i < 5; i++ {
			handlePoliticalCompass(context.Background(), createMockRequest("political_compass", map[string]interface{}{"answer": "agree"}))
		}

		// Check that question count is correct
		if questionCount != 6 { // 5 answers + 1 initial question shown
			t.Errorf("Expected questionCount 6, got %d", questionCount)
		}

		if len(quizState.Responses) != 5 {
			t.Errorf("Expected 5 recorded responses, got %d", len(quizState.Responses))
		}

		// Start and partially complete 8values
		handleEightValues(context.Background(), createMockRequest("eight_values", map[string]interface{}{"answer": ""}))
		for i := 0; i < 3; i++ {
			handleEightValues(context.Background(), createMockRequest("eight_values", map[string]interface{}{"answer": "agree"}))
		}

		// Check that 8values state is separate
		if eightValuesQuestionCount != 4 { // 3 answers + 1 initial question shown
			t.Errorf("Expected eightValuesQuestionCount 4, got %d", eightValuesQuestionCount)
		}

		if len(eightValuesQuizState.Responses) != 3 {
			t.Errorf("Expected 3 recorded 8values responses, got %d", len(eightValuesQuizState.Responses))
		}

		// Political compass state should be unchanged
		if questionCount != 6 {
			t.Errorf("Political compass questionCount should still be 6, got %d", questionCount)
		}

		if len(quizState.Responses) != 5 {
			t.Errorf("Political compass responses should still be 5, got %d", len(quizState.Responses))
		}
	})
}
