package main

import (
	"testing"

	"github.com/mark3labs/mcp-go/server"
)

// Test that covers the version flag handling path in main (indirectly)
func TestVersionHandling(t *testing.T) {
	// We can't easily test main() directly due to os.Exit(), but we can test
	// that the version variable is properly set
	if Version == "" {
		t.Error("Version should not be empty")
	}
}

// Test server creation
func TestServerCreation(t *testing.T) {
	s := server.NewMCPServer(
		"Test Server",
		"1.0.0",
		server.WithToolCapabilities(true),
	)
	if s == nil {
		t.Fatal("Expected server to be created, got nil")
	}
}

// Test that setupServer works correctly
func TestSetupServer(t *testing.T) {
	server := setupServer()
	if server == nil {
		t.Fatal("Expected server to be created, got nil")
	}
}
