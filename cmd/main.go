package main

import (
	mcp "github.com/metoro-io/mcp-golang"
	"github.com/metoro-io/mcp-golang/transport"
	"github.com/metoro-io/mcp-golang/transport/stdio"
)

// setupServer creates and configures an MCP server with all tools and prompts registered
func setupServer(transport transport.Transport) (*mcp.Server, error) {
	// Create a new server with the transport
	server := mcp.NewServer(transport)

	// Register political compass question tool
	err := server.RegisterTool("political_compass", "Presents a political compass question and accepts a response", handlePoliticalCompass)
	if err != nil {
		return nil, err
	}

	// Register reset quiz tool
	err = server.RegisterTool("reset_quiz", "Resets the political compass quiz progress", handleResetQuiz)
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
