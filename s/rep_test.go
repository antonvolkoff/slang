package s

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var repTestcases = map[string]string{
	"(+ 1 2)":                            "3",
	"(+ 5 (* 2 3))":                      "11",
	"(- (+ 5 (* 2 3)) 3)":                "8",
	"(/ (- (+ 5 (* 2 3)) 3) 4)":          "2",
	"(/ (- (+ 515 (* 222 311)) 302) 27)": "2565",
}

func TestRep(t *testing.T) {
	for exp, out := range repTestcases {
		result, err := Rep(exp)

		assert.NoError(t, err)
		assert.Equal(t, out, result)
	}
}
