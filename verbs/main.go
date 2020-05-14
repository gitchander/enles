package main

import (
	"bufio"
	"flag"
	"fmt"
	"math/rand"
	"os"
	"strings"
)

func main() {

	var filename string
	flag.StringVar(&filename, "filename", "irregular_verbs.csv", "verbs source file name")
	flag.Parse()

	vs, err := LoadVerbs(filename)
	checkError(err)

	// for _, v := range vs {
	// 	fmt.Println(v)
	// }
	// return

	err = runTest(vs)
	checkError(err)
}

func runTest(vs []Verb) error {

	r := randNow()

	vfs := verbForms
	br := bufio.NewReader(os.Stdin)

	const prefix = "\t"

	var (
		v  = vs[r.Intn(len(vs))]
		vf = vfs[r.Intn(len(vfs))]
	)

	for {

		fmt.Printf("Enter %d forms of the verb %q (%s): ", len(vfs), v.Value(vf), vf)
		//fmt.Print("\n> ")

		lineBytes, _, err := br.ReadLine()
		if err != nil {
			return err
		}

		line := string(lineBytes)
		if (line == "quit") || (line == "exit") {
			break
		}

		if line == "stat" {
			continue
		}

		fields := strings.Fields(line)
		//fmt.Println("fields:", fields, line)
		if len(fields) != len(vfs) {
			fmt.Println()
			fmt.Printf("%sYou mast enter %d forms of the verb\n", prefix, len(vfs))
			fmt.Println()
			continue
		}

		var success bool = true
		for i, vf := range vfs {
			if fields[i] != v.Value(vf) {
				success = false
				break
			}
		}

		fmt.Println()
		if success {
			fmt.Printf("%s%s\n", prefix, rightString(r))
		} else {
			var vs []string
			for _, vf := range vfs {
				vs = append(vs, v.Value(vf))
			}
			fmt.Printf("%sWrong! There are: %s\n", prefix, strings.Join(vs, " "))
		}
		fmt.Println()

		v = vs[r.Intn(len(vs))]
		vf = vfs[r.Intn(len(vfs))]
	}
	return nil
}

var rights = []string{
	"Right",
	"Good",
	"Very good",
	"Success",
}

func rightString(r *rand.Rand) string {
	return rights[r.Intn(len(rights))]
}

func randVerbAndForm(r *rand.Rand, vs []Verb, vfs []VerbForm) (v Verb, vf VerbForm) {
	v = vs[r.Intn(len(vs))]
	vf = vfs[r.Intn(len(vfs))]
	return
}
