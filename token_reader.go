package main

import (
	"strings"
	"unicode"
)

type TokenReader struct {
	Tokens []rune
	source string
}

func NewTokenReader(source string) *TokenReader {
	return &TokenReader{source: source}
}

func (t *TokenReader) Run() {
	r := strings.NewReader(t.source)

	var ch rune
	for {
		ch, _, _ = r.ReadRune()

		switch {
		case unicode.IsSpace(ch):
			// Skip
		default:
			t.Tokens = append(t.Tokens, ch)
		}

		if r.Len() == 0 {
			break
		}
	}
}
