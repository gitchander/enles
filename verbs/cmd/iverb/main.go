package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/gitchander/enles/utils/random"
	"github.com/gitchander/enles/verbs/v1/verbs"
)

func main() {
	var filename string
	flag.StringVar(&filename, "filename", "irregular_verbs.csv", "verbs source file name")
	flag.Parse()
	err := run(filename)
	checkError(err)
}

func run(filename string) error {
	vs, err := verbs.LoadVerbs(filename)
	if err != nil {
		return err
	}
	return runWithVerbs(vs)
}

func runWithVerbs(vs []verbs.Verb) error {

	const prefix = "\t"

	r := random.NewRandNow()

	lr := NewLineReader(os.Stdin)

	var tasker Tasker = newVerbTasker(r, vs)
	var task Task

	for {
		if task == nil {
			newTask, ok := tasker.Next()
			if !ok {
				break
			}
			task = newTask
		}

		fmt.Print(task.Question())

		line, err := lr.ReadLine()
		if err != nil {
			return err
		}

		if (line == "quit") || (line == "exit") {
			break
		}

		// if line == "stat" {
		// 	continue
		// }

		ok, failMessage := task.CheckAnswer(line)
		if !ok {
			fmt.Println()
			fmt.Printf("%s%s\n", prefix, failMessage)
			fmt.Println()
			continue
		}

		fmt.Println()
		fmt.Printf("%s%s\n", prefix, randCorrectPhrase(r))
		fmt.Println()

		task = nil // reset task
	}
	return nil
}
