package main

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

var verbForms = []VerbForm{
	Infinitive,
	SimplePast,
	PastParticiple,
}
