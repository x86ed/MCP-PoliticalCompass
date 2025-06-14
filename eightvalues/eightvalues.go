package eightvalues

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
