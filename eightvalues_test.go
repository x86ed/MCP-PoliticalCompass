package main

import (
	"context"
	"strings"
	"testing"

	"github.com/x86ed/MCP-PoliticalCompass/v3/eightvalues"
)

func TestEightValuesToolStart(t *testing.T) {
	// Reset state first
	resetState()

	request := createMockRequest("eightvalues", map[string]interface{}{
		"answer": "",
	})

	response, err := handleEightValues(context.Background(), request)
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	text := extractTextContent(response)

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
	startRequest := createMockRequest("eightvalues", map[string]interface{}{
		"answer": "",
	})
	_, err := handleEightValues(context.Background(), startRequest)
	if err != nil {
		t.Fatalf("Expected no error starting quiz, got %v", err)
	}

	// Test invalid response
	invalidRequest := createMockRequest("eightvalues", map[string]interface{}{
		"answer": "invalid_response",
	})

	result, err := handleEightValues(context.Background(), invalidRequest)
	if err != nil {
		t.Fatalf("Unexpected Go error: %v", err)
	}

	// Check if result indicates an error
	text := extractTextContent(result)
	expectedError := "invalid response: invalid_response"
	if !strings.Contains(text, expectedError) {
		t.Errorf("Expected error containing '%s', got: %s", expectedError, text)
	}
}

func TestEightValuesAllResponseTypes(t *testing.T) {
	responses := []string{"strongly_disagree", "disagree", "neutral", "agree", "strongly_agree"}

	for _, response := range responses {
		t.Run(response, func(t *testing.T) {
			// Reset state
			resetState()

			// Start quiz
			startRequest := createMockRequest("eightvalues", map[string]interface{}{
				"answer": "",
			})
			_, err := handleEightValues(context.Background(), startRequest)
			if err != nil {
				t.Fatalf("Expected no error starting quiz, got %v", err)
			}

			// Test response
			responseRequest := createMockRequest("eightvalues", map[string]interface{}{
				"answer": response,
			})
			result, err := handleEightValues(context.Background(), responseRequest)
			if err != nil {
				t.Fatalf("Expected no error for response %s, got %v", response, err)
			}

			text := extractTextContent(result)

			if !strings.Contains(text, "Response recorded!") {
				t.Errorf("Expected response recorded message for %s, got: %s", response, text)
			}
		})
	}
}

func TestResetEightValuesTool(t *testing.T) {
	// Start a quiz and answer some questions
	resetState()

	startRequest := createMockRequest("eightvalues", map[string]interface{}{
		"answer": "",
	})
	_, err := handleEightValues(context.Background(), startRequest)
	if err != nil {
		t.Fatalf("Expected no error starting quiz, got %v", err)
	}

	// Answer a few questions
	for i := 0; i < 3; i++ {
		responseRequest := createMockRequest("eightvalues", map[string]interface{}{
			"answer": "agree",
		})
		_, err := handleEightValues(context.Background(), responseRequest)
		if err != nil {
			t.Fatalf("Expected no error on question %d, got %v", i+1, err)
		}
	}

	// Reset the quiz
	resetRequest := createMockRequest("reset_eightvalues", map[string]interface{}{})
	response, err := handleResetEightValues(context.Background(), resetRequest)
	if err != nil {
		t.Fatalf("Expected no error resetting quiz, got %v", err)
	}

	text := extractTextContent(response)

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
	resetState()

	statusRequest := createMockRequest("eightvalues_status", map[string]interface{}{})
	response, err := handleEightValuesStatus(context.Background(), statusRequest)
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	text := extractTextContent(response)

	if !strings.Contains(text, "8values Quiz Status") {
		t.Errorf("Expected status header, got: %s", text)
	}

	if !strings.Contains(text, "Questions answered: 0") {
		t.Errorf("Expected question count, got: %s", text)
	}
}

func TestEightValuesSVGGeneration(t *testing.T) {
	// Test SVG generation using the eightvalues package
	svg := eightvalues.GenerateSVG(50.0, 50.0, 50.0, 50.0)

	if !strings.Contains(svg, "<svg") {
		t.Errorf("Expected SVG content")
	}

	if !strings.Contains(svg, "</svg>") {
		t.Errorf("Expected complete SVG")
	}
}

func TestGetQuadrant(t *testing.T) {
	tests := []struct {
		name     string
		econ     float64
		social   float64
		expected string
	}{
		{"Center", 0, 0, "Authoritarian Left"}, // (0,0) falls into the else case
		{"Auth Left", -5, -5, "Authoritarian Left"},
		{"Auth Right", 5, -5, "Authoritarian Right"},
		{"Lib Left", -5, 5, "Libertarian Left"},
		{"Lib Right", 5, 5, "Libertarian Right"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := getQuadrant(tt.econ, tt.social)
			if result != tt.expected {
				t.Errorf("getQuadrant(%f, %f) = %s, want %s",
					tt.econ, tt.social, result, tt.expected)
			}
		})
	}
}
