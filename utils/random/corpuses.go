package random

import (
	"math/rand"
)

func RandIntByCorpus(r *rand.Rand, corpus []int) int {
	return corpus[r.Intn(len(corpus))]
}

func RandStringByCorpus(r *rand.Rand, corpus []string) string {
	return corpus[r.Intn(len(corpus))]
}
