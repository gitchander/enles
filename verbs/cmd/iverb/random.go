package main

import (
	"math/rand"

	"github.com/gitchander/enles/utils/random"
	"github.com/gitchander/enles/verbs/v1/verbs"
)

var correctPhrases = []string{
	"Right",
	"Good",
	"Very good",
	"Success",
}

// random correct phrase
func randCorrectPhrase(r *rand.Rand) string {
	return random.RandStringByCorpus(r, correctPhrases)
}

func randVerb(r *rand.Rand, corpus []verbs.Verb) verbs.Verb {
	return corpus[r.Intn(len(corpus))]
}

func randVerbForm(r *rand.Rand, corpus []verbs.VerbForm) verbs.VerbForm {
	return corpus[r.Intn(len(corpus))]
}
