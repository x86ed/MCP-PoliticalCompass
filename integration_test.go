package main

import (
	"testing"
)

func TestToolHandlersDirectly(t *testing.T) {
	t.Run("political compass tool handler", func(t *testing.T) {
		resetState()

		response, err := handlePoliticalCompass(PoliticalCompassArgs{Response: ""})
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}

		if response == nil {
			t.Fatal("political compass handler response is nil")
		}

		if len(response.Content) == 0 {
			t.Fatal("political compass handler response content is empty")
		}

		if response.Content[0].TextContent == nil {
			t.Fatal("political compass handler response content is not TextContent")
		}
	})

	t.Run("reset quiz tool handler", func(t *testing.T) {
		resetState()

		handlePoliticalCompass(PoliticalCompassArgs{Response: ""})

		response, err := handleResetQuiz(ResetQuizArgs{})
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}

		if response == nil {
			t.Fatal("reset quiz handler response is nil")
		}

		if len(response.Content) == 0 {
			t.Fatal("reset quiz handler response content is empty")
		}

		if response.Content[0].TextContent == nil {
			t.Fatal("reset quiz handler response content is not TextContent")
		}
	})
}
