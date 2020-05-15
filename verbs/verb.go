package main

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/gocarina/gocsv"
)

type Verb struct {
	Infinitive     string `json:"infinitive"      csv:"infinitive"`
	SimplePast     string `json:"simple_past"     csv:"simple_past"`
	PastParticiple string `json:"past_participle" csv:"past_participle"`
}

func (v Verb) Value(vf VerbForm) string {
	switch vf {
	case Infinitive:
		return v.Infinitive
	case SimplePast:
		return v.SimplePast
	case PastParticiple:
		return v.PastParticiple
	default:
		panic(fmt.Errorf("invalid %s", vf))
	}
}

func (v *Verb) SetValue(vf VerbForm, value string) {
	switch vf {
	case Infinitive:
		v.Infinitive = value
	case SimplePast:
		v.SimplePast = value
	case PastParticiple:
		v.PastParticiple = value
	default:
		panic(fmt.Errorf("invalid %s", vf))
	}
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
		//return loadVerbsCSV(filename)
		return loadVerbsStandatdCSV(filename, nil)
	default:
	}
	return nil, fmt.Errorf("file %q is not support", filename)
}

func loadVerbsStandatdCSV(filename string, keys []string) ([]Verb, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	r := csv.NewReader(file)
	if len(keys) == 0 {
		record, err := r.Read()
		if err != nil {
			return nil, err
		}
		keys = record
	}
	var vs []Verb
	var v Verb
	for line := 2; ; line++ {
		record, err := r.Read()
		if err != nil {
			if err == io.EOF {
				break
			}
			return nil, err
		}
		ws, err := parseRecordVerb(record)
		if err != nil {
			err = fmt.Errorf("record on line %d: %s", line, err)
			return nil, err
		}

		verbByWFs(ws, &v)
		//printJSON(ws)

		vs = append(vs, v)
	}
	return vs, nil
}

func verbByWFs(ws [][]*WordInfo, v *Verb) error {
	for i, vf := range verbForms {
		v.SetValue(vf, ws[i][0].Word)
	}
	return nil
}

func printJSON(v interface{}) {
	data, err := json.MarshalIndent(v, "", "\t")
	checkError(err)
	fmt.Println(string(data))
}
