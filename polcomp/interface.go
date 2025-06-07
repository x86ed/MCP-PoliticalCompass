package polcomp

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
