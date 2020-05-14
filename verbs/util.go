package main

import (
	"log"
	"math/rand"
	"time"
)

func checkError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func randNow() *rand.Rand {
	return rand.New(rand.NewSource(time.Now().UnixNano()))
}

func serialInts(n int) []int {
	ds := make([]int, n)
	for i := range ds {
		ds[i] = i
	}
	return ds
}
