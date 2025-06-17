package politiscales

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
	},
	{
		Name:      "essentialism",
		Pair:      "identity",
		Color:     "#34b634",
		Label:     "",
		Threshold: 0.0,
	},
	{
		Name:      "rehabilitative_justice",
		Pair:      "justice",
		Color:     "#14bee1",
		Label:     "justice",
		Threshold: 0.0,
	},
	{
		Name:      "punitive_justice",
		Pair:      "justice",
		Color:     "#e6cc27",
		Label:     "order",
		Threshold: 0.0,
	},
	{
		Name:      "progressive",
		Pair:      "culture",
		Color:     "#850083",
		Label:     "",
		Threshold: 0.0,
	},
	{
		Name:      "conservative",
		Pair:      "culture",
		Color:     "#970000",
		Label:     "family",
		Threshold: 0.0,
	},
	{
		Name:      "internationalism",
		Pair:      "globalism",
		Color:     "#3e6ffd",
		Label:     "humanity",
		Threshold: 0.0,
	},
	{
		Name:      "nationalism",
		Pair:      "globalism",
		Color:     "#ff8500",
		Label:     "fatherland",
		Threshold: 0.0,
	},
	{
		Name:      "communism",
		Pair:      "economy",
		Color:     "#cc0000",
		Label:     "socialism",
		Threshold: 0.0,
	},
	{
		Name:      "capitalism",
		Pair:      "economy",
		Color:     "#ffb800",
		Label:     "work",
		Threshold: 0.0,
	},
	{
		Name:      "regulation",
		Pair:      "markets",
		Color:     "#269B32",
		Label:     "",
		Threshold: 0.0,
	},
	{
		Name:      "laissez_faire",
		Pair:      "markets",
		Color:     "#6608C0",
		Label:     "liberty",
		Threshold: 0.0,
	},
	{
		Name:      "ecology",
		Pair:      "environment",
		Color:     "#a0e90d",
		Label:     "ecology",
		Threshold: 0.0,
	},
	{
		Name:      "production",
		Pair:      "environment",
		Color:     "#4deae9",
		Label:     "",
		Threshold: 0.0,
	},
	{
		Name:      "revolution",
		Pair:      "radicalism",
		Color:     "#eb1a66",
		Label:     "revolution",
		Threshold: 0.0,
	},
	{
		Name:      "reform",
		Pair:      "radicalism",
		Color:     "#0ee4c8",
		Label:     "",
		Threshold: 0.0,
	},

	// Unpaired axes
	{
		Name:      "anarchism",
		Pair:      "",
		Color:     "",
		Label:     "",
		Threshold: 0.9,
	},
	{
		Name:      "pragmatism",
		Pair:      "",
		Color:     "",
		Label:     "",
		Threshold: 0.5,
	},
	{
		Name:      "feminism",
		Pair:      "",
		Color:     "",
		Label:     "",
		Threshold: 0.9,
	},
	{
		Name:      "complotism",
		Pair:      "",
		Color:     "",
		Label:     "",
		Threshold: 0.9,
	},
	{
		Name:      "veganism",
		Pair:      "",
		Color:     "",
		Label:     "",
		Threshold: 0.5,
	},
	{
		Name:      "monarchism",
		Pair:      "",
		Color:     "",
		Label:     "",
		Threshold: 0.5,
	},
	{
		Name:      "religion",
		Pair:      "",
		Color:     "",
		Label:     "",
		Threshold: 0.5,
	},
}
