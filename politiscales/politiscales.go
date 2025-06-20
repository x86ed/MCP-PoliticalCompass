package politiscales

import "fmt"

// Response values for question answers
// These correspond to the button onclick values in the UI
const (
	StronglyAgree    = 1.0        // Strongly Agree response value
	Agree            = 2.0 / 3.0  // Agree response value
	Neutral          = 0.0        // Neutral/Unsure response value
	Disagree         = -2.0 / 3.0 // Disagree response value
	StronglyDisagree = -1.0       // Strongly Disagree response value
)

// Question represents a political compass question with economic and social scoring data
type Question struct {
	Index      int32 // Question index/ID
	YesWeights []Weight
	NoWeights  []Weight
	Text       string // Question text
}

// Weight represents the effect of a response on a specific political axis
type Weight struct {
	Axis  string
	Value float64
}

type Axis struct {
	Name      string
	Pair      string
	Color     string
	Label     string
	Threshold float64
	Slogan    string
}

// Axes contains all the political axes with their properties
var Axes = []Axis{
	// Paired axes
	{
		Name:      "constructivism",
		Pair:      "identity",
		Color:     "#a425b6",
		Label:     "equality",
		Threshold: 0.0,
		Slogan:    "Social Constructor",
	},
	{
		Name:      "essentialism",
		Pair:      "identity",
		Color:     "#34b634",
		Label:     "",
		Threshold: 0.0,
		Slogan:    "Natural Order",
	},
	{
		Name:      "rehabilitative_justice",
		Pair:      "justice",
		Color:     "#14bee1",
		Label:     "justice",
		Threshold: 0.0,
		Slogan:    "Restorative Justice",
	},
	{
		Name:      "punitive_justice",
		Pair:      "justice",
		Color:     "#e6cc27",
		Label:     "order",
		Threshold: 0.0,
		Slogan:    "Law and Order",
	},
	{
		Name:      "progressive",
		Pair:      "culture",
		Color:     "#850083",
		Label:     "",
		Threshold: 0.0,
		Slogan:    "Forward Thinking",
	},
	{
		Name:      "conservative",
		Pair:      "culture",
		Color:     "#970000",
		Label:     "family",
		Threshold: 0.0,
		Slogan:    "Traditional Values",
	},
	{
		Name:      "internationalism",
		Pair:      "globalism",
		Color:     "#3e6ffd",
		Label:     "humanity",
		Threshold: 0.0,
		Slogan:    "Global Citizen",
	},
	{
		Name:      "nationalism",
		Pair:      "globalism",
		Color:     "#ff8500",
		Label:     "fatherland",
		Threshold: 0.0,
		Slogan:    "Nation First",
	},
	{
		Name:      "communism",
		Pair:      "economy",
		Color:     "#cc0000",
		Label:     "socialism",
		Threshold: 0.0,
		Slogan:    "Workers United",
	},
	{
		Name:      "capitalism",
		Pair:      "economy",
		Color:     "#ffb800",
		Label:     "work",
		Threshold: 0.0,
		Slogan:    "Free Markets",
	},
	{
		Name:      "regulation",
		Pair:      "markets",
		Color:     "#269B32",
		Label:     "",
		Threshold: 0.0,
		Slogan:    "Guided Economy",
	},
	{
		Name:      "laissez_faire",
		Pair:      "markets",
		Color:     "#6608C0",
		Label:     "liberty",
		Threshold: 0.0,
		Slogan:    "Market Freedom",
	},
	{
		Name:      "ecology",
		Pair:      "environment",
		Color:     "#a0e90d",
		Label:     "ecology",
		Threshold: 0.0,
		Slogan:    "Green Future",
	},
	{
		Name:      "production",
		Pair:      "environment",
		Color:     "#4deae9",
		Label:     "",
		Threshold: 0.0,
		Slogan:    "Progress First",
	},
	{
		Name:      "revolution",
		Pair:      "radicalism",
		Color:     "#eb1a66",
		Label:     "revolution",
		Threshold: 0.0,
		Slogan:    "Radical Change",
	},
	{
		Name:      "reform",
		Pair:      "radicalism",
		Color:     "#0ee4c8",
		Label:     "",
		Threshold: 0.0,
		Slogan:    "Gradual Progress",
	},

	// Unpaired axes
	{
		Name:      "anarchism",
		Pair:      "",
		Color:     "#000000",
		Label:     "Anarchist",
		Threshold: 0.9,
		Slogan:    "No Gods No Masters",
	},
	{
		Name:      "pragmatism",
		Pair:      "",
		Color:     "#808080",
		Label:     "Pragmatist",
		Threshold: 0.5,
		Slogan:    "Practical Solutions",
	},
	{
		Name:      "feminism",
		Pair:      "",
		Color:     "#ff69b4",
		Label:     "Feminist",
		Threshold: 0.9,
		Slogan:    "Gender Equality",
	},
	{
		Name:      "complotism",
		Pair:      "",
		Color:     "#8b0000",
		Label:     "Conspiracist",
		Threshold: 0.9,
		Slogan:    "Question Everything",
	},
	{
		Name:      "veganism",
		Pair:      "",
		Color:     "#228b22",
		Label:     "Vegan",
		Threshold: 0.5,
		Slogan:    "Animal Rights",
	},
	{
		Name:      "monarchism",
		Pair:      "",
		Color:     "#ffd700",
		Label:     "Monarchist",
		Threshold: 0.5,
		Slogan:    "Royal Tradition",
	},
	{
		Name:      "religion",
		Pair:      "",
		Color:     "#4b0082",
		Label:     "Missionary",
		Threshold: 0.5,
		Slogan:    "Faithful Believer",
	},
}

