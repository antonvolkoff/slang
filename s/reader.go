package s

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
	"unicode"
)

type Reader struct {
	position int
	tokens   []string
}

func NewReader() *Reader {
	return &Reader{position: -1}
}

func (r *Reader) Parse(code string) (Item, error) {
	r.tokens = r.Tokenize(code)
	if len(r.tokens) == 0 {
		return nil, fmt.Errorf("unexpected EOF while reading")
	}

	return r.ReadFromTokens()
}

func (r *Reader) ReadFromTokens() (Item, error) {
	var i Item
	token := r.peek()

	switch token {
	case "(":
		i := List{}
		for {
			cn, err := r.ReadFromTokens()
			if err != nil {
				return nil, err
			}
			i.Add(cn)

			if r.next() == ")" {
				r.peek() // Move to next one
				break
			}
		}
		return i, nil

	case ")":
		return nil, fmt.Errorf("unexpected ) at %d", r.position)

	case "{":
		i := Hash{}
		for {
			key, err := r.ReadFromTokens()
			if err != nil {
				return nil, err
			}
			value, err := r.ReadFromTokens()
			if err != nil {
				return nil, err
			}
			kv := KeyValue{Key: key, Value: value}
			i.Add(kv)

			if r.next() == "}" {
				r.peek() // Move to next one
				break
			}
		}
		return i, nil

	case "}":
		return nil, fmt.Errorf("unexpected } at %d", r.position)

	default:
		return r.readAtom(token)
	}

	return i, nil
}

func (r *Reader) Tokenize(code string) []string {
	results := make([]string, 0, 1)
	// Work around lack of quoting in backtick
	re := regexp.MustCompile(`[\s,]*(~@|[\[\]{}()'` + "`" +
		`~^@]|"(?:\\.|[^\\"])*"|;.*|[^\s\[\]{}('"` + "`" +
		`,;)]*)`)
	for _, group := range re.FindAllStringSubmatch(code, -1) {
		if (group[1] == "") || (group[1][0] == ';') {
			continue
		}
		results = append(results, group[1])
	}
	return results
}

func (r *Reader) peek() string {
	r.position += 1
	token := r.tokens[r.position]
	return token
}

func (r *Reader) next() string {
	return r.tokens[r.position+1]
}

func (r *Reader) readAtom(token string) (Item, error) {
	switch {
	case unicode.IsNumber(rune(token[0])):
		i := Integer{}
		val, err := strconv.Atoi(token)
		if err != nil {
			return nil, err
		}
		i.Value = int64(val)

		return i, nil

	case string(token[0]) == `"` && string(token[len(token)-1]) == `"`:
		i := String{}
		var val string
		val = token[1 : len(token)-1]
		val = strings.Replace(val, `\"`, `"`, -1)
		val = strings.Replace(val, `\n`, "\n", -1)
		i.Value = val
		return i, nil

	case string(token[0]) == ":":
		i := Keyword{Value: token[1:]}
		return i, nil

	case token == "nil":
		return Nil{}, nil
	case token == "true":
		return True{}, nil
	case token == "false":
		return False{}, nil
	default:
		return Symbol{Value: token}, nil
	}
}
