package main

import (
	"fmt"
	"strings"
)

type Printer struct {
	node *Node
}

func NewPrinter(node *Node) *Printer {
	return &Printer{node: node}
}

func (p *Printer) ToString() (string, error) {
	return p.nodeToString(p.node)
}

func (p *Printer) nodeToString(n *Node) (string, error) {
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

	case "number":
		output = fmt.Sprintf("%d", n.Value)

	case "symbol":
		output = fmt.Sprintf("%s", n.Value)

	default:
		return "", fmt.Errorf("Unknown type '%s'", n.Type)
	}

	return output, nil
}
