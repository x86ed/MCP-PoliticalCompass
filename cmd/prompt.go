package main

import (
	"strings"

	mcp "github.com/metoro-io/mcp-golang"
)

// PromptArgs represents the arguments for custom prompts
type PromptArgs struct {
	Input string `json:"input" jsonschema:"required,description=The input text to process"`
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
