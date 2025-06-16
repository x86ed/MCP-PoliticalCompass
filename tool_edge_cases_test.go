package main

import (
	"strings"
	"testing"

	"github.com/x86ed/MCP-PoliticalCompass/v2/eightvalues"
	politicalcompass "github.com/x86ed/MCP-PoliticalCompass/v2/political-compass"
	"github.com/x86ed/MCP-PoliticalCompass/v2/politiscales"
)

// Test handleQuizStatus edge cases to improve coverage
func TestHandleQuizStatusEdgeCases(t *testing.T) {
	// Test with no responses
	t.Run("No responses", func(t *testing.T) {
		resetState()

		response, err := handleQuizStatus(QuizStatusArgs{})
		if err != nil {
			t.Fatalf("Expected no error, got: %v", err)
		}

		content := response.Content[0].TextContent.Text
		if !strings.Contains(content, "Questions answered: 0") {
			t.Error("Expected 0 questions answered")
		}
		if !strings.Contains(content, "No questions answered yet") {
			t.Error("Expected no questions message")
		}
	})

	// Test with partial responses and all response types
	t.Run("Partial responses with all types", func(t *testing.T) {
		resetState()

		// Manually add responses to test distribution
		quizState.Responses = []politicalcompass.Response{
			politicalcompass.StronglyDisagree,
			politicalcompass.Disagree,
			politicalcompass.Agree,
			politicalcompass.StronglyAgree,
		}

		response, err := handleQuizStatus(QuizStatusArgs{})
		if err != nil {
			t.Fatalf("Expected no error, got: %v", err)
		}

		content := response.Content[0].TextContent.Text
		if !strings.Contains(content, "Questions answered: 4") {
			t.Error("Expected 4 questions answered")
		}

		// Check response distribution
		expectedDistribution := []string{
			"Strongly Disagree: 1 (25.0%)",
			"Disagree: 1 (25.0%)",
			"Agree: 1 (25.0%)",
			"Strongly Agree: 1 (25.0%)",
		}

		for _, expected := range expectedDistribution {
			if !strings.Contains(content, expected) {
				t.Errorf("Expected %s in response distribution", expected)
			}
		}
	})

	// Test with completed quiz
	t.Run("Completed quiz", func(t *testing.T) {
		resetState()

		// Simulate completed quiz by ensuring we have responses for all questions
		totalEconomicScore = 5.0
		totalSocialScore = -3.0
		questionCount = len(politicalcompass.AllQuestions)
		currentIndex = len(politicalcompass.AllQuestions)

		// Initialize questions to get proper shuffled order
		initializeQuestions()

		// Add ALL responses (this is what determines completion)
		quizState.Responses = make([]politicalcompass.Response, len(politicalcompass.AllQuestions))
		for i := 0; i < len(politicalcompass.AllQuestions); i++ {
			quizState.Responses[i] = politicalcompass.Agree
		}

		response, err := handleQuizStatus(QuizStatusArgs{})
		if err != nil {
			t.Fatalf("Expected no error, got: %v", err)
		}

		content := response.Content[0].TextContent.Text
		if !strings.Contains(content, "Questions remaining: 0") {
			t.Errorf("Expected 0 remaining questions. Content: %s", content)
		}
		if !strings.Contains(content, "Final Scores:") {
			t.Error("Expected final scores section")
		}
		if !strings.Contains(content, "Your Quadrant:") {
			t.Error("Expected quadrant information")
		}
		if !strings.Contains(content, "Quiz complete!") {
			t.Error("Expected quiz complete message")
		}
	})
}

