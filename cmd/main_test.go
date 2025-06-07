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
