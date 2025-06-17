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
		Threshold: 0.6,
		Slogan:    "No Gods No Masters",
	},
	{
		Name:      "pragmatism",
		Pair:      "",
		Color:     "#808080",
		Label:     "Pragmatist",
		Threshold: 0.3,
		Slogan:    "Practical Solutions",
	},
	{
		Name:      "feminism",
		Pair:      "",
		Color:     "#ff69b4",
		Label:     "Feminist",
		Threshold: 0.6,
		Slogan:    "Gender Equality",
	},
	{
		Name:      "complotism",
		Pair:      "",
		Color:     "#8b0000",
		Label:     "Conspiracist",
		Threshold: 0.7,
		Slogan:    "Question Everything",
	},
	{
		Name:      "veganism",
		Pair:      "",
		Color:     "#228b22",
		Label:     "Vegan",
		Threshold: 0.3,
		Slogan:    "Animal Rights",
	},
	{
		Name:      "monarchism",
		Pair:      "",
		Color:     "#ffd700",
		Label:     "Monarchist",
		Threshold: 0.3,
		Slogan:    "Royal Tradition",
	},
	{
		Name:      "religion",
		Pair:      "",
		Color:     "#4b0082",
		Label:     "Missionary",
		Threshold: 0.3,
		Slogan:    "Faithful Believer",
	},
}
