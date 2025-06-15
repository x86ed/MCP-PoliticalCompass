package main

import (
	"flag"
	"fmt"
	"os"

	mcp "github.com/metoro-io/mcp-golang"
	"github.com/metoro-io/mcp-golang/transport"
	"github.com/metoro-io/mcp-golang/transport/stdio"
)

// Version is set during build time via ldflags
var Version = "dev"

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

	// Register quiz status tool
	err = server.RegisterTool("quiz_status", "Shows current quiz progress and statistics", handleQuizStatus)
	if err != nil {
		return nil, err
	}

	// Register 8values quiz tool
	err = server.RegisterTool("eight_values", "Presents an 8values political question and accepts a response", handleEightValues)
	if err != nil {
		return nil, err
	}

	// Register reset 8values quiz tool
	err = server.RegisterTool("reset_eight_values", "Resets the 8values quiz progress", handleResetEightValues)
	if err != nil {
		return nil, err
	}

	// Register 8values quiz status tool
	err = server.RegisterTool("eight_values_status", "Shows current 8values quiz progress and statistics", handleEightValuesStatus)
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
	// Command-line flags
	showVersion := flag.Bool("version", false, "Show version")
	// Add more flags as needed

	// Parse command-line flags
	flag.Parse()

	// Handle version flag
	if *showVersion {
		fmt.Println("Version:", Version)
		os.Exit(0)
	}

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
