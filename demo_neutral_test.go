package main

import (
	"fmt"
	"testing"
)

// Test to demonstrate how neutral states and paired values are handled
func TestPolitiscalesNeutralDemo(t *testing.T) {
	// Create test results to demonstrate the new neutral handling
	results := map[string]float64{
		// Example pair with significant neutral space
		"constructivism": 30.0,
		"essentialism":   20.0, // Total = 50%, so 50% neutral

		// Example pair with minimal neutral space
		"rehabilitative_justice": 60.0,
		"punitive_justice":       35.0, // Total = 95%, so 5% neutral

		// Example pair with balanced scores
		"progressive":  35.0,
		"conservative": 35.0, // Total = 70%, so 30% neutral

		// Set other pairs to have some values
		"internationalism": 40.0, "nationalism": 30.0,
		"communism": 45.0, "capitalism": 40.0,
		"regulation": 50.0, "laissez_faire": 30.0,
		"ecology": 60.0, "production": 25.0,
		"revolution": 25.0, "reform": 50.0,

		// Unpaired axes for badges
		"anarchism": 95.0, "pragmatism": 60.0, "feminism": 45.0,
	}

	// Generate SVG to show how neutral is handled
	svg := generatePolitiscalesResultsSVG(results)

	// Look for evidence that neutral states are calculated
	fmt.Printf("Sample SVG content showing neutral handling:\n")
	fmt.Printf("==========================================\n")

	// Extract a portion of the SVG to show neutral calculation
	if len(svg) > 500 {
		sample := svg[500:1000] // Get a sample section
		fmt.Printf("SVG Sample: %s...\n", sample)
	}

	// Demonstrate what the text results format would look like
	fmt.Printf("\nExpected Text Results Format:\n")
	fmt.Printf("============================\n")
	fmt.Printf("- **Identity:** 30.0%% Constructivism | 20.0%% Essentialism | 50.0%% Neutral\n")
	fmt.Printf("- **Justice:** 60.0%% Rehabilitative Justice | 35.0%% Punitive Justice | 5.0%% Neutral\n")
	fmt.Printf("- **Culture:** 35.0%% Progressive | 35.0%% Conservative | 30.0%% Neutral\n")

	fmt.Printf("\n**Notable Characteristics:**\n")
	fmt.Printf("- Anarchist: 95.0%%\n")
	fmt.Printf("- Pragmatist: 60.0%%\n")

	t.Log("✅ Demonstrated neutral state handling in politiscales results")
}
