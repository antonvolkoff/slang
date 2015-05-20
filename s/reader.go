package s

import (
	"fmt"
	"strconv"
	"strings"
	"unicode"
)

type Node struct {
	Type     string
	Value    interface{}
	Children []*Node
}

type Reader struct {
	position int
	tokens   []string
}

func NewReader() *Reader {
	return &Reader{position: -1}
}

func (r *Reader) Parse(code string) (*Node, error) {
	r.tokens = r.Tokenize(code)
	if len(r.tokens) == 0 {
		return nil, fmt.Errorf("unexpected EOF while reading")
	}

	return r.ReadFromTokens()
}

func (r *Reader) ReadFromTokens() (*Node, error) {
	n := &Node{}
	token := r.peek()

	switch token {
	case "(":
		n.Type = "list"
		for {
			cn, err := r.ReadFromTokens()
			if err != nil {
				return nil, err
			}
			n.Children = append(n.Children, cn)

			if r.next() == ")" {
				r.peek() // Move to next one
				break
			}
		}

	case ")":
		return nil, fmt.Errorf("unexpected ) at %d", r.position)

	default:
		r.readAtom(n, token)
	}

	return n, nil
}

func (r *Reader) Tokenize(code string) []string {
	code = strings.Replace(code, "(", "( ", -1)
	code = strings.Replace(code, ")", " )", -1)

	return strings.Fields(code)
}

func (r *Reader) peek() string {
	r.position += 1
	token := r.tokens[r.position]
	return token
}

func (r *Reader) next() string {
	return r.tokens[r.position+1]
}

func (r *Reader) readAtom(n *Node, token string) {
	if unicode.IsNumber(rune(token[0])) {
		n.Type = "number"

		val, _ := strconv.Atoi(token)
		n.Value = val
	} else {
		n.Type = "symbol"
		n.Value = token
	}
}
