package convert

import (
	"encoding/csv"
	"fmt"
	"io"
	"os"

	vs1 "github.com/gitchander/enles/verbs/v1/verbs"
	vs2 "github.com/gitchander/enles/verbs/v2/verbs"
)

func LoadVerbsV2CSV(filename string, keys []string) ([]vs2.Verb, error) {
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
	var vs []vs2.Verb
	for line := 2; ; line++ {
		record, err := r.Read()
		if err != nil {
			if err == io.EOF {
				break
			}
			return nil, err
		}
		ws, err := vs1.ParseRecord(record)
		if err != nil {
			err = fmt.Errorf("record on line %d: %s", line, err)
			return nil, err
		}

		var v vs2.Verb

		verbV2ByWFs(ws, &v)

		vs = append(vs, v)
	}
	return vs, nil
}

func verbV2ByWFs(ws [][]*vs1.WordInfo, v *vs2.Verb) error {

	var words []vs2.Word

	var allVerbForms = []vs2.VerbForm{
		vs2.Infinitive,
		vs2.SimplePast,
		vs2.PastParticiple,
	}

	for i, vf := range allVerbForms {

		var variants []vs2.Variant

		for _, w := range ws[i] {

			if w.Regular {
				v.Regular = true
			}

			variant := vs2.Variant{
				Word:          w.Word,
				Note:          w.Note,
				Transcription: w.Transcription,
			}

			variants = append(variants, variant)
		}

		word := vs2.Word{
			VerbForm: vf.String(),
			Variants: variants,
		}

		words = append(words, word)
	}

	v.Forms = words

	return nil
}
