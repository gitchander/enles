package main

import (
	"fmt"
	"strings"
)

const (
	valuesSeparator = "/"
)

func parseRecordVerb(record []string) ([][]*WordInfo, error) {
	n := len(verbForms)
	if len(record) != n {
		return nil, fmt.Errorf("wrong number of fields for record: have %d, want %d", len(record), n)
	}
	ws := make([][]*WordInfo, n)
	for i := range verbForms {
		w, err := parseField(record[i])
		if err != nil {
			err = fmt.Errorf("field index %d: %s", i, err)
			return nil, err
		}
		ws[i] = w
	}
	return ws, nil
}

// parse variants of the verb form
func parseField(field string) ([]*WordInfo, error) {
	if field == "" {
		return nil, fmt.Errorf("field is empty")
	}
	var wis []*WordInfo
	vs := strings.Split(field, valuesSeparator)
	for _, v := range vs {
		wi, err := ParseWordInfo(v)
		if err != nil {
			err = fmt.Errorf("word tag %q: %s", v, err)
			return nil, err
		}
		wis = append(wis, wi)
	}
	return wis, nil
}

type WordInfo struct {
	Word          string `json:"word"`
	Note          string `json:"note,omitempty"`
	Transcription string `json:"transcription,omitempty"`
	Regular       bool   `json:"regular,omitempty"`
}

func ParseWordInfo(s string) (*WordInfo, error) {

	var wi WordInfo

	s = strings.TrimSpace(s)
	f := func(r rune) bool {
		return !IsSymbolOfWord(r)
	}
	index := strings.IndexFunc(s, f)
	if index == -1 {
		wi.Word = s
		if wi.Word == "" {
			return nil, fmt.Errorf("word is not exist")
		}
		return &wi, nil
	}
	wi.Word = s[:index]
	s = s[index:]

	if wi.Word == "" {
		return nil, fmt.Errorf("word is not exist")
	}

	for len(s) > 0 {
		s = strings.TrimSpace(s)
		if len(s) == 0 {
			break
		}

		if s[0] == '(' {

			index := strings.IndexByte(s, ')')
			if index == -1 {
				return nil, fmt.Errorf("there isn't close bracket %q", ')')
			}
			wi.Note = s[1:index]
			s = s[index+1:]

		} else if s[0] == '[' {

			index := strings.IndexByte(s, ']')
			if index == -1 {
				return nil, fmt.Errorf("there isn't close bracket %q", ']')
			}
			wi.Transcription = s[1:index]
			s = s[index+1:]

		} else {
			strRegular := "REGULAR"
			if strings.HasPrefix(s, strRegular) {
				wi.Regular = true
				s = s[len(strRegular):]
			} else {
				return nil, fmt.Errorf("wrong tag: %q", s)
			}
		}
	}

	return &wi, nil
}

// unicode.IsSpace()
func IsSymbolOfWord(r rune) bool {
	if ('a' <= r) && (r <= 'z') {
		return true
	}
	if ('A' <= r) && (r <= 'Z') {
		return true
	}
	if r == '-' {
		return true
	}
	return false
}
