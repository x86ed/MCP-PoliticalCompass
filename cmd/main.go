package main

import (
	"fmt"
	"strings"
	"time"

	mcp "github.com/metoro-io/mcp-golang"
	"github.com/metoro-io/mcp-golang/transport"
	"github.com/metoro-io/mcp-golang/transport/stdio"
)

// HelloArgs represents the arguments for the hello tool
type HelloArgs struct {
	Name string `json:"name" jsonschema:"required,description=The name to say hello to"`
}

// CalculateArgs represents the arguments for the calculate tool
type CalculateArgs struct {
	Operation string  `json:"operation" jsonschema:"required,enum=add,enum=subtract,enum=multiply,enum=divide,description=The mathematical operation to perform"`
	A         float64 `json:"a" jsonschema:"required,description=First number"`
	B         float64 `json:"b" jsonschema:"required,description=Second number"`
}

// TimeArgs represents the arguments for the current time tool
type TimeArgs struct {
	Format string `json:"format,omitempty" jsonschema:"description=Optional time format (default: RFC3339)"`
}

// PromptArgs represents the arguments for custom prompts
type PromptArgs struct {
	Input string `json:"input" jsonschema:"required,description=The input text to process"`
}

// Handler functions for tools and prompts
func handleHello(args HelloArgs) (*mcp.ToolResponse, error) {
	message := fmt.Sprintf("Hello, %s!", args.Name)
	return mcp.NewToolResponse(mcp.NewTextContent(message)), nil
}

func handleCalculate(args CalculateArgs) (*mcp.ToolResponse, error) {
	var result float64
	switch args.Operation {
	case "add":
		result = args.A + args.B
	case "subtract":
		result = args.A - args.B
	case "multiply":
		result = args.A * args.B
	case "divide":
		if args.B == 0 {
			return nil, fmt.Errorf("division by zero")
		}
		result = args.A / args.B
	default:
		return nil, fmt.Errorf("unknown operation: %s", args.Operation)
	}
	message := fmt.Sprintf("Result of %s: %.2f", args.Operation, result)
	return mcp.NewToolResponse(mcp.NewTextContent(message)), nil
}

func handleTime(args TimeArgs) (*mcp.ToolResponse, error) {
	format := time.RFC3339
	if args.Format != "" {
		format = args.Format
	}
	message := time.Now().Format(format)
	return mcp.NewToolResponse(mcp.NewTextContent(message)), nil
}

func handleUppercasePrompt(args PromptArgs) (*mcp.PromptResponse, error) {
	text := strings.ToUpper(args.Input)
	return mcp.NewPromptResponse("uppercase", mcp.NewPromptMessage(mcp.NewTextContent(text), mcp.RoleUser)), nil
}

func handleReversePrompt(args PromptArgs) (*mcp.PromptResponse, error) {
	// Reverse the string
	runes := []rune(args.Input)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	text := string(runes)
	return mcp.NewPromptResponse("reverse", mcp.NewPromptMessage(mcp.NewTextContent(text), mcp.RoleUser)), nil
}

// setupServer creates and configures an MCP server with all tools and prompts registered
func setupServer(transport transport.Transport) (*mcp.Server, error) {
	// Create a new server with the transport
	server := mcp.NewServer(transport)

	// Register hello tool
	err := server.RegisterTool("hello", "Says hello to the provided name", handleHello)
	if err != nil {
		return nil, err
	}

	// Register calculate tool
	err = server.RegisterTool("calculate", "Performs basic mathematical operations", handleCalculate)
	if err != nil {
		return nil, err
	}

	// Register current time tool
	err = server.RegisterTool("time", "Returns the current time", handleTime)
	if err != nil {
		return nil, err
	}

	// Register example prompts
	err = server.RegisterPrompt("uppercase", "Converts text to uppercase", handleUppercasePrompt)
	if err != nil {
		return nil, err
	}

	err = server.RegisterPrompt("reverse", "Reverses the input text", handleReversePrompt)
	if err != nil {
		return nil, err
	}

	return server, nil
}

// createServerTransport creates the transport for the server
func createServerTransport() transport.Transport {
	return stdio.NewStdioServerTransport()
}

func main() {
	// Create a transport for the server
	serverTransport := createServerTransport()

	// Setup the server
	server, err := setupServer(serverTransport)
	if err != nil {
		panic(err)
	}

	// Start the server
	if err := server.Serve(); err != nil {
		panic(err)
	}

	// Keep the server running
	select {}
}
