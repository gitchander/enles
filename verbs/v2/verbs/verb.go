package verbs

type Verb struct {
	Regular bool   `json:"regular,omitempty"`
	Note    string `json:"note,omitempty"`
	Forms   []Word `json:"forms"`
}

type Word struct {
	VerbForm string    `json:"verb_form"`
	Variants []Variant `json:"variants"`
}

type Variant struct {
	Word          string `json:"word"`
	Note          string `json:"note,omitempty"`
	Transcription string `json:"transcription,omitempty"`
}
