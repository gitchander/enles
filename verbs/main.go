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

	for _, v := range verbs {
		fmt.Println(v.Infinitive)
	}

}

func checkError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

type Verb struct {
	Infinitive     string        `json:"infinitive"`
	PastSimple     string        `json:"past-simple"`
	PastParticiple string        `json:"past-participle"`
	Translations   []Translation `json:"translations"`
}

type Translation struct {
	Language string   `json:"language"`
	Variants []string `json:"variants"`

	// Samples  []string `json:"samples"`
}

func runTest(verbs []Verb) {

	r := randNow()

	for {
		v := verbs[r.Intn(len(verbs))]

		switch k := r.Intn(3); k {
		case 0:
			fmt.Printf("infinitive %q, ps ?\n", v.Infinitive)

			var line string
			fmt.Scanln(&line)

			if line != v.PastSimple {
				fmt.Println("Wrong!")
			} else {
				fmt.Println("Right!")
			}

		case 1:

		}
	}
}

func randNow() *rand.Rand {
	return rand.New(rand.NewSource(time.Now().UnixNano()))
}

func randMixInts(r *rand.Rand, as []int) {
	// to do
}