// Generate SVG results display matching the original PolitiScales format
func GeneratePolitiscalesResultsSVG(results map[string]float64) string {
	// Define the axis pairs in display order using data from politiscales module
	// But maintain specific display order and labels for consistency
	axisPairs := []struct {
		leftAxis, rightAxis   string
		leftLabel, rightLabel string
		leftColor, rightColor string
	}{
		{"constructivism", "essentialism", "Constructivism", "Essentialism", "#a425b6", "#34b634"},
		{"rehabilitative_justice", "punitive_justice", "Rehabilitative Justice", "Punitive Justice", "#14bee1", "#e6cc27"},
		{"progressive", "conservative", "Progressive", "Conservative", "#850083", "#970000"},
		{"internationalism", "nationalism", "Internationalism", "Nationalism", "#3e6ffd", "#ff8500"},
		{"communism", "capitalism", "Communism", "Capitalism", "#cc0000", "#ffb800"},
		{"regulation", "laissez_faire", "Regulation", "Laissez-faire", "#269B32", "#6608C0"},
		{"ecology", "production", "Ecology", "Production", "#a0e90d", "#4deae9"},
		{"revolution", "reform", "Revolution", "Reform", "#eb1a66", "#0ee4c8"},
	}

	// Count qualifying badges to calculate required height
	var qualifyingBadges []struct {
		name  string
		label string
		score float64
		color string
	}

	for _, axis := range Axes {
		if axis.Pair == "" { // Unpaired axes only
			score := results[axis.Name]
			threshold := axis.Threshold * 100
			if score >= threshold && score > 0 {
				label := axis.Label
				if label == "" {
					label = axis.Name // Fallback to axis name if no label
				}
				color := axis.Color
				if color == "" {
					color = "#666666" // Default color if none specified
				}
				qualifyingBadges = append(qualifyingBadges, struct {
					name  string
					label string
					score float64
					color string
				}{axis.Name, label, score, color})
			}
		}
	}

	// Sort badges by score descending
	for i := 0; i < len(qualifyingBadges)-1; i++ {
		for j := i + 1; j < len(qualifyingBadges); j++ {
			if qualifyingBadges[j].score > qualifyingBadges[i].score {
				qualifyingBadges[i], qualifyingBadges[j] = qualifyingBadges[j], qualifyingBadges[i]
			}
		}
	}

	// Calculate required height: base + axes + spacing + badges section
	baseHeight := 600
	badgesHeight := 0
	if len(qualifyingBadges) > 0 {
		badgesHeight = 90 + (len(qualifyingBadges) * 25) // Header + badges
	}
	totalHeight := baseHeight + badgesHeight

	svg := fmt.Sprintf(`<svg width="800" height="%d" xmlns="http://www.w3.org/2000/svg">
  <defs>
    <style>
      .axis-label { font-family: Arial, sans-serif; font-size: 14px; font-weight: bold; }
      .percentage-text { font-family: Arial, sans-serif; font-size: 12px; fill: white; text-anchor: middle; }
      .title { font-family: Arial, sans-serif; font-size: 24px; font-weight: bold; text-anchor: middle; }
    </style>
  </defs>
  
  <!-- Background -->
  <rect width="800" height="%d" fill="#f8f9fa"/>
  
  <!-- Title -->
  <text x="400" y="40" class="title" fill="#333">PolitiScales Results</text>`, totalHeight, totalHeight)

	y := 80
	for _, pair := range axisPairs {
		leftScore := results[pair.leftAxis]
		rightScore := results[pair.rightAxis]

		// Calculate neutral space
		total := leftScore + rightScore
		neutral := 100 - total

		// Calculate widths for 600px bar
		barWidth := 600
		leftWidth := int((leftScore / 100) * float64(barWidth))
		rightWidth := int((rightScore / 100) * float64(barWidth))
		neutralWidth := barWidth - leftWidth - rightWidth

		if neutralWidth < 0 {
			neutralWidth = 0
		}

		// Axis labels
		svg += fmt.Sprintf(`
  <!-- %s vs %s -->
  <text x="100" y="%d" class="axis-label" fill="#333" text-anchor="end">%s</text>
  <text x="700" y="%d" class="axis-label" fill="#333">%s</text>`,
			pair.leftLabel, pair.rightLabel, y+15, pair.leftLabel, y+15, pair.rightLabel)

		// Progress bar
		barY := y + 20
		currentX := 100

		// Left side
		if leftWidth > 0 {
			svg += fmt.Sprintf(`
  <rect x="%d" y="%d" width="%d" height="30" fill="%s"/>
  <text x="%d" y="%d" class="percentage-text">%.0f%%</text>`,
				currentX, barY, leftWidth, pair.leftColor,
				currentX+leftWidth/2, barY+20, leftScore)
		}
		currentX += leftWidth

		// Neutral section
		if neutralWidth > 0 {
			svg += fmt.Sprintf(`
  <rect x="%d" y="%d" width="%d" height="30" fill="#e0e0e0"/>
  <text x="%d" y="%d" class="percentage-text" fill="#666">%.0f%%</text>`,
				currentX, barY, neutralWidth,
				currentX+neutralWidth/2, barY+20, neutral)
		}
		currentX += neutralWidth

		// Right side
		if rightWidth > 0 {
			svg += fmt.Sprintf(`
  <rect x="%d" y="%d" width="%d" height="30" fill="%s"/>
  <text x="%d" y="%d" class="percentage-text">%.0f%%</text>`,
				currentX, barY, rightWidth, pair.rightColor,
				currentX+rightWidth/2, barY+20, rightScore)
		}

		y += 65
	}

	// Add slogan section
	y += 20

	// Generate slogan based on top characteristics using data from politiscales module
	type characteristic struct {
		name  string
		value float64
	}

	var characteristics []characteristic
	for name, value := range results {
		if value > 0 {
			characteristics = append(characteristics, characteristic{name, value})
		}
	}

	// Sort by value descending
	for i := 0; i < len(characteristics)-1; i++ {
		for j := i + 1; j < len(characteristics); j++ {
			if characteristics[j].value > characteristics[i].value {
				characteristics[i], characteristics[j] = characteristics[j], characteristics[i]
			}
		}
	}

	// Create slogans map from module data
	slogans := make(map[string]string)
	for _, axis := range Axes {
		if axis.Slogan != "" {
			slogans[axis.Name] = axis.Slogan
		}
	}

	sloganParts := []string{}
	for i, char := range characteristics {
		if i >= 3 {
			break
		}
		if sloganText, exists := slogans[char.name]; exists && char.value >= 50 {
			sloganParts = append(sloganParts, sloganText)
		}
	}

	// If no high-scoring characteristics, try with lower threshold
	if len(sloganParts) == 0 {
		for i, char := range characteristics {
			if i >= 3 {
				break
			}
			if sloganText, exists := slogans[char.name]; exists && char.value >= 30 {
				sloganParts = append(sloganParts, sloganText)
			}
		}
	}

	var slogan string
	if len(sloganParts) > 0 {
		for i, part := range sloganParts {
			if i > 0 {
				slogan += " · "
			}
			slogan += part
		}
	} else {
		slogan = "Political Moderate"
	}

	svg += fmt.Sprintf(`
  <text x="400" y="%d" class="title" fill="#333" font-size="18">Political Identity</text>
  <text x="400" y="%d" class="axis-label" fill="#666" font-size="14" text-anchor="middle">%s</text>`,
		y, y+25, slogan)

	// Add bonus characteristics section
	y += 50
	svg += fmt.Sprintf(`
  <text x="400" y="%d" class="title" fill="#333" font-size="18">Additional Characteristics</text>`, y)

	// Use the badges we already calculated
	bonusY := y + 40

	displayedBonus := 0
	for _, badge := range qualifyingBadges {
		svg += fmt.Sprintf(`
  <circle cx="150" cy="%d" r="8" fill="%s"/>
  <text x="170" y="%d" class="axis-label" fill="#333">%s (%.1f%%)</text>`,
			bonusY+displayedBonus*25, badge.color,
			bonusY+displayedBonus*25+5, badge.label, badge.score)
		displayedBonus++
	}

	svg += `
</svg>`

	return svg
}
