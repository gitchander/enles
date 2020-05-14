package main

import (
	"math/rand"
)

type Swapper interface {
	Len() int
	Swap(i, j int)
}

type IntSlice []int

var _ Swapper = IntSlice(nil)

func (p IntSlice) Len() int      { return len(p) }
func (p IntSlice) Swap(i, j int) { p[i], p[j] = p[j], p[i] }

func Scramble(r *rand.Rand, p Swapper) {
	for n := p.Len(); n > 1; n-- {
		i := r.Intn(n)
		p.Swap(i, n-1)
	}
}
