package politiscales

import (
	"strings"
	"testing"
)

func TestPolitiscalesModule(t *testing.T) {
	// Basic test to ensure the module compiles and loads
	// Add more specific tests here as the module develops
	t.Log("Politiscales module test placeholder")
}

func TestQuestionsSlice(t *testing.T) {
	// Test that we have the correct number of questions (117 from weights.js)
	expectedCount := 117
	if len(Questions) != expectedCount {
		t.Errorf("Expected %d questions, got %d", expectedCount, len(Questions))
	}

	// Test that all questions have valid indices
	for i, question := range Questions {
		if question.Index != int32(i) {
			t.Errorf("Question %d has incorrect index: expected %d, got %d", i, i, question.Index)
		}

		// Test that Text is not empty and contains the key from weights.js
		if question.Text == "" {
			t.Errorf("Question %d has empty text", i)
		}

		// Test that the first few questions have the expected keys from weights.js
		if i < 10 {
			expectedKeys := []string{
				"constructivism_becoming_woman",
				"constructivism_racism_presence",
				"constructivism_science_society",
				"constructivism_gender_categories",
				"constructivism_criminality_nature",
				"constructivism_sexual_orientation",
				"constructivism_ethnic_differences",
				"essentialism_gender_biology",
				"essentialism_hormones_character",
				"essentialism_sexual_aggression",
			}
			if question.Text != expectedKeys[i] {
				t.Errorf("Question %d has incorrect text: expected %s, got %s", i, expectedKeys[i], question.Text)
			}
		}

		// Test that YesWeights and NoWeights are properly defined
		if len(question.YesWeights) == 0 && len(question.NoWeights) == 0 {
			t.Errorf("Question %d has no weights defined", i)
		}

		// Test that weights have valid axis names and values
		for _, weight := range question.YesWeights {
			if weight.Axis == "" {
				t.Errorf("Question %d has YesWeight with empty axis", i)
			}
			if weight.Value <= 0 {
				t.Errorf("Question %d has YesWeight with invalid value: %f", i, weight.Value)
			}
		}

		for _, weight := range question.NoWeights {
			if weight.Axis == "" {
				t.Errorf("Question %d has NoWeight with empty axis", i)
			}
			if weight.Value <= 0 {
				t.Errorf("Question %d has NoWeight with invalid value: %f", i, weight.Value)
			}
		}
	}
}

func TestAxesSlice(t *testing.T) {
	// Test that we have the correct number of axes (20 paired + 7 unpaired = 27)
	expectedCount := 27
	if len(Axes) != expectedCount {
		t.Errorf("Expected %d axes, got %d", expectedCount, len(Axes))
	}

	// Track which axes we've seen to check for duplicates
	seenAxes := make(map[string]bool)
	pairedAxesCount := 0
	unpairedAxesCount := 0

	for _, axis := range Axes {
		// Test that Name is not empty
		if axis.Name == "" {
			t.Error("Found axis with empty name")
		}

		// Check for duplicates
		if seenAxes[axis.Name] {
			t.Errorf("Duplicate axis found: %s", axis.Name)
		}
		seenAxes[axis.Name] = true

		// Count paired vs unpaired axes
		if axis.Pair != "" {
			pairedAxesCount++
		} else {
			unpairedAxesCount++
		}

		// Test that paired axes have colors
		if axis.Pair != "" && axis.Color == "" {
			t.Errorf("Paired axis %s should have a color", axis.Name)
		}

		// Test that unpaired axes have thresholds > 0
		if axis.Pair == "" && axis.Threshold <= 0 {
			t.Errorf("Unpaired axis %s should have a threshold > 0", axis.Name)
		}

		// Test threshold values are valid (between 0 and 1)
		if axis.Threshold < 0 || axis.Threshold > 1 {
			t.Errorf("Axis %s has invalid threshold: %f", axis.Name, axis.Threshold)
		}
	}

	// Verify we have the expected counts
	if pairedAxesCount != 20 {
		t.Errorf("Expected 20 paired axes, got %d", pairedAxesCount)
	}
	if unpairedAxesCount != 7 {
		t.Errorf("Expected 7 unpaired axes, got %d", unpairedAxesCount)
	}

	// Test specific axes from the JavaScript data
	expectedAxes := map[string]struct {
		pair      string
		color     string
		label     string
		threshold float64
	}{
		"constructivism":   {"identity", "#a425b6", "equality", 0.0},
		"essentialism":     {"identity", "#34b634", "", 0.0},
		"internationalism": {"globalism", "#3e6ffd", "humanity", 0.0},
		"nationalism":      {"globalism", "#ff8500", "fatherland", 0.0},
		"anarchism":        {"", "", "", 0.9},
		"pragmatism":       {"", "", "", 0.5},
		"feminism":         {"", "", "", 0.9},
	}

	for _, axis := range Axes {
		if expected, exists := expectedAxes[axis.Name]; exists {
			if axis.Pair != expected.pair {
				t.Errorf("Axis %s: expected pair %s, got %s", axis.Name, expected.pair, axis.Pair)
			}
			if axis.Color != expected.color {
				t.Errorf("Axis %s: expected color %s, got %s", axis.Name, expected.color, axis.Color)
			}
			if axis.Label != expected.label {
				t.Errorf("Axis %s: expected label %s, got %s", axis.Name, expected.label, axis.Label)
			}
			if axis.Threshold != expected.threshold {
				t.Errorf("Axis %s: expected threshold %f, got %f", axis.Name, expected.threshold, axis.Threshold)
			}
		}
	}
}

