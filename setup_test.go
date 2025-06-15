package main

import (
	"testing"

	"github.com/metoro-io/mcp-golang/transport"
	stdio "github.com/metoro-io/mcp-golang/transport/stdio"
)

// Removed unused MockTransport struct to reduce dead code.

func TestSetupServer(t *testing.T) {
	// Test successful server setup
	transport := stdio.NewStdioServerTransport()
	server, err := setupServer(transport)
	if err != nil {
		t.Fatalf("Expected no error setting up server, got: %v", err)
	}

	if server == nil {
		t.Fatal("Expected server to be created, got nil")
	}
}

func TestCreateServerTransport(t *testing.T) {
	transport := createServerTransport()
	if transport == nil {
		t.Fatal("Expected transport to be created, got nil")
	}
}

// Test that covers the version flag handling path in main (indirectly)
func TestVersionHandling(t *testing.T) {
	// We can't easily test main() directly due to os.Exit(), but we can test
	// that the version variable is properly set
	if Version == "" {
		t.Error("Version should not be empty")
	}
}
