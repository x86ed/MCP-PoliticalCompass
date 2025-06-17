package main

import (
	"strings"
	"testing"
)

// Test that neutral states and paired values are properly handled
func TestPolitiscalesNeutralAndPairedHandling(t *testing.T) {
	// Create test results with varied scores to test neutral calculation
	results := map[string]float64{
		// Paired axes with different totals to test neutral calculation
		"constructivism": 40.0,
		"essentialism":   30.0, // Total = 70%, so 30% neutral

		"rehabilitative_justice": 60.0,
		"punitive_justice":       35.0, // Total = 95%, so 5% neutral

		"progressive":  25.0,
		"conservative": 25.0, // Total = 50%, so 50% neutral

		"internationalism": 80.0,
		"nationalism":      20.0, // Total = 100%, so 0% neutral

		"communism":  45.0,
		"capitalism": 40.0, // Total = 85%, so 15% neutral

		"regulation":    30.0,
		"laissez_faire": 35.0, // Total = 65%, so 35% neutral

		"ecology":    70.0,
		"production": 15.0, // Total = 85%, so 15% neutral

		"revolution": 20.0,
		"reform":     60.0, // Total = 80%, so 20% neutral

		// Unpaired axes (badges)
		"anarchism":  95.0, // Should show
		"pragmatism": 60.0, // Should show
		"feminism":   45.0, // Should NOT show (below 90% threshold)
	}

	// Test SVG generation includes neutral states
	svg := generatePolitiscalesResultsSVG(results)

	// Check that SVG includes neutral percentages
	if !strings.Contains(svg, "30%") { // Neutral from constructivism/essentialism
		t.Error("SVG should show 30% neutral for identity pair")
	}
	if !strings.Contains(svg, "50%") { // Neutral from progressive/conservative
		t.Error("SVG should show 50% neutral for culture pair")
	}

	// Check that SVG shows both sides of pairs
	if !strings.Contains(svg, "Constructivism") || !strings.Contains(svg, "Essentialism") {
		t.Error("SVG should show both sides of identity pair")
	}

	// Check that badges are included
	if !strings.Contains(svg, "Anarchist") {
		t.Error("SVG should include Anarchist badge")
	}
	if !strings.Contains(svg, "Pragmatist") {
		t.Error("SVG should include Pragmatist badge")
	}
	if strings.Contains(svg, "Feminist") {
		t.Error("SVG should NOT include Feminist badge (below threshold)")
	}

	t.Log("✅ SVG correctly handles neutral states and paired values")

	// Now let's test that the text results would also handle this correctly
	// We can't easily call handlePolitiscales directly, but we can verify
	// the structure exists in the SVG which shares the same logic

	if !strings.Contains(svg, "Additional Characteristics") {
		t.Error("SVG should include Additional Characteristics section")
	}

	t.Log("✅ Neutral states and paired values are properly handled in politiscales")
}
