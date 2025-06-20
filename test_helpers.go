package main

import (
	"github.com/mark3labs/mcp-go/mcp"
)

// Helper functions for testing with the new MCP library

// createMockRequest creates a mock CallToolRequest for testing
func createMockRequest(toolName string, args map[string]interface{}) mcp.CallToolRequest {
	return mcp.CallToolRequest{
		Params: mcp.CallToolParams{
			Name:      toolName,
			Arguments: args,
		},
	}
}

// createRequestWithAnswer creates a request with an "answer" argument
func createRequestWithAnswer(answer string) mcp.CallToolRequest {
	return createMockRequest("test_tool", map[string]interface{}{
		"answer": answer,
	})
}

// createRequestWithLanguage creates a request with a "language" argument
func createRequestWithLanguage(language string) mcp.CallToolRequest {
	return createMockRequest("test_tool", map[string]interface{}{
		"language": language,
	})
}

// createEmptyRequest creates a request with no arguments
func createEmptyRequest() mcp.CallToolRequest {
	return createMockRequest("test_tool", map[string]interface{}{})
}

// extractTextContent extracts text content from a CallToolResult
func extractTextContent(result *mcp.CallToolResult) string {
	if result == nil || len(result.Content) == 0 {
		return ""
	}

	content := result.Content[0]
	if textContent, ok := mcp.AsTextContent(content); ok {
		return textContent.Text
	}

	return ""
}

// isErrorResult checks if a CallToolResult represents an error
func isErrorResult(result *mcp.CallToolResult) bool {
	return result != nil && result.IsError
}
