package eightvalues

import (
	"testing"
)

func TestQuestionStructure(t *testing.T) {
	// Test that all questions have valid structure
	for i, question := range Questions {
		if question.Index != int32(i) {
			t.Errorf("Question %d has incorrect index %d", i, question.Index)
		}

		if question.Text == "" {
			t.Errorf("Question %d has empty text", i)
		}

		// Verify effect array has 4 elements
		if len(question.Effect) != 4 {
			t.Errorf("Question %d effect array should have 4 elements, got %d", i, len(question.Effect))
		}
	}
}

func TestQuestionCount(t *testing.T) {
	expectedCount := 70 // 8values has 70 questions
	if len(Questions) != expectedCount {
		t.Errorf("Expected %d questions, got %d", expectedCount, len(Questions))
	}
}

func TestResponseConstants(t *testing.T) {
	// Test that response constants are properly defined
	if StronglyDisagree >= 0 {
		t.Error("StronglyDisagree should be negative")
	}
	if Disagree >= 0 {
		t.Error("Disagree should be negative")
	}
	if Neutral != 0 {
		t.Error("Neutral should be zero")
	}
	if Agree <= 0 {
		t.Error("Agree should be positive")
	}
	if StronglyAgree <= 0 {
		t.Error("StronglyAgree should be positive")
	}
}

func TestAxisConstants(t *testing.T) {
	// Test that axis constants are properly defined
	if Economic < 0 || Economic > 3 {
		t.Errorf("Economic axis constant should be 0-3, got %d", Economic)
	}
	if Diplomatic < 0 || Diplomatic > 3 {
		t.Errorf("Diplomatic axis constant should be 0-3, got %d", Diplomatic)
	}
	if Government < 0 || Government > 3 {
		t.Errorf("Government axis constant should be 0-3, got %d", Government)
	}
	if Society < 0 || Society > 3 {
		t.Errorf("Society axis constant should be 0-3, got %d", Society)
	}
}

func TestQuestionEffects(t *testing.T) {
	// Test that questions have meaningful effects (not all zeros)
	for i, question := range Questions {
		allZero := true
		for _, effect := range question.Effect {
			if effect != 0 {
				allZero = false
				break
			}
		}
		if allZero {
			t.Errorf("Question %d has all zero effects", i)
		}
	}
}

func TestSampleQuestions(t *testing.T) {
	// Test specific known questions to ensure they're loaded correctly
	if len(Questions) > 0 {
		firstQuestion := Questions[0]
		if firstQuestion.Index != 0 {
			t.Error("First question should have index 0")
		}
		if firstQuestion.Text == "" {
			t.Error("First question should have text")
		}
	}

	if len(Questions) > 10 {
		// Test a question in the middle
		midQuestion := Questions[10]
		if midQuestion.Index != int32(10) {
			t.Error("Question 10 should have index 10")
		}
	}
}