func TestENCopyMap(t *testing.T) {
	// Test that ENCopy map has the expected number of UI elements
	expectedUIElements := 10
	if len(ENCopy) != expectedUIElements {
		t.Errorf("Expected %d UI elements in ENCopy, got %d", expectedUIElements, len(ENCopy))
	}

	// Test that all required UI keys exist
	requiredKeys := []string{
		"start_test", "question_x_of_n", "back_home", "prev_question",
		"strong_disagree", "disagree", "neutral", "agree", "strong_agree", "result",
	}

	for _, key := range requiredKeys {
		if value, exists := ENCopy[key]; !exists {
			t.Errorf("Required UI key '%s' missing from ENCopy", key)
		} else if value == "" {
			t.Errorf("UI key '%s' has empty value", key)
		}
	}

	// Test specific UI translations
	expectedTranslations := map[string]string{
		"start_test":      "Start the test",
		"strong_disagree": "Strongly disagree",
		"result":          "Result",
	}

	for key, expectedValue := range expectedTranslations {
		if actualValue, exists := ENCopy[key]; !exists {
			t.Errorf("Key '%s' missing from ENCopy", key)
		} else if actualValue != expectedValue {
			t.Errorf("Key '%s': expected '%s', got '%s'", key, expectedValue, actualValue)
		}
	}
}

func TestENQuestionsMap(t *testing.T) {
	// Test that ENQuestions map has the expected number of questions (117)
	expectedQuestionCount := 117
	if len(ENQuestions) != expectedQuestionCount {
		t.Errorf("Expected %d questions in ENQuestions, got %d", expectedQuestionCount, len(ENQuestions))
	}

	// Test that all question keys from Questions slice exist in ENQuestions
	missingKeys := []string{}
	for _, question := range Questions {
		if translation, exists := ENQuestions[question.Text]; !exists {
			missingKeys = append(missingKeys, question.Text)
		} else if translation == "" {
			t.Errorf("Question key '%s' has empty translation", question.Text)
		}
	}

	if len(missingKeys) > 0 {
		t.Errorf("Found %d missing question keys in ENQuestions. First 5: %v", len(missingKeys), missingKeys[:min(5, len(missingKeys))])
	}

	// Test specific question translations
	testCases := []struct {
		key         string
		shouldExist bool
		contains    string
	}{
		{
			key:         "constructivism_becoming_woman",
			shouldExist: true,
			contains:    "One is not born",
		},
		{
			key:         "essentialism_gender_biology",
			shouldExist: true,
			contains:    "biological differences",
		},
		{
			key:         "nonexistent_key",
			shouldExist: false,
			contains:    "",
		},
	}

	for _, tc := range testCases {
		translation, exists := ENQuestions[tc.key]
		if tc.shouldExist && !exists {
			t.Errorf("Expected question key '%s' to exist in ENQuestions", tc.key)
		} else if !tc.shouldExist && exists {
			t.Errorf("Expected question key '%s' to not exist in ENQuestions", tc.key)
		} else if tc.shouldExist && exists {
			if tc.contains != "" && !contains(translation, tc.contains) {
				t.Errorf("Question '%s' translation should contain '%s', got: '%s'", tc.key, tc.contains, translation)
			}
		}
	}
}

func TestARCopyMap(t *testing.T) {
	// Arabic JSON has no top-level UI keys, so ARCopy should be empty
	expectedUIElements := 0
	if len(ARCopy) != expectedUIElements {
		t.Errorf("Expected %d UI elements in ARCopy, got %d", expectedUIElements, len(ARCopy))
	}
}

func TestARQuestionsMap(t *testing.T) {
	expectedQuestionCount := 117
	if len(ARQuestions) != expectedQuestionCount {
		t.Errorf("Expected %d questions in ARQuestions, got %d", expectedQuestionCount, len(ARQuestions))
	}

	missingKeys := []string{}
	for _, question := range Questions {
		if _, exists := ARQuestions[question.Text]; !exists {
			missingKeys = append(missingKeys, question.Text)
		}
	}
	if len(missingKeys) > 0 {
		t.Errorf("Found %d missing question keys in ARQuestions. First 5: %v", len(missingKeys), missingKeys[:min(5, len(missingKeys))])
	}
}

