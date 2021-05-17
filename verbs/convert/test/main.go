package main

import (
	"encoding/json"
	"flag"
	"io/ioutil"
	"log"

	"github.com/gitchander/enles/verbs/convert"
)

func main() {
	var filename string
	flag.StringVar(&filename, "filename", "", "source file name")
	flag.Parse()

	err := run(filename)
	checkError(err)
}

func checkError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func run(filename string) error {
	vs, err := convert.LoadVerbsV2CSV(filename, nil)
	if err != nil {
		return err
	}
	data, err := json.MarshalIndent(vs, "", "\t")
	if err != nil {
		return err
	}
	return ioutil.WriteFile("irregular_verbs.json", data, 0666)
}
