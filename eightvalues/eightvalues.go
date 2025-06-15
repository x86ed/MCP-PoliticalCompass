package eightvalues

import "fmt"

// Effect indices for the Question.Effect array
// Each constant represents an index position in the Effect array
const (
	Economic   = iota // Index 0: Economic axis effect
	Diplomatic        // Index 1: Diplomatic axis effect
	Government        // Index 2: Government axis effect
	Society           // Index 3: Society axis effect
)

// Response values for question answers
// These correspond to the button onclick values in the UI
const (
	StronglyAgree    = 1.0  // Strongly Agree response value
	Agree            = 0.5  // Agree response value
	Neutral          = 0.0  // Neutral/Unsure response value
	Disagree         = -0.5 // Disagree response value
	StronglyDisagree = -1.0 // Strongly Disagree response value
)

// Question represents a 8values question with economic and social scoring data
type Question struct {
	Index  int32      // Question index/ID
	Effect [4]float64 // Effect scoring values (array of 4 floats)
	Text   string     // Question text
}

// GenerateSVG creates an SVG visualization of the user's 8values position
func GenerateSVG(econPercentage, diplPercentage, govtPercentage, sctyPercentage float64) string {
	// SVG dimensions - match 8values.js canvas size
	width := 800
	height := 650

	// Calculate complementary percentages (following 8values.js logic)
	wealthPercentage := 100 - econPercentage
	mightPercentage := 100 - diplPercentage
	authorityPercentage := 100 - govtPercentage
	traditionPercentage := 100 - sctyPercentage

	// Helper function to get ideological label (matches 8values setLabel function)
	getLabel := func(val float64, labelArray []string) string {
		if val > 100 {
			return ""
		} else if val > 90 {
			return labelArray[0]
		} else if val > 75 {
			return labelArray[1]
		} else if val > 60 {
			return labelArray[2]
		} else if val >= 40 {
			return labelArray[3]
		} else if val >= 25 {
			return labelArray[4]
		} else if val >= 10 {
			return labelArray[5]
		} else if val >= 0 {
			return labelArray[6]
		}
		return ""
	}

	// Label arrays (from 8values.js)
	econArray := []string{"Communist", "Socialist", "Social", "Centrist", "Market", "Capitalist", "Laissez-Faire"}
	diplArray := []string{"Cosmopolitan", "Internationalist", "Peaceful", "Balanced", "Patriotic", "Nationalist", "Chauvinist"}
	govtArray := []string{"Anarchist", "Libertarian", "Liberal", "Moderate", "Statist", "Authoritarian", "Totalitarian"}
	sctyArray := []string{"Revolutionary", "Very Progressive", "Progressive", "Neutral", "Traditional", "Very Traditional", "Reactionary"}

	// Get labels for each axis
	economicLabel := getLabel(econPercentage, econArray)
	diplomaticLabel := getLabel(diplPercentage, diplArray)
	governmentLabel := getLabel(govtPercentage, govtArray)
	societyLabel := getLabel(sctyPercentage, sctyArray)

	svg := fmt.Sprintf(`<svg width="%d" height="%d" xmlns="http://www.w3.org/2000/svg">
  <!-- Background -->
  <rect width="%d" height="%d" fill="#EEEEEE"/>
  
  <!-- Title -->
  <text x="20" y="90" text-anchor="start" font-family="Montserrat, sans-serif" font-size="80" font-weight="700" fill="#222222">8values</text>
  
  <!-- Icons (positioned as in 8values.js) -->
  <!-- Economic Icons -->
  <rect x="20" y="170" width="100" height="100" fill="#f44336" stroke="#222222" stroke-width="2" rx="10"/>
  <text x="70" y="235" text-anchor="middle" font-family="Arial, sans-serif" font-size="40" fill="white">⚖</text>
  <rect x="680" y="170" width="100" height="100" fill="#00897b" stroke="#222222" stroke-width="2" rx="10"/>
  <text x="730" y="235" text-anchor="middle" font-family="Arial, sans-serif" font-size="40" fill="white">💰</text>
  
  <!-- Diplomatic Icons -->
  <rect x="20" y="290" width="100" height="100" fill="#ff9800" stroke="#222222" stroke-width="2" rx="10"/>
  <text x="70" y="355" text-anchor="middle" font-family="Arial, sans-serif" font-size="40" fill="white">🏛</text>
  <rect x="680" y="290" width="100" height="100" fill="#03a9f4" stroke="#222222" stroke-width="2" rx="10"/>
  <text x="730" y="355" text-anchor="middle" font-family="Arial, sans-serif" font-size="40" fill="white">🌍</text>
  
  <!-- Government Icons -->
  <rect x="20" y="410" width="100" height="100" fill="#ffeb3b" stroke="#222222" stroke-width="2" rx="10"/>
  <text x="70" y="475" text-anchor="middle" font-family="Arial, sans-serif" font-size="40" fill="#222222">🗽</text>
  <rect x="680" y="410" width="100" height="100" fill="#3f51b5" stroke="#222222" stroke-width="2" rx="10"/>
  <text x="730" y="475" text-anchor="middle" font-family="Arial, sans-serif" font-size="40" fill="white">⚔</text>
  
  <!-- Society Icons -->
  <rect x="20" y="530" width="100" height="100" fill="#8bc34a" stroke="#222222" stroke-width="2" rx="10"/>
  <text x="70" y="595" text-anchor="middle" font-family="Arial, sans-serif" font-size="40" fill="white">🏛</text>
  <rect x="680" y="530" width="100" height="100" fill="#9c27b0" stroke="#222222" stroke-width="2" rx="10"/>
  <text x="730" y="595" text-anchor="middle" font-family="Arial, sans-serif" font-size="40" fill="white">⚡</text>
  
  <!-- Background bars (black) -->
  <rect x="120" y="180" width="560" height="80" fill="#222222"/>
  <rect x="120" y="300" width="560" height="80" fill="#222222"/>
  <rect x="120" y="420" width="560" height="80" fill="#222222"/>
  <rect x="120" y="540" width="560" height="80" fill="#222222"/>
  
  <!-- Left-side bars (equality, might, liberty, tradition) -->
  <rect x="120" y="184" width="%.1f" height="72" fill="#f44336"/>
  <rect x="120" y="304" width="%.1f" height="72" fill="#ff9800"/>
  <rect x="120" y="424" width="%.1f" height="72" fill="#ffeb3b"/>
  <rect x="120" y="544" width="%.1f" height="72" fill="#8bc34a"/>
  
  <!-- Right-side bars (wealth, peace, authority, progress) -->
  <rect x="%.1f" y="184" width="%.1f" height="72" fill="#00897b"/>
  <rect x="%.1f" y="304" width="%.1f" height="72" fill="#03a9f4"/>
  <rect x="%.1f" y="424" width="%.1f" height="72" fill="#3f51b5"/>
  <rect x="%.1f" y="544" width="%.1f" height="72" fill="#9c27b0"/>
  
  <!-- Left-side percentage text (only if > 30%%) -->`,
		width, height, width, height,
		// Left bars (equality, might, liberty, tradition)
		5.6*econPercentage-2,      // equality bar width
		5.6*mightPercentage-2,     // might bar width
		5.6*govtPercentage-2,      // liberty bar width
		5.6*traditionPercentage-2, // tradition bar width
		// Right bars positioning and widths (wealth, peace, authority, progress)
		682-5.6*wealthPercentage, 5.6*wealthPercentage-2, // wealth bar
		682-5.6*diplPercentage, 5.6*diplPercentage-2, // peace bar
		682-5.6*authorityPercentage, 5.6*authorityPercentage-2, // authority bar
		682-5.6*sctyPercentage, 5.6*sctyPercentage-2) // progress bar

	// Add percentage text (only if > 30% as in 8values.js)
	if econPercentage > 30 {
		svg += fmt.Sprintf(`
  <text x="130" y="237.5" text-anchor="start" font-family="Montserrat, sans-serif" font-size="50" fill="#222222">%.1f%%</text>`, econPercentage)
	}
	if mightPercentage > 30 {
		svg += fmt.Sprintf(`
  <text x="130" y="357.5" text-anchor="start" font-family="Montserrat, sans-serif" font-size="50" fill="#222222">%.1f%%</text>`, mightPercentage)
	}
	if govtPercentage > 30 {
		svg += fmt.Sprintf(`
  <text x="130" y="477.5" text-anchor="start" font-family="Montserrat, sans-serif" font-size="50" fill="#222222">%.1f%%</text>`, govtPercentage)
	}
	if traditionPercentage > 30 {
		svg += fmt.Sprintf(`
  <text x="130" y="597.5" text-anchor="start" font-family="Montserrat, sans-serif" font-size="50" fill="#222222">%.1f%%</text>`, traditionPercentage)
	}
	if wealthPercentage > 30 {
		svg += fmt.Sprintf(`
  <text x="670" y="237.5" text-anchor="end" font-family="Montserrat, sans-serif" font-size="50" fill="#222222">%.1f%%</text>`, wealthPercentage)
	}
	if diplPercentage > 30 {
		svg += fmt.Sprintf(`
  <text x="670" y="357.5" text-anchor="end" font-family="Montserrat, sans-serif" font-size="50" fill="#222222">%.1f%%</text>`, diplPercentage)
	}
	if authorityPercentage > 30 {
		svg += fmt.Sprintf(`
  <text x="670" y="477.5" text-anchor="end" font-family="Montserrat, sans-serif" font-size="50" fill="#222222">%.1f%%</text>`, authorityPercentage)
	}
	if sctyPercentage > 30 {
		svg += fmt.Sprintf(`
  <text x="670" y="597.5" text-anchor="end" font-family="Montserrat, sans-serif" font-size="50" fill="#222222">%.1f%%</text>`, sctyPercentage)
	}

	// Add axis labels
	svg += fmt.Sprintf(`
  
  <!-- Axis Labels -->
  <text x="400" y="175" text-anchor="middle" font-family="Montserrat, sans-serif" font-size="30" font-weight="300" fill="#222222">Economic Axis: %s</text>
  <text x="400" y="295" text-anchor="middle" font-family="Montserrat, sans-serif" font-size="30" font-weight="300" fill="#222222">Diplomatic Axis: %s</text>
  <text x="400" y="415" text-anchor="middle" font-family="Montserrat, sans-serif" font-size="30" font-weight="300" fill="#222222">Government Axis: %s</text>
  <text x="400" y="535" text-anchor="middle" font-family="Montserrat, sans-serif" font-size="30" font-weight="300" fill="#222222">Society Axis: %s</text>
  
  <!-- Attribution -->
  <text x="780" y="60" text-anchor="end" font-family="Montserrat, sans-serif" font-size="30" font-weight="300" fill="#222222">8values.github.io</text>
</svg>`, economicLabel, diplomaticLabel, governmentLabel, societyLabel)

	return svg
}
