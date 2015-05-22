package s

import (
	"fmt"
	"strings"
)

type Printer struct {
	item Item
}

func NewPrinter(item Item) *Printer {
	return &Printer{item: item}
}

func (p *Printer) ToString() (string, error) {
	return p.nodeToString(p.item)
}

func (p *Printer) nodeToString(i Item) (string, error) {
	var output string

	switch v := i.(type) {
	case List:
		var body []string
		for _, child := range v.Value {
			str, err := p.nodeToString(child)
			if err != nil {
				return output, err
			}
			body = append(body, str)
		}
		output = "(" + strings.Join(body, " ") + ")"

	case Hash:
		var body []string
		for _, kv := range v.Value {
			keyStr, err := p.nodeToString(kv.Key)
			if err != nil {
				return output, err
			}
			body = append(body, keyStr)

			valueStr, err := p.nodeToString(kv.Value)
			if err != nil {
				return output, err
			}
			body = append(body, valueStr)
		}
		output = "{" + strings.Join(body, " ") + "}"

	case Integer:
		output = fmt.Sprintf("%d", v.Value)

	case Symbol:
		output = fmt.Sprintf("%s", v.Value)

	case Keyword:
		output = fmt.Sprintf(":%s", v.Value)

	case Nil:
		output = "nil"

	case True:
		output = "true"

	case False:
		output = "false"

	case String:
		output = fmt.Sprintf(`"%s"`, v.Value)

	default:
		return "", fmt.Errorf("Unknown type '%s'", i)
	}

	return output, nil
}
