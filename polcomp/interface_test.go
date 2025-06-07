package polcomp

import (
	"testing"
)

func TestAllQuestions(t *testing.T) {
	// Test that we have 62 questions
	if len(AllQuestions) != 62 {
		t.Errorf("Expected 62 questions, got %d", len(AllQuestions))
	}

	// Test that all questions have valid indices
	for i, question := range AllQuestions {
		if question.Index != int32(i) {
			t.Errorf("Question %d has incorrect index: expected %d, got %d", i, i, question.Index)
		}

		// Test that Text is not empty
		if question.Text == "" {
			t.Errorf("Question %d has empty text", i)
		}

		// Test that arrays have the correct size (they should by definition since they're [4]float64)
		if len(question.Economic) != 4 {
			t.Errorf("Question %d has incorrect Economic array size: expected 4, got %d", i, len(question.Economic))
		}
		if len(question.Social) != 4 {
			t.Errorf("Question %d has incorrect Social array size: expected 4, got %d", i, len(question.Social))
		}
	}
}

func TestResponseEnum(t *testing.T) {
	// Test Response enum values
	tests := []struct {
		response Response
		expected string
	}{
		{StronglyDisagree, "Strongly Disagree"},
		{Disagree, "Disagree"},
		{Agree, "Agree"},
		{StronglyAgree, "Strongly Agree"},
	}

	for _, test := range tests {
		if test.response.String() != test.expected {
			t.Errorf("Response %d: expected %s, got %s", test.response, test.expected, test.response.String())
		}
	}

	// Test unknown response
	var unknownResponse Response = 99
	if unknownResponse.String() != "Unknown" {
		t.Errorf("Unknown response: expected 'Unknown', got %s", unknownResponse.String())
	}
}

func TestQuestionStructure(t *testing.T) {
	// Test a specific question to verify structure is correct
	question := AllQuestions[0]

	if question.Index != 0 {
		t.Errorf("First question index should be 0, got %d", question.Index)
	}

	expectedText := "If economic globalisation is inevitable, it should primarily serve humanity rather than the interests of trans-national corporations."
	if question.Text != expectedText {
		t.Errorf("First question text mismatch")
	}

	// This question should have economic scoring but no social scoring
	expectedEconomic := [4]float64{7, 5, 0, -2}
	if question.Economic != expectedEconomic {
		t.Errorf("First question economic scores: expected %v, got %v", expectedEconomic, question.Economic)
	}

	expectedSocial := [4]float64{0, 0, 0, 0}
	if question.Social != expectedSocial {
		t.Errorf("First question social scores: expected %v, got %v", expectedSocial, question.Social)
	}
}

func TestDataIntegrity(t *testing.T) {
	// Test that scoring arrays contain meaningful data (not all zeros for every question)
	hasEconomicScoring := false
	hasSocialScoring := false

	for _, question := range AllQuestions {
		// Check if any question has non-zero economic scoring
		for _, score := range question.Economic {
			if score != 0 {
				hasEconomicScoring = true
				break
			}
		}

		// Check if any question has non-zero social scoring
		for _, score := range question.Social {
			if score != 0 {
				hasSocialScoring = true
				break
			}
		}

		if hasEconomicScoring && hasSocialScoring {
			break
		}
	}

	if !hasEconomicScoring {
		t.Error("No questions have economic scoring data")
	}

	if !hasSocialScoring {
		t.Error("No questions have social scoring data")
	}

	// Test that we have diverse question texts
	uniqueTexts := make(map[string]bool)
	for _, question := range AllQuestions {
		if len(question.Text) < 10 {
			t.Errorf("Question %d has suspiciously short text: %s", question.Index, question.Text)
		}

		if uniqueTexts[question.Text] {
			t.Errorf("Duplicate question text found: %s", question.Text)
		}
		uniqueTexts[question.Text] = true
	}

	if len(uniqueTexts) != 62 {
		t.Errorf("Expected 62 unique question texts, got %d", len(uniqueTexts))
	}
}
