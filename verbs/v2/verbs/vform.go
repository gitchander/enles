package verbs

import (
	"fmt"
)

type VerbForm int

const (
	Infinitive VerbForm = iota
	SimplePast
	PastParticiple
)

func (vf VerbForm) String() string {
	switch vf {
	case Infinitive:
		return "infinitive"
	case SimplePast:
		return "simple-past"
	case PastParticiple:
		return "past-participle"
	default:
		return fmt.Sprintf("VerbForm(%d)", vf)
	}
}

var valueToStringVerbForm = map[VerbForm]string{
	Infinitive:     "infinitive",
	SimplePast:     "simple-past",
	PastParticiple: "past-participle",
}

var stringToValueVerbForm = map[string]VerbForm{
	"infinitive":      Infinitive,
	"simple-past":     SimplePast,
	"past-participle": PastParticiple,
}

func ParseVerbForm(s string) (VerbForm, error) {
	v, ok := stringToValueVerbForm[s]
	if !ok {
		return 0, fmt.Errorf("Invalid VerbForm %q", s)
	}
	return v, nil
}
