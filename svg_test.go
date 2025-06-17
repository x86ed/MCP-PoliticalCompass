package main

import (
	"fmt"
	"strings"
	"testing"
)

func TestSVGGeneration(t *testing.T) {
	// Test SVG generation with different positions
	testCases := []struct {
		name     string
		economic float64
		social   float64
	}{
		{"Center", 0.0, 0.0},
		{"Libertarian Left", 2.5, 3.0},
		{"Authoritarian Right", -2.0, -1.5},
		{"Extreme Libertarian Left", 5.0, 5.0},
		{"Extreme Authoritarian Right", -5.0, -5.0},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			svg := generatePoliticalCompassSVG(tc.economic, tc.social)

			// Basic validation
			if len(svg) == 0 {
				t.Error("SVG should not be empty")
			}

			if !strings.Contains(svg, "<svg") {
				t.Error("SVG should contain opening svg tag")
			}

			if !strings.Contains(svg, "</svg>") {
				t.Error("SVG should contain closing svg tag")
			}

			if !strings.Contains(svg, fmt.Sprintf("Position: (%.2f, %.2f)", tc.economic, tc.social)) {
				t.Error("SVG should contain position coordinates")
			}

			// Test that the SVG is properly formatted
			if !strings.Contains(svg, "xmlns=\"http://www.w3.org/2000/svg\"") {
				t.Error("SVG should have proper namespace")
			}

			// Debug output commented out to reduce test noise
			// fmt.Printf("Generated SVG for %s (%.2f, %.2f):\n%s\n\n", tc.name, tc.economic, tc.social, svg)
		})
	}
}
