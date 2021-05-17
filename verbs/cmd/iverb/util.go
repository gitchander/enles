package main

import (
	"bufio"
	"io"
	"log"
)

func checkError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

type LineReader struct {
	br *bufio.Reader
}

func NewLineReader(r io.Reader) *LineReader {
	return &LineReader{
		br: bufio.NewReader(r),
	}
}

func (p *LineReader) ReadLine() (string, error) {
	line, _, err := p.br.ReadLine()
	if err != nil {
		return "", err
	}
	return string(line), nil
}