func TestESCopyMap(t *testing.T) {
	// Spanish JSON has no top-level UI keys, so ESCopy should be empty
	expectedUIElements := 0
	if len(ESCopy) != expectedUIElements {
		t.Errorf("Expected %d UI elements in ESCopy, got %d", expectedUIElements, len(ESCopy))
	}
}

func TestESQuestionsMap(t *testing.T) {
	expectedQuestionCount := 117
	if len(ESQuestions) != expectedQuestionCount {
		t.Errorf("Expected %d questions in ESQuestions, got %d", expectedQuestionCount, len(ESQuestions))
	}

	missingKeys := []string{}
	for _, question := range Questions {
		if _, exists := ESQuestions[question.Text]; !exists {
			missingKeys = append(missingKeys, question.Text)
		}
	}
	if len(missingKeys) > 0 {
		t.Errorf("Found %d missing question keys in ESQuestions. First 5: %v", len(missingKeys), missingKeys[:min(5, len(missingKeys))])
	}
}

func TestFRCopyMap(t *testing.T) {
	// French JSON has no top-level UI keys, so FRCopy should be empty
	expectedUIElements := 0
	if len(FRCopy) != expectedUIElements {
		t.Errorf("Expected %d UI elements in FRCopy, got %d", expectedUIElements, len(FRCopy))
	}
}

func TestFRQuestionsMap(t *testing.T) {
	expectedQuestionCount := 117
	if len(FRQuestions) != expectedQuestionCount {
		t.Errorf("Expected %d questions in FRQuestions, got %d", expectedQuestionCount, len(FRQuestions))
	}

	missingKeys := []string{}
	for _, question := range Questions {
		if _, exists := FRQuestions[question.Text]; !exists {
			missingKeys = append(missingKeys, question.Text)
		}
	}
	if len(missingKeys) > 0 {
		t.Errorf("Found %d missing question keys in FRQuestions. First 5: %v", len(missingKeys), missingKeys[:min(5, len(missingKeys))])
	}
}

func TestITCopyMap(t *testing.T) {
	// Italian JSON has no top-level UI keys, so ITCopy should be empty
	expectedUIElements := 0
	if len(ITCopy) != expectedUIElements {
		t.Errorf("Expected %d UI elements in ITCopy, got %d", expectedUIElements, len(ITCopy))
	}
}

func TestITQuestionsMap(t *testing.T) {
	expectedQuestionCount := 117
	if len(ITQuestions) != expectedQuestionCount {
		t.Errorf("Expected %d questions in ITQuestions, got %d", expectedQuestionCount, len(ITQuestions))
	}

	missingKeys := []string{}
	for _, question := range Questions {
		if _, exists := ITQuestions[question.Text]; !exists {
			missingKeys = append(missingKeys, question.Text)
		}
	}
	if len(missingKeys) > 0 {
		t.Errorf("Found %d missing question keys in ITQuestions. First 5: %v", len(missingKeys), missingKeys[:min(5, len(missingKeys))])
	}
}

func TestRUCopyMap(t *testing.T) {
	// Russian JSON has no top-level UI keys, so RUCopy should be empty
	expectedUIElements := 0
	if len(RUCopy) != expectedUIElements {
		t.Errorf("Expected %d UI elements in RUCopy, got %d", expectedUIElements, len(RUCopy))
	}
}

func TestRUQuestionsMap(t *testing.T) {
	expectedQuestionCount := 117
	if len(RUQuestions) != expectedQuestionCount {
		t.Errorf("Expected %d questions in RUQuestions, got %d", expectedQuestionCount, len(RUQuestions))
	}

	missingKeys := []string{}
	for _, question := range Questions {
		if _, exists := RUQuestions[question.Text]; !exists {
			missingKeys = append(missingKeys, question.Text)
		}
	}
	if len(missingKeys) > 0 {
		t.Errorf("Found %d missing question keys in RUQuestions. First 5: %v", len(missingKeys), missingKeys[:min(5, len(missingKeys))])
	}
}

func TestZHCopyMap(t *testing.T) {
	// Chinese JSON has no top-level UI keys, so ZHCopy should be empty
	expectedUIElements := 0
	if len(ZHCopy) != expectedUIElements {
		t.Errorf("Expected %d UI elements in ZHCopy, got %d", expectedUIElements, len(ZHCopy))
	}
}

func TestZHQuestionsMap(t *testing.T) {
	expectedQuestionCount := 117
	if len(ZHQuestions) != expectedQuestionCount {
		t.Errorf("Expected %d questions in ZHQuestions, got %d", expectedQuestionCount, len(ZHQuestions))
	}

	missingKeys := []string{}
	for _, question := range Questions {
		if _, exists := ZHQuestions[question.Text]; !exists {
			missingKeys = append(missingKeys, question.Text)
		}
	}
	if len(missingKeys) > 0 {
		t.Errorf("Found %d missing question keys in ZHQuestions. First 5: %v", len(missingKeys), missingKeys[:min(5, len(missingKeys))])
	}
}

func contains(s, substr string) bool {
	return strings.Contains(s, substr)
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
