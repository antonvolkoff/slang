package s

import (
	"testing"

	// "github.com/k0kubun/pp"
	"github.com/stretchr/testify/assert"
)

var testcases = map[string]Item{
	// true, false and nil
	"nil":   Nil{},
	"true":  True{},
	"false": False{},

	// Numbers
	"1":     Integer{Value: 1},
	"7":     Integer{Value: 7},
	"  7  ": Integer{Value: 7},

	// Symbols
	"+":         Symbol{Value: "+"},
	"abc":       Symbol{Value: "abc"},
	"   abc   ": Symbol{Value: "abc"},
	"abc5":      Symbol{Value: "abc5"},
	"abc-def":   Symbol{Value: "abc-def"},

	// Strings
	`"abc"`:               String{Value: "abc"},
	`   "abc"   `:         String{Value: "abc"},
	`"abc (with parens)"`: String{Value: "abc (with parens)"},
	`"abc\"def"`:          String{Value: "abc\"def"},
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
	"  ( +   1   (+   2 3   )   )  ": List{Value: []Item{
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

	// Commas as whitespace
	"(1 2, 3,,,,),,": List{Value: []Item{
		Integer{Value: 1},
		Integer{Value: 2},
		Integer{Value: 3},
	}},

	// Keywords
	":kw": Keyword{Value: "kw"},
	"(:kw1 :kw2 :kw3)": List{Value: []Item{
		Keyword{Value: "kw1"},
		Keyword{Value: "kw2"},
		Keyword{Value: "kw3"},
	}},

	// Hash
	`{"abc" 1}`: Hash{Value: []KeyValue{
		KeyValue{Key: String{Value: "abc"}, Value: Integer{Value: 1}},
	}},
	`{"a" {"b" 2}}`: Hash{Value: []KeyValue{
		KeyValue{Key: String{Value: "a"}, Value: Hash{Value: []KeyValue{
			KeyValue{Key: String{Value: "b"}, Value: Integer{Value: 2}},
		}}},
	}},
}

func TestReader_Parse(t *testing.T) {
	for code, node := range testcases {
		r := NewReader()
		n, err := r.Parse(code)

		assert.NoError(t, err)
		if assert.NotNil(t, n) {
			assert.Equal(t, node, n)
		}
	}
}
