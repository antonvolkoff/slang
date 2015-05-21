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

func TestRep_Def(t *testing.T) {
	res1, err1 := Rep(`(def x 2)`)
	assert.NoError(t, err1)
	assert.Equal(t, "2", res1)

	res2, err2 := Rep("x")
	assert.NoError(t, err2)
	assert.Equal(t, "2", res2)

	res3, err3 := Rep("(+ 2 x)")
	assert.NoError(t, err3)
	assert.Equal(t, "4", res3)

	res4, err4 := Rep("(def y (+ 1 7))")
	assert.NoError(t, err4)
	assert.Equal(t, "8", res4)
}

func TestRep_Let(t *testing.T) {
	res1, err1 := Rep(`(let {z 9} z)`)
	assert.NoError(t, err1)
	assert.Equal(t, "9", res1)

	res2, err2 := Rep(`(let {z (+ 2 3)} (+ 1 z))`)
	assert.NoError(t, err2)
	assert.Equal(t, "6", res2)

	res3, err3 := Rep(`(let {p (+ 2 3) q (+ 2 p)} (+ p q))`)
	assert.NoError(t, err3)
	assert.Equal(t, "12", res3)
}

func TestRep_Outer(t *testing.T) {
	res1, err1 := Rep(`(def a 4)`)
	assert.NoError(t, err1)
	assert.Equal(t, "4", res1)

	res2, err2 := Rep(`(let {q 9} q)`)
	assert.NoError(t, err2)
	assert.Equal(t, "9", res2)

	res3, err3 := Rep(`(let {q 9} a)`)
	assert.NoError(t, err3)
	assert.Equal(t, "4", res3)

	res4, err4 := Rep(`(let {z 2} (let {q 9} a))`)
	assert.NoError(t, err4)
	assert.Equal(t, "4", res4)

	res5, err5 := Rep(`(let {z a} z)`)
	assert.NoError(t, err5)
	assert.Equal(t, "4", res5)
}

var listCases = map[string]string{
	"(list)":               "()",
	"(list? (list))":       "true",
	"(empty? (list))":      "true",
	"(empty? (list 1))":    "false",
	"(list 1 2 3)":         "(1 2 3)",
	"(count (list 1 2 3))": "3",
	"(count (list))":       "0",
	"(count nil)":          "0",
}

func TestRep_Lists(t *testing.T) {
	for input, output := range listCases {
		res, err := Rep(input)
		assert.NoError(t, err)
		assert.Equal(t, output, res)
	}
}

var ifCases = map[string]string{
	"(if true 7 8)":              "7",
	"(if false 7 8)":             "8",
	"(if true (+ 1 7) (+ 1 8))":  "8",
	"(if false (+ 1 7) (+ 1 8))": "9",
	"(if nil 7 8)":               "8",
	"(if 0 7 8)":                 "7",
	`(if "" 7 8)`:                "7",
	"(if (list) 7 8)":            "7",
	"(if (list 1 2 3) 7 8)":      "7",
	"(if false (+ 1 7))":         "nil",
	"(if true (+ 1 7))":          "8",
}

func TestRep_If(t *testing.T) {
	for input, output := range ifCases {
		res, err := Rep(input)
		assert.NoError(t, err)
		assert.Equal(t, output, res)
	}
}
