// Package politicalcompass provides types and data for political compass questionnaires.
package politicalcompass

import "fmt"

// Response represents the possible responses to a political compass question
type Response int

const (
	StronglyDisagree Response = iota
	Disagree
	Agree
	StronglyAgree
)

// String returns the string representation of the Response
func (r Response) String() string {
	switch r {
	case StronglyDisagree:
		return "Strongly Disagree"
	case Disagree:
		return "Disagree"
	case Agree:
		return "Agree"
	case StronglyAgree:
		return "Strongly Agree"
	default:
		return "Unknown"
	}
}

// Question represents a political compass question with economic and social scoring data
type Question struct {
	Index    int32      // Question index/ID
	Economic [4]float64 // Economic scoring values (array of 4 floats)
	Social   [4]float64 // Social scoring values (array of 4 floats)
	Text     string     // Question text
}

// GenerateSVG generates an SVG visualization of political compass results
func GenerateSVG(economicScore, socialScore float64) string {
	// SVG dimensions and margins
	width := 400
	height := 400
	margin := 50

	// Calculate center point
	centerX := width / 2
	centerY := height / 2

	// Calculate user position on the graph
	// Economic score: -10 to +10 maps to left-right (0 to width-2*margin)
	// Social score: -10 to +10 maps to top-bottom (inverted: +10 is top)
	userX := centerX + int(economicScore*(float64(width-2*margin)/20.0))
	userY := centerY - int(socialScore*(float64(height-2*margin)/20.0))

	// Clamp values to stay within bounds
	if userX < margin {
		userX = margin
	}
	if userX > width-margin {
		userX = width - margin
	}
	if userY < margin {
		userY = margin
	}
	if userY > height-margin {
		userY = height - margin
	}

	svg := fmt.Sprintf(`<svg width="%d" height="%d" xmlns="http://www.w3.org/2000/svg">
  <!-- Background -->
  <rect width="%d" height="%d" fill="#f8f9fa" stroke="#dee2e6" stroke-width="1"/>
  
  <!-- Quadrant backgrounds -->
  <!-- Authoritarian Left (top-left) -->
  <rect x="%d" y="%d" width="%d" height="%d" fill="#ffebee" opacity="0.7"/>
  <!-- Authoritarian Right (top-right) -->
  <rect x="%d" y="%d" width="%d" height="%d" fill="#fff3e0" opacity="0.7"/>
  <!-- Libertarian Left (bottom-left) -->
  <rect x="%d" y="%d" width="%d" height="%d" fill="#e8f5e8" opacity="0.7"/>
  <!-- Libertarian Right (bottom-right) -->
  <rect x="%d" y="%d" width="%d" height="%d" fill="#e3f2fd" opacity="0.7"/>
  
  <!-- Grid lines -->
  <!-- Vertical center line -->
  <line x1="%d" y1="%d" x2="%d" y2="%d" stroke="#6c757d" stroke-width="2"/>
  <!-- Horizontal center line -->
  <line x1="%d" y1="%d" x2="%d" y2="%d" stroke="#6c757d" stroke-width="2"/>
  
  <!-- Grid marks -->
  <!-- Economic axis marks -->
  <line x1="%d" y1="%d" x2="%d" y2="%d" stroke="#adb5bd" stroke-width="1"/>
  <line x1="%d" y1="%d" x2="%d" y2="%d" stroke="#adb5bd" stroke-width="1"/>
  <!-- Social axis marks -->
  <line x1="%d" y1="%d" x2="%d" y2="%d" stroke="#adb5bd" stroke-width="1"/>
  <line x1="%d" y1="%d" x2="%d" y2="%d" stroke="#adb5bd" stroke-width="1"/>
  
  <!-- Axis labels -->
  <text x="%d" y="%d" text-anchor="middle" font-family="Arial, sans-serif" font-size="12" fill="#495057">Economic</text>
  <text x="%d" y="%d" text-anchor="middle" font-family="Arial, sans-serif" font-size="12" fill="#495057">Left</text>
  <text x="%d" y="%d" text-anchor="middle" font-family="Arial, sans-serif" font-size="12" fill="#495057">Right</text>
  <text x="%d" y="%d" text-anchor="middle" font-family="Arial, sans-serif" font-size="12" fill="#495057" transform="rotate(-90 %d %d)">Social</text>
  <text x="%d" y="%d" text-anchor="middle" font-family="Arial, sans-serif" font-size="12" fill="#495057" transform="rotate(-90 %d %d)">Libertarian</text>
  <text x="%d" y="%d" text-anchor="middle" font-family="Arial, sans-serif" font-size="12" fill="#495057" transform="rotate(-90 %d %d)">Authoritarian</text>
  
  <!-- Quadrant labels -->
  <text x="%d" y="%d" text-anchor="middle" font-family="Arial, sans-serif" font-size="11" font-weight="bold" fill="#666">Authoritarian</text>
  <text x="%d" y="%d" text-anchor="middle" font-family="Arial, sans-serif" font-size="11" font-weight="bold" fill="#666">Left</text>
  <text x="%d" y="%d" text-anchor="middle" font-family="Arial, sans-serif" font-size="11" font-weight="bold" fill="#666">Authoritarian</text>
  <text x="%d" y="%d" text-anchor="middle" font-family="Arial, sans-serif" font-size="11" font-weight="bold" fill="#666">Right</text>
  <text x="%d" y="%d" text-anchor="middle" font-family="Arial, sans-serif" font-size="11" font-weight="bold" fill="#666">Libertarian</text>
  <text x="%d" y="%d" text-anchor="middle" font-family="Arial, sans-serif" font-size="11" font-weight="bold" fill="#666">Left</text>
  <text x="%d" y="%d" text-anchor="middle" font-family="Arial, sans-serif" font-size="11" font-weight="bold" fill="#666">Libertarian</text>
  <text x="%d" y="%d" text-anchor="middle" font-family="Arial, sans-serif" font-size="11" font-weight="bold" fill="#666">Right</text>
  
  <!-- User position -->
  <circle cx="%d" cy="%d" r="8" fill="#dc3545" stroke="#ffffff" stroke-width="2"/>
  <text x="%d" y="%d" text-anchor="middle" font-family="Arial, sans-serif" font-size="10" font-weight="bold" fill="#ffffff">●</text>
  
  <!-- Position coordinates -->
  <text x="%d" y="%d" text-anchor="start" font-family="Arial, sans-serif" font-size="10" fill="#495057">
    Position: (%.2f, %.2f)
  </text>
</svg>`,
		width, height, // SVG dimensions
		width, height, // Background rect
		margin, margin, centerX-margin, centerY-margin, // Auth Left
		centerX, margin, centerX-margin, centerY-margin, // Auth Right
		margin, centerY, centerX-margin, centerY-margin, // Lib Left
		centerX, centerY, centerX-margin, centerY-margin, // Lib Right
		centerX, margin, centerX, height-margin, // Vertical center line
		margin, centerY, width-margin, centerY, // Horizontal center line
		margin+(centerX-margin)/2, centerY-5, margin+(centerX-margin)/2, centerY+5, // Economic left mark
		centerX+(centerX-margin)/2, centerY-5, centerX+(centerX-margin)/2, centerY+5, // Economic right mark
		centerX-5, margin+(centerY-margin)/2, centerX+5, margin+(centerY-margin)/2, // Social top mark
		centerX-5, centerY+(centerY-margin)/2, centerX+5, centerY+(centerY-margin)/2, // Social bottom mark
		centerX, height-15, // Economic label
		margin-5, centerY+15, // Left label
		width-margin+5, centerY+15, // Right label
		15, centerY, 15, centerY, // Social label
		15, margin+(centerY-margin)/2, 15, margin+(centerY-margin)/2, // Libertarian label
		15, centerY+(centerY-margin)/2, 15, centerY+(centerY-margin)/2, // Authoritarian label
		margin+(centerX-margin)/2, margin+15, // Auth Left label
		margin+(centerX-margin)/2, margin+25, // Auth Left label 2
		centerX+(centerX-margin)/2, margin+15, // Auth Right label
		centerX+(centerX-margin)/2, margin+25, // Auth Right label 2
		margin+(centerX-margin)/2, height-margin+15, // Lib Left label
		margin+(centerX-margin)/2, height-margin+25, // Lib Left label 2
		centerX+(centerX-margin)/2, height-margin+15, // Lib Right label
		centerX+(centerX-margin)/2, height-margin+25, // Lib Right label 2
		userX, userY, // User position circle
		userX, userY+1, // User position text
		margin, height-30, // Position text
		economicScore, socialScore, // Position coordinates
	)

	return svg
}
