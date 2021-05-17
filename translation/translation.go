package translation

type Translation struct {
	Language string   `json:"language"`
	Variants []string `json:"variants"`
	// Samples  []string `json:"samples"`
}