// Test handleEightValuesStatus edge cases to improve coverage
func TestHandleEightValuesStatusEdgeCases(t *testing.T) {
	// Test with no responses
	t.Run("No responses", func(t *testing.T) {
		resetState()

		response, err := handleEightValuesStatus(EightValuesStatusArgs{})
		if err != nil {
			t.Fatalf("Expected no error, got: %v", err)
		}

		content := response.Content[0].TextContent.Text
		if !strings.Contains(content, "Questions answered: 0") {
			t.Error("Expected 0 questions answered")
		}
		if !strings.Contains(content, "No questions answered yet") {
			t.Error("Expected no questions message")
		}
	})

	// Test with partial responses and all response types
	t.Run("Partial responses with all types", func(t *testing.T) {
		resetState()

		// Manually add responses to test distribution
		eightValuesQuizState.Responses = []float64{
			eightvalues.StronglyDisagree,
			eightvalues.Disagree,
			eightvalues.Neutral,
			eightvalues.Agree,
			eightvalues.StronglyAgree,
		}

		response, err := handleEightValuesStatus(EightValuesStatusArgs{})
		if err != nil {
			t.Fatalf("Expected no error, got: %v", err)
		}

		content := response.Content[0].TextContent.Text
		if !strings.Contains(content, "Questions answered: 5") {
			t.Error("Expected 5 questions answered")
		}

		// Check response distribution
		expectedDistribution := []string{
			"Strongly Disagree: 1 (20.0%)",
			"Disagree: 1 (20.0%)",
			"Neutral: 1 (20.0%)",
			"Agree: 1 (20.0%)",
			"Strongly Agree: 1 (20.0%)",
		}

		for _, expected := range expectedDistribution {
			if !strings.Contains(content, expected) {
				t.Errorf("Expected %s in response distribution", expected)
			}
		}
	})

	// Test with completed quiz showing all axis results
	t.Run("Completed quiz with scores", func(t *testing.T) {
		resetState()

		// Simulate completed quiz with specific scores
		eightValuesEconScore = 30.0  // More socialist
		eightValuesDiplScore = -20.0 // More nationalist
		eightValuesGovtScore = 15.0  // More libertarian
		eightValuesSctyScore = -25.0 // More traditional
		eightValuesQuestionCount = len(eightvalues.Questions)
		eightValuesCurrentIndex = len(eightvalues.Questions)

		// Initialize questions
		initializeEightValuesQuestions()

		// Add ALL responses (this is what determines completion)
		eightValuesQuizState.Responses = make([]float64, len(eightvalues.Questions))
		for i := 0; i < len(eightvalues.Questions); i++ {
			eightValuesQuizState.Responses[i] = eightvalues.Neutral
		}

		response, err := handleEightValuesStatus(EightValuesStatusArgs{})
		if err != nil {
			t.Fatalf("Expected no error, got: %v", err)
		}

		content := response.Content[0].TextContent.Text
		if !strings.Contains(content, "Questions remaining: 0") {
			t.Errorf("Expected 0 remaining questions. Content: %s", content)
		}
		if !strings.Contains(content, "Final Scores:") {
			t.Error("Expected final scores section")
		}
		if !strings.Contains(content, "Economic Axis:") {
			t.Error("Expected economic axis results")
		}
		if !strings.Contains(content, "Diplomatic Axis:") {
			t.Error("Expected diplomatic axis results")
		}
		if !strings.Contains(content, "Government Axis:") {
			t.Error("Expected government axis results")
		}
		if !strings.Contains(content, "Society Axis:") {
			t.Error("Expected society axis results")
		}
		if !strings.Contains(content, "Quiz complete!") {
			t.Error("Expected quiz complete message")
		}
	})

	// Test extreme scores for boundary conditions
	t.Run("Extreme scores", func(t *testing.T) {
		resetState()

		// Set extreme scores to test boundary conditions
		eightValuesEconScore = 80.0  // Maximum socialist
		eightValuesDiplScore = -80.0 // Maximum nationalist
		eightValuesGovtScore = 80.0  // Maximum libertarian
		eightValuesSctyScore = -80.0 // Maximum traditional
		eightValuesQuestionCount = len(eightvalues.Questions)
		eightValuesCurrentIndex = len(eightvalues.Questions)

		// Initialize questions
		initializeEightValuesQuestions()

		// Add responses
		eightValuesQuizState.Responses = make([]float64, len(eightvalues.Questions))
		for i := 0; i < len(eightvalues.Questions); i++ {
			eightValuesQuizState.Responses[i] = eightvalues.StronglyAgree
		}

		response, err := handleEightValuesStatus(EightValuesStatusArgs{})
		if err != nil {
			t.Fatalf("Expected no error, got: %v", err)
		}

		content := response.Content[0].TextContent.Text

		// Should show extreme classifications
		if !strings.Contains(content, "100.0%") {
			t.Error("Expected 100% scores for extreme values")
		}
	})

	// Test with only some response types to check partial distributions
	t.Run("Partial response types", func(t *testing.T) {
		resetState()

		// Add only some response types
		eightValuesQuizState.Responses = []float64{
			eightvalues.StronglyAgree,
			eightvalues.StronglyAgree,
			eightvalues.Agree,
		}

		response, err := handleEightValuesStatus(EightValuesStatusArgs{})
		if err != nil {
			t.Fatalf("Expected no error, got: %v", err)
		}

		content := response.Content[0].TextContent.Text
		if !strings.Contains(content, "Strongly Agree: 2 (66.7%)") {
			t.Error("Expected 2 Strongly Agree responses at 66.7%")
		}
		if !strings.Contains(content, "Agree: 1 (33.3%)") {
			t.Error("Expected 1 Agree response at 33.3%")
		}

		// Should not contain responses that weren't given
		if strings.Contains(content, "Strongly Disagree:") {
			t.Error("Should not show Strongly Disagree in distribution")
		}
		if strings.Contains(content, "Disagree:") {
			t.Error("Should not show Disagree in distribution")
		}
		if strings.Contains(content, "Neutral:") {
			t.Error("Should not show Neutral in distribution")
		}
	})
}

