package s

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var printTestcases = map[string]*Node{
	// true, false and nil
	"nil":   &Node{Type: "nil"},
	"true":  &Node{Type: "true"},
	"false": &Node{Type: "false"},

	// Numbers
	"1": &Node{Type: "number", Value: 1},
	"7": &Node{Type: "number", Value: 7},

	// Symbols
	"+":       &Node{Type: "symbol", Value: "+"},
	"abc":     &Node{Type: "symbol", Value: "abc"},
	"abc5":    &Node{Type: "symbol", Value: "abc5"},
	"abc-def": &Node{Type: "symbol", Value: "abc-def"},

	// Strings
	`"abc"`:               &Node{Type: "string", Value: "abc"},
	`"abc (with parens)"`: &Node{Type: "string", Value: "abc (with parens)"},
	`"abc"def"`:           &Node{Type: "string", Value: "abc\"def"},
	`""`:                  &Node{Type: "string", Value: ""},

	// Lists
	"(+ 1 2)": &Node{Type: "list", Children: []*Node{
		&Node{Type: "symbol", Value: "+"},
		&Node{Type: "number", Value: 1},
		&Node{Type: "number", Value: 2},
	}},
	"((3 4))": &Node{Type: "list", Children: []*Node{
		&Node{Type: "list", Children: []*Node{
			&Node{Type: "number", Value: 3},
			&Node{Type: "number", Value: 4},
		}},
	}},
	"(+ 1 (+ 2 3))": &Node{Type: "list", Children: []*Node{
		&Node{Type: "symbol", Value: "+"},
		&Node{Type: "number", Value: 1},
		&Node{Type: "list", Children: []*Node{
			&Node{Type: "symbol", Value: "+"},
			&Node{Type: "number", Value: 2},
			&Node{Type: "number", Value: 3},
		}},
	}},
	"(* 1 2)": &Node{Type: "list", Children: []*Node{
		&Node{Type: "symbol", Value: "*"},
		&Node{Type: "number", Value: 1},
		&Node{Type: "number", Value: 2},
	}},
	"(** 1 2)": &Node{Type: "list", Children: []*Node{
		&Node{Type: "symbol", Value: "**"},
		&Node{Type: "number", Value: 1},
		&Node{Type: "number", Value: 2},
	}},
}

func TestPrinter_ToString(t *testing.T) {
	for code, node := range printTestcases {
		p := NewPrinter(node)
		output, err := p.ToString()

		assert.NoError(t, err)
		assert.Equal(t, code, output)
	}
}
