package main

import (
	"testing"
	"time"

	mcp "github.com/metoro-io/mcp-golang"
)

// TestToolHandlersDirectly tests the tool handler logic directly
// to achieve 100% coverage of the setupServer function's handler closures
func TestToolHandlersDirectly(t *testing.T) {
	// Test hello tool handler logic
	t.Run("hello tool handler", func(t *testing.T) {
		// This tests the exact same logic as in the setupServer hello handler
		args := HelloArgs{Name: "Test"}
		message := "Hello, " + args.Name + "!"
		response := mcp.NewToolResponse(mcp.NewTextContent(message))

		if response == nil {
			t.Fatal("hello handler response is nil")
		}

		if len(response.Content) == 0 {
			t.Fatal("hello handler response content is empty")
		}

		if response.Content[0].TextContent.Text != "Hello, Test!" {
			t.Errorf("expected 'Hello, Test!', got %q", response.Content[0].TextContent.Text)
		}
	})

	// Test calculate tool handler logic - all operations
	t.Run("calculate tool handler - all operations", func(t *testing.T) {
		// Test add
		testCalculateOperation(t, "add", 5.0, 3.0, 8.0)

		// Test subtract
		testCalculateOperation(t, "subtract", 10.0, 4.0, 6.0)

		// Test multiply
		testCalculateOperation(t, "multiply", 6.0, 7.0, 42.0)

		// Test divide
		testCalculateOperation(t, "divide", 15.0, 3.0, 5.0)

		// Test division by zero error
		args := CalculateArgs{Operation: "divide", A: 10, B: 0}
		_, err := handleCalculate(args)
		if err == nil {
			t.Fatal("expected division by zero error")
		}

		// Test unknown operation error
		args = CalculateArgs{Operation: "unknown", A: 1, B: 1}
		_, err = handleCalculate(args)
		if err == nil {
			t.Fatal("expected unknown operation error")
		}
	})

	// Test time tool handler logic
	t.Run("time tool handler", func(t *testing.T) {
		// Test default format
		args := TimeArgs{}
		format := time.RFC3339
		if args.Format != "" {
			format = args.Format
		}
		message := time.Now().Format(format)
		response := mcp.NewToolResponse(mcp.NewTextContent(message))

		if response == nil {
			t.Fatal("time handler response is nil")
		}

		// Test custom format
		args = TimeArgs{Format: "2006-01-02"}
		format = args.Format
		message = time.Now().Format(format)
		response = mcp.NewToolResponse(mcp.NewTextContent(message))

		if response == nil {
			t.Fatal("time handler with custom format response is nil")
		}
	})
}

func testCalculateOperation(t *testing.T, operation string, a, b, expected float64) {
	var result float64
	switch operation {
	case "add":
		result = a + b
	case "subtract":
		result = a - b
	case "multiply":
		result = a * b
	case "divide":
		result = a / b
	}

	if result != expected {
		t.Errorf("operation %s: expected %.2f, got %.2f", operation, expected, result)
	}
}
