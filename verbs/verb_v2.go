package main

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func loadVerbsV2CSV(filename string, keys []string) ([]VerbV2, error) {
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
	var vs []VerbV2
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

		var v VerbV2

		verbV2ByWFs(ws, &v)
		//printJSON(v)

		vs = append(vs, v)
	}
	return vs, nil
}

type VerbV2 struct {
	Regular bool   `json:"regular,omitempty"`
	Note    string `json:"note,omitempty"`
	Forms   []Word `json:"forms"`
}

type Word struct {
	VerbForm string    `json:"verb_form"`
	Variants []Variant `json:"variants"`
}

type Variant struct {
	Word          string `json:"word"`
	Note          string `json:"note,omitempty"`
	Transcription string `json:"transcription,omitempty"`
}

func verbV2ByWFs(ws [][]*WordInfo, v *VerbV2) error {

	var words []Word

	for i, vf := range verbForms {

		var variants []Variant

		for _, w := range ws[i] {

			if w.Regular {
				v.Regular = true
			}

			variant := Variant{
				Word:          w.Word,
				Note:          w.Note,
				Transcription: w.Transcription,
			}

			variants = append(variants, variant)
		}

		word := Word{
			VerbForm: vf.String(),
			Variants: variants,
		}

		words = append(words, word)
	}

	v.Forms = words

	return nil
}

func test2(filename string) error {
	vs, err := loadVerbsV2CSV(filename, nil)
	if err != nil {
		return err
	}
	data, err := json.MarshalIndent(vs, "", "\t")
	if err != nil {
		return err
	}
	return ioutil.WriteFile("irregular_verbs.json", data, 0666)
}