// Test SVG generation edge cases
func TestGeneratePoliticalCompassSVGEdgeCases(t *testing.T) {
	testCases := []struct {
		name          string
		economic      float64
		social        float64
		shouldContain []string
	}{
		{
			name:          "Center position",
			economic:      0.0,
			social:        0.0,
			shouldContain: []string{"<svg", "width=", "height=", "<circle"},
		},
		{
			name:          "Extreme top-left",
			economic:      -10.0,
			social:        10.0,
			shouldContain: []string{"<svg", "Libertarian Left", "<circle"},
		},
		{
			name:          "Extreme bottom-right",
			economic:      10.0,
			social:        -10.0,
			shouldContain: []string{"<svg", "Authoritarian Right", "<circle"},
		},
		{
			name:          "Beyond bounds",
			economic:      15.0,  // Beyond normal range
			social:        -15.0, // Beyond normal range
			shouldContain: []string{"<svg", "<circle"},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			svg := generatePoliticalCompassSVG(tc.economic, tc.social)

			if svg == "" {
				t.Error("Expected non-empty SVG")
			}

			for _, expected := range tc.shouldContain {
				if !strings.Contains(svg, expected) {
					t.Errorf("Expected SVG to contain %s", expected)
				}
			}

			// Verify it's valid SVG structure
			if !strings.HasPrefix(svg, "<svg") {
				t.Error("SVG should start with <svg tag")
			}
			if !strings.HasSuffix(svg, "</svg>") {
				t.Error("SVG should end with </svg> tag")
			}
		})
	}
}

// Test initialization functions
func TestInitializationFunctions(t *testing.T) {
	// Test political compass question initialization
	t.Run("Political compass initialization", func(t *testing.T) {
		resetState()
		shuffledQuestions = nil // Reset to test initialization

		initializeQuestions()

		if len(shuffledQuestions) != len(politicalcompass.AllQuestions) {
			t.Errorf("Expected %d questions, got %d", len(politicalcompass.AllQuestions), len(shuffledQuestions))
		}

		// Verify all question indices are present
		questionMap := make(map[int]bool)
		for _, idx := range shuffledQuestions {
			questionMap[idx] = true
		}

		for i := 0; i < len(politicalcompass.AllQuestions); i++ {
			if !questionMap[i] {
				t.Errorf("Question index %d missing from shuffled questions", i)
			}
		}
	})

	// Test 8values question initialization
	t.Run("8values initialization", func(t *testing.T) {
		resetState()
		eightValuesShuffledQuestions = nil // Reset to test initialization

		initializeEightValuesQuestions()

		if len(eightValuesShuffledQuestions) != len(eightvalues.Questions) {
			t.Errorf("Expected %d questions, got %d", len(eightvalues.Questions), len(eightValuesShuffledQuestions))
		}

		// Verify all question indices are present
		questionMap := make(map[int]bool)
		for _, idx := range eightValuesShuffledQuestions {
			questionMap[idx] = true
		}

		for i := 0; i < len(eightvalues.Questions); i++ {
			if !questionMap[i] {
				t.Errorf("Question index %d missing from shuffled questions", i)
			}
		}
	})

	// Test politiscales question initialization
	t.Run("Politiscales initialization", func(t *testing.T) {
		resetState()
		politiscalesShuffledQuestions = nil // Reset to test initialization

		initializePolitiscalesQuestions()

		if len(politiscalesShuffledQuestions) != len(politiscales.Questions) {
			t.Errorf("Expected %d questions, got %d", len(politiscales.Questions), len(politiscalesShuffledQuestions))
		}

		// Verify all question indices are present
		questionMap := make(map[int]bool)
		for _, idx := range politiscalesShuffledQuestions {
			questionMap[idx] = true
		}

		for i := 0; i < len(politiscales.Questions); i++ {
			if !questionMap[i] {
				t.Errorf("Question index %d missing from shuffled questions", i)
			}
		}
	})
}

// Test multiple consecutive initializations (should not re-shuffle)
func TestInitializationIdempotency(t *testing.T) {
	t.Run("Political compass idempotency", func(t *testing.T) {
		resetState()
		initializeQuestions()
		firstShuffle := make([]int, len(shuffledQuestions))
		copy(firstShuffle, shuffledQuestions)

		initializeQuestions() // Second call

		// Should be identical (no re-shuffle)
		if len(shuffledQuestions) != len(firstShuffle) {
			t.Error("Shuffled questions length changed on second initialization")
		}

		for i, val := range shuffledQuestions {
			if i < len(firstShuffle) && val != firstShuffle[i] {
				t.Error("Questions were re-shuffled on second initialization")
				break
			}
		}
	})
}
