package main

import (
	"fmt"
	"strings"
)

func parseRecordVerb(record []string) ([][]*WordInfo, error) {
	n := len(verbForms)
	if len(record) != n {
		return nil, fmt.Errorf("wrong number of fields for record: have %d, want %d", len(record), n)
	}
	var ws [][]*WordInfo
	for i := range verbForms {
		w, err := parseField(record[i])
		if err != nil {
			return nil, err
		}
		ws = append(ws, w)
	}
	return ws, nil
}

func parseField(field string) ([]*WordInfo, error) {
	var wis []*WordInfo
	ws := strings.Split(field, "/")
	for _, w := range ws {
		wi, err := ParseWordInfo(w)
		if err != nil {
			return nil, err
		}
		wis = append(wis, wi)
	}
	return wis, nil
}

type WordInfo struct {
	Word          string `json:"word"`
	Note          string `json:"note,omitempty"`          // ( note )
	Transcription string `json:"transcription,omitempty"` // [ transcription ]
	Regular       bool   `json:"regular,omitempty"`
}

func ParseWordInfo(s string) (*WordInfo, error) {

	s = strings.TrimSpace(s)
	f := func(r rune) bool {
		return !IsSymbolOfWord(r)
	}
	index := strings.IndexFunc(s, f)
	if index == -1 {
		wi := &WordInfo{
			Word: s,
		}
		return wi, nil
	}
	word := s[:index]
	s = s[index:]

	var note string
	var transcription string
	var regular bool

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
			note = s[1:index]
			s = s[index+1:]

		} else if s[0] == '[' {

			index := strings.IndexByte(s, ']')
			if index == -1 {
				return nil, fmt.Errorf("there isn't close bracket %q", ']')
			}
			transcription = s[1:index]
			s = s[index+1:]

		} else {
			strRegular := "REGULAR"
			if strings.HasPrefix(s, strRegular) {
				regular = true
				s = s[len(strRegular):]
			} else {
				return nil, fmt.Errorf("wrong tag: %q", s)
			}
		}
	}

	wi := &WordInfo{
		Word:          word,
		Note:          note,
		Transcription: transcription,
		Regular:       regular,
	}

	return wi, nil
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
