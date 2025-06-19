package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/mark3labs/mcp-go/mcp"
	"github.com/mark3labs/mcp-go/server"
)

// Version is set during build time via ldflags
var Version = "dev"

// setupServer creates and configures an MCP server with all tools registered
func setupServer() *server.MCPServer {
	// Create a new server
	s := server.NewMCPServer(
		"Political Compass MCP Server",
		Version,
		server.WithToolCapabilities(true),
	)

	// Register political compass question tool
	politicalCompassTool := mcp.NewTool("political_compass",
		mcp.WithDescription("Presents a political compass question and accepts a response"),
		mcp.WithString("answer", mcp.Required(), mcp.Enum("strongly_agree", "agree", "neutral", "disagree", "strongly_disagree"), mcp.Description("The user's response to the question")),
	)
	s.AddTool(politicalCompassTool, handlePoliticalCompass)

	// Register reset quiz tool
	resetQuizTool := mcp.NewTool("reset_quiz",
		mcp.WithDescription("Resets the political compass quiz progress"),
	)
	s.AddTool(resetQuizTool, handleResetQuiz)

	// Register quiz status tool
	quizStatusTool := mcp.NewTool("quiz_status",
		mcp.WithDescription("Shows current quiz progress and statistics"),
	)
	s.AddTool(quizStatusTool, handleQuizStatus)

	// Register 8values quiz tool
	eightValuesTool := mcp.NewTool("eight_values",
		mcp.WithDescription("Presents an 8values political question and accepts a response"),
		mcp.WithString("answer", mcp.Required(), mcp.Enum("strongly_agree", "agree", "neutral", "disagree", "strongly_disagree"), mcp.Description("The user's response to the question")),
	)
	s.AddTool(eightValuesTool, handleEightValues)

	// Register reset 8values quiz tool
	resetEightValuesTool := mcp.NewTool("reset_eight_values",
		mcp.WithDescription("Resets the 8values quiz progress"),
	)
	s.AddTool(resetEightValuesTool, handleResetEightValues)

	// Register 8values quiz status tool
	eightValuesStatusTool := mcp.NewTool("eight_values_status",
		mcp.WithDescription("Shows current 8values quiz progress and statistics"),
	)
	s.AddTool(eightValuesStatusTool, handleEightValuesStatus)

	// Register politiscales quiz tool
	politiscalesTool := mcp.NewTool("politiscales",
		mcp.WithDescription("Presents a politiscales political question and accepts a response"),
		mcp.WithString("answer", mcp.Required(), mcp.Enum("strongly_agree", "agree", "neutral", "disagree", "strongly_disagree"), mcp.Description("The user's response to the question")),
	)
	s.AddTool(politiscalesTool, handlePolitiscales)

	// Register reset politiscales quiz tool
	resetPolitiscalesTool := mcp.NewTool("reset_politiscales",
		mcp.WithDescription("Resets the politiscales quiz progress"),
	)
	s.AddTool(resetPolitiscalesTool, handleResetPolitiscales)

	// Register politiscales quiz status tool
	politiscalesStatusTool := mcp.NewTool("politiscales_status",
		mcp.WithDescription("Shows current politiscales quiz progress and statistics"),
	)
	s.AddTool(politiscalesStatusTool, handlePolitiscalesStatus)

	// Register set politiscales language tool
	setPolitiscalesLanguageTool := mcp.NewTool("set_politiscales_language",
		mcp.WithDescription("Sets the language for the politiscales quiz"),
		mcp.WithString("language", mcp.Required(), mcp.Enum("en", "fr", "es", "it", "ru", "zh", "ar"), mcp.Description("The language code for the quiz")),
	)
	s.AddTool(setPolitiscalesLanguageTool, handleSetPolitiscalesLanguage)

	return s
}

func main() {
	showVersion := flag.Bool("version", false, "Show version")
	flag.Parse()

	if *showVersion {
		fmt.Println("Version:", Version)
		os.Exit(0)
	}

	s := setupServer()

	if err := server.ServeStdio(s); err != nil {
		fmt.Fprintf(os.Stderr, "Error serving: %v\n", err)
		os.Exit(1)
	}
}
