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

func TestRep_Variable(t *testing.T) {
	res1, err1 := Rep(`(def "x" 2)`)
	assert.NoError(t, err1)
	assert.Equal(t, "2", res1)

	res2, err2 := Rep("x")
	assert.NoError(t, err2)
	assert.Equal(t, "2", res2)

	res3, err3 := Rep("(+ 2 x)")
	assert.NoError(t, err3)
	assert.Equal(t, "4", res3)
}
