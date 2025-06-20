package main

import (
	"context"
	"testing"

	"github.com/x86ed/MCP-PoliticalCompass/v3/politiscales"
)

func TestPolitiscalesQuiz(t *testing.T) {
	// Reset state before testing
	resetState()

	// Test initial state
	if politiscalesQuestionCount != 0 {
		t.Errorf("Expected initial question count to be 0, got %d", politiscalesQuestionCount)
	}

	if politiscalesLanguage != "en" {
		t.Errorf("Expected default language to be 'en', got %s", politiscalesLanguage)
	}

	// Test language setting
	response, err := handleSetPolitiscalesLanguage(context.Background(), createMockRequest("set_politiscales_language", map[string]interface{}{"language": "fr"}))
	if err != nil {
		t.Errorf("Error setting language: %v", err)
	}

	if politiscalesLanguage != "fr" {
		t.Errorf("Expected language to be 'fr', got %s", politiscalesLanguage)
	}

	if response == nil {
		t.Error("Expected response, got nil")
	}

	// Test invalid language
	invalidResponse, err := handleSetPolitiscalesLanguage(context.Background(), createMockRequest("set_politiscales_language", map[string]interface{}{"language": "invalid"}))
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if !isErrorResult(invalidResponse) {
		t.Error("Expected error result for invalid language")
	}

	// Reset language to English for quiz test
	_, err = handleSetPolitiscalesLanguage(context.Background(), createMockRequest("set_politiscales_language", map[string]interface{}{"language": "en"}))
	if err != nil {
		t.Errorf("Error setting language to English: %v", err)
	}

	// Test quiz start
	response, err = handlePolitiscales(context.Background(), createMockRequest("politiscales", map[string]interface{}{"answer": ""}))
	if err != nil {
		t.Errorf("Error starting quiz: %v", err)
	}

	if politiscalesQuestionCount != 1 {
		t.Errorf("Expected question count to be 1, got %d", politiscalesQuestionCount)
	}

	if response == nil {
		t.Error("Expected response, got nil")
	}

	// Test status
	statusResponse, err := handlePolitiscalesStatus(context.Background(), createMockRequest("politiscales_status", map[string]interface{}{}))
	if err != nil {
		t.Errorf("Error getting status: %v", err)
	}

	if statusResponse == nil {
		t.Error("Expected status response, got nil")
	}

	// Test reset
	resetResponse, err := handleResetPolitiscales(context.Background(), createMockRequest("reset_politiscales", map[string]interface{}{}))
	if err != nil {
		t.Errorf("Error resetting quiz: %v", err)
	}

	if resetResponse == nil {
		t.Error("Expected reset response, got nil")
	}

	if politiscalesQuestionCount != 0 {
		t.Errorf("Expected question count to be 0 after reset, got %d", politiscalesQuestionCount)
	}
}

func TestPolitiscalesQuestionLocalization(t *testing.T) {
	// Test English (default)
	politiscalesLanguage = "en"
	text := getPolitiscalesQuestionText("constructivism_becoming_woman")
	expectedEN := "\"One is not born, but rather becomes, a woman.\""
	if text != expectedEN {
		t.Errorf("Expected English text '%s', got '%s'", expectedEN, text)
	}

	// Test French
	politiscalesLanguage = "fr"
	text = getPolitiscalesQuestionText("constructivism_becoming_woman")
	// This should return French text if available, or fallback to English
	if text == "" {
		t.Error("Expected some text, got empty string")
	}

	// Test fallback for unknown question
	text = getPolitiscalesQuestionText("unknown_question")
	if text != "unknown_question" {
		t.Errorf("Expected fallback to question key 'unknown_question', got '%s'", text)
	}

	// Reset to English
	politiscalesLanguage = "en"
}

func TestPolitiscalesScoring(t *testing.T) {
	// Reset state
	resetState()

	// Test that we can get question text
	if len(politiscales.Questions) == 0 {
		t.Error("Expected politiscales questions to be loaded")
	}

	// Test that axes are defined
	if len(politiscales.Axes) == 0 {
		t.Error("Expected politiscales axes to be defined")
	}

	// Test that English translations exist
	if len(politiscales.ENQuestions) == 0 {
		t.Error("Expected English question translations to exist")
	}
}
