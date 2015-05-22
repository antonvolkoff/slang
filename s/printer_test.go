package s

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var printTestcases = map[string]Item{
	// true, false and nil
	"nil":   Nil{},
	"true":  True{},
	"false": False{},

	// Numbers
	"1": Integer{Value: 1},
	"7": Integer{Value: 7},

	// Symbols
	"+":       Symbol{Value: "+"},
	"abc":     Symbol{Value: "abc"},
	"abc5":    Symbol{Value: "abc5"},
	"abc-def": Symbol{Value: "abc-def"},

	// Strings
	`"abc"`:               String{Value: "abc"},
	`"abc (with parens)"`: String{Value: "abc (with parens)"},
	`"abc"def"`:           String{Value: "abc\"def"},
	`""`:                  String{Value: ""},

	// Lists
	"(+ 1 2)": List{Value: []Item{
		Symbol{Value: "+"},
		Integer{Value: 1},
		Integer{Value: 2},
	}},
	"((3 4))": List{Value: []Item{
		List{Value: []Item{
			Integer{Value: 3},
			Integer{Value: 4},
		}},
	}},
	"(+ 1 (+ 2 3))": List{Value: []Item{
		Symbol{Value: "+"},
		Integer{Value: 1},
		List{Value: []Item{
			Symbol{Value: "+"},
			Integer{Value: 2},
			Integer{Value: 3},
		}},
	}},
	"(* 1 2)": List{Value: []Item{
		Symbol{Value: "*"},
		Integer{Value: 1},
		Integer{Value: 2},
	}},
	"(** 1 2)": List{Value: []Item{
		Symbol{Value: "**"},
		Integer{Value: 1},
		Integer{Value: 2},
	}},

	// Keywords
	":kw": Keyword{Value: "kw"},

	// Map
	`{"a" 1}`: Hash{Value: []KeyValue{
		KeyValue{Key: String{Value: "a"}, Value: Integer{Value: 1}},
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
