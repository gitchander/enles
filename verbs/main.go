package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"math/rand"
	"time"
)

func main() {
	var filename string
	flag.StringVar(&filename, "filename", "", "source file name")
	flag.Parse()

	data, err := ioutil.ReadFile(filename)
	checkError(err)

	var verbs []Verb
	err = json.Unmarshal(data, &verbs)
	checkError(err)

	// for _, v := range verbs {
	// 	fmt.Println(v.Infinitive)
	// }

	runTest(verbs)
}

func checkError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

type Verb struct {
	Infinitive     string `json:"infinitive"`
	PastSimple     string `json:"past-simple"`
	PastParticiple string `json:"past-participle"`
}

type Translation struct {
	Language string   `json:"language"`
	Variants []string `json:"variants"`
	// Samples  []string `json:"samples"`
}

func runTest(verbs []Verb) {

	r := randNow()
	ds := makeSerialInts(3)

	for {

		v := verbs[r.Intn(len(verbs))]

		if v.Infinitive == "" {
			continue
		}

		Scramble(r, IntSlice(ds))

		t0 := indexToString(ds[0])
		v0 := indexToValue(ds[0], v)

		t1 := indexToString(ds[1])
		v1 := indexToValue(ds[1], v)

		fmt.Printf("%s:%q, %s - ?\n", t0, v0, t1)
		fmt.Print("> ")

		var line string
		fmt.Scanln(&line)

		if line != v1 {
			fmt.Printf("  wrong! it is %q\n", v1)
		} else {
			fmt.Println("  right")
		}
		fmt.Println()
	}
}

func randNow() *rand.Rand {
	return rand.New(rand.NewSource(time.Now().UnixNano()))
}

func makeSerialInts(n int) []int {
	ds := make([]int, n)
	for i := range ds {
		ds[i] = i
	}
	return ds
}

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

func indexToString(i int) string {
	switch i {
	case 0:
		return "infinitive"
	case 1:
		return "past-simple"
	case 2:
		return "past-participle"
	default:
		return "unknown"
	}
}

func indexToValue(i int, v Verb) string {
	switch i {
	case 0:
		return v.Infinitive
	case 1:
		return v.PastSimple
	case 2:
		return v.PastParticiple
	default:
		return "unknown"
	}
}
