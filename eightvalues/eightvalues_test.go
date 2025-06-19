package eightvalues

import (
	"strings"
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

func TestGenerateSVG(t *testing.T) {
	// Test basic SVG generation
	svg := GenerateSVG(50.0, 50.0, 50.0, 50.0)

	// Check that it returns a valid SVG
	if !strings.HasPrefix(svg, "<svg") {
		t.Error("Generated output should start with <svg")
	}
	if !strings.HasSuffix(svg, "</svg>") {
		t.Error("Generated output should end with </svg>")
	}

	// Check for required SVG attributes
	if !strings.Contains(svg, `width="800"`) {
		t.Error("SVG should have width=800")
	}
	if !strings.Contains(svg, `height="650"`) {
		t.Error("SVG should have height=650")
	}

	// Check for 8values title
	if !strings.Contains(svg, "8values") {
		t.Error("SVG should contain 8values title")
	}

	// Check for axis labels
	if !strings.Contains(svg, "Economic Axis:") {
		t.Error("SVG should contain Economic Axis label")
	}
	if !strings.Contains(svg, "Diplomatic Axis:") {
		t.Error("SVG should contain Diplomatic Axis label")
	}
	if !strings.Contains(svg, "Government Axis:") {
		t.Error("SVG should contain Government Axis label")
	}
	if !strings.Contains(svg, "Society Axis:") {
		t.Error("SVG should contain Society Axis label")
	}
}

func TestGenerateSVGWithExtremeValues(t *testing.T) {
	// Test with extreme values (0% and 100%)
	svg := GenerateSVG(100.0, 0.0, 100.0, 0.0)

	// Should still generate valid SVG
	if !strings.HasPrefix(svg, "<svg") {
		t.Error("Generated output should start with <svg")
	}
	if !strings.HasSuffix(svg, "</svg>") {
		t.Error("Generated output should end with </svg>")
	}

	// Check that percentages appear when > 30%
	if !strings.Contains(svg, "100.0%") {
		t.Error("SVG should contain 100.0% for high values")
	}
}

func TestGenerateSVGLabeling(t *testing.T) {
	// Test extreme values to check labeling
	tests := []struct {
		name     string
		econ     float64
		dipl     float64
		govt     float64
		scty     float64
		expected []string
	}{
		{
			name:     "extreme equality",
			econ:     100.0,
			dipl:     50.0,
			govt:     50.0,
			scty:     50.0,
			expected: []string{"Communist"},
		},
		{
			name:     "extreme wealth",
			econ:     0.0,
			dipl:     50.0,
			govt:     50.0,
			scty:     50.0,
			expected: []string{"Laissez-Faire"},
		},
		{
			name:     "balanced centrist",
			econ:     50.0,
			dipl:     50.0,
			govt:     50.0,
			scty:     50.0,
			expected: []string{"Centrist", "Balanced", "Moderate", "Neutral"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			svg := GenerateSVG(tt.econ, tt.dipl, tt.govt, tt.scty)

			for _, expectedLabel := range tt.expected {
				if !strings.Contains(svg, expectedLabel) {
					t.Errorf("SVG should contain label '%s' for test '%s'", expectedLabel, tt.name)
				}
			}
		})
	}
}

func TestGenerateSVGPercentageDisplay(t *testing.T) {
	// Test that percentages only appear when > 30%
	svg := GenerateSVG(25.0, 35.0, 25.0, 35.0)

	// Should not contain 25.0% (below threshold)
	if strings.Contains(svg, "25.0%") {
		t.Error("SVG should not contain percentages below 30%")
	}

	// Should contain 35.0% (above threshold)
	if !strings.Contains(svg, "35.0%") {
		t.Error("SVG should contain percentages above 30%")
	}
}
