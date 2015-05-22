package s

import (
	"fmt"
	"strings"
)

type Printer struct {
	node Item
}

func NewPrinter(node Item) *Printer {
	return &Printer{node: node}
}

func (p *Printer) ToString() (string, error) {
	return p.nodeToString(p.node)
}

func (p *Printer) nodeToString(n Item) (string, error) {
	var output string

	switch n.Type {
	case "list":
		var body []string
		for _, child := range n.Children {
			str, err := p.nodeToString(child)
			if err != nil {
				return output, err
			}
			body = append(body, str)
		}
		output = "(" + strings.Join(body, " ") + ")"

	case "hash":
		var body []string
		for key, value := range n.Value.(map[Item]Item) {
			keyStr, err := p.nodeToString(key)
			if err != nil {
				return output, err
			}
			body = append(body, keyStr)

			valueStr, err := p.nodeToString(value)
			if err != nil {
				return output, err
			}
			body = append(body, valueStr)
		}
		output = "{" + strings.Join(body, " ") + "}"

	case "number":
		output = fmt.Sprintf("%d", n.Value)

	case "symbol":
		output = fmt.Sprintf("%s", n.Value)

	case "keyword":
		output = fmt.Sprintf(":%s", n.Value)

	case "nil":
		output = "nil"

	case "true":
		output = "true"

	case "false":
		output = "false"

	case "string":
		output = fmt.Sprintf(`"%s"`, n.Value)

	default:
		return "", fmt.Errorf("Unknown type '%s'", n.Type)
	}

	return output, nil
}
