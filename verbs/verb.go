package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/gocarina/gocsv"
)

type Verb struct {
	Infinitive     string `json:"infinitive"      csv:"infinitive"`
	PastSimple     string `json:"past_simple"     csv:"past_simple"`
	PastParticiple string `json:"past_participle" csv:"past_participle"`
}

func loadVerbsJSON(filename string) ([]Verb, error) {
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	var vs []Verb
	err = json.Unmarshal(data, &vs)
	if err != nil {
		return nil, err
	}
	return vs, nil
}

func loadVerbsCSV(filename string) ([]Verb, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	var vs []Verb
	err = gocsv.UnmarshalFile(file, &vs)
	if err != nil {
		return nil, err
	}
	return vs, nil
}

func LoadVerbs(filename string) ([]Verb, error) {
	ext := filepath.Ext(filename)
	switch ext {
	case ".json":
		return loadVerbsJSON(filename)
	case ".csv":
		return loadVerbsCSV(filename)
	default:
	}
	return nil, fmt.Errorf("file %q is not support", filename)
}
