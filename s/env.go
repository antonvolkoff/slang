package s

import (
	"fmt"
)

type EnvFunc func([]*Node) *Node

type Env struct {
	defs   map[string]EnvFunc
	parent *Env
}

func NewEnv() *Env {
	env := &Env{
		defs:   make(map[string]EnvFunc),
		parent: nil,
	}

	return env
}

func (e *Env) Init() {
	e.defs["+"] = func(nodes []*Node) *Node {
		result := 0
		for _, node := range nodes {
			result += node.Value.(int)
		}

		return &Node{Type: "number", Value: result}
	}
	e.defs["-"] = func(nodes []*Node) *Node {
		result := nodes[0].Value.(int)
		for _, node := range nodes[1:] {
			result -= node.Value.(int)
		}

		return &Node{Type: "number", Value: result}
	}
	e.defs["*"] = func(nodes []*Node) *Node {
		result := nodes[0].Value.(int)
		for _, node := range nodes[1:] {
			result *= node.Value.(int)
		}

		return &Node{Type: "number", Value: result}
	}
	e.defs["/"] = func(nodes []*Node) *Node {
		result := nodes[0].Value.(int)
		for _, node := range nodes[1:] {
			result /= node.Value.(int)
		}

		return &Node{Type: "number", Value: result}
	}

	// List functions
	e.defs["list"] = func(nodes []*Node) *Node {
		var value []*Node
		if nodes == nil {
			value = []*Node{}
		} else {
			value = nodes
		}

		n := &Node{Type: "list", Children: value}
		return n
	}
	e.defs["list?"] = func(nodes []*Node) *Node {
		if nodes[0].Type == "list" {
			return &Node{Type: "true"}
		}
		return &Node{Type: "false"}
	}
	e.defs["empty?"] = func(nodes []*Node) *Node {
		list := nodes[0]
		if len(list.Children) == 0 {
			return &Node{Type: "true"}
		}

		return &Node{Type: "false"}
	}
	e.defs["count"] = func(nodes []*Node) *Node {
		list := nodes[0]
		count := len(list.Children)
		return &Node{Type: "number", Value: count}
	}

	// If condition
	e.defs["if"] = func(nodes []*Node) *Node {
		cond := nodes[0]
		ifTrue := nodes[1]
		var ifFalse *Node
		if len(nodes) == 3 {
			ifFalse = nodes[2]
		} else {
			ifFalse = &Node{Type: "nil"}
		}

		if cond.Type == "false" || cond.Type == "nil" {
			return ifFalse
		} else {
			return ifTrue
		}
	}

	// Basic cond
	e.defs["="] = func(nodes []*Node) *Node {
		left := nodes[0]
		right := nodes[1]

		if left.Type != left.Type {
			return &Node{Type: "false"}
		}

		switch {
		case left.Type == "list" && right.Type == "list":
			leftValue := left.Children
			rightValue := right.Children

			if len(leftValue) != len(rightValue) {
				return &Node{Type: "false"}
			}

			for i := 0; i < len(leftValue); i++ {
				if leftValue[i].Type != rightValue[i].Type || leftValue[i].Value != rightValue[i].Value {
					return &Node{Type: "false"}
				}
			}

			return &Node{Type: "true"}

		default:
			if left.Value == right.Value {
				return &Node{Type: "true"}
			} else {
				return &Node{Type: "false"}
			}
		}
	}
	e.defs[">"] = func(nodes []*Node) *Node {
		left := nodes[0].Value.(int)
		right := nodes[1].Value.(int)
		if left > right {
			return &Node{Type: "true"}
		} else {
			return &Node{Type: "false"}
		}
	}
	e.defs[">="] = func(nodes []*Node) *Node {
		left := nodes[0].Value.(int)
		right := nodes[1].Value.(int)
		if left >= right {
			return &Node{Type: "true"}
		} else {
			return &Node{Type: "false"}
		}
	}
	e.defs["<="] = func(nodes []*Node) *Node {
		left := nodes[0].Value.(int)
		right := nodes[1].Value.(int)
		if left <= right {
			return &Node{Type: "true"}
		} else {
			return &Node{Type: "false"}
		}
	}
	e.defs["<"] = func(nodes []*Node) *Node {
		left := nodes[0].Value.(int)
		right := nodes[1].Value.(int)
		if left < right {
			return &Node{Type: "true"}
		} else {
			return &Node{Type: "false"}
		}
	}
}

func (e *Env) Call(sym string, nodes []*Node) (*Node, error) {
	fn, ok := e.defs[sym]
	if !ok {
		if e.parent != nil {
			return e.parent.Call(sym, nodes)
		}
		return nil, fmt.Errorf("Undefined call to %s", sym)
	}
	return fn(nodes), nil
}

func (e *Env) Define(symbol *Node, value *Node) *Node {
	e.defs[symbol.Value.(string)] = func(nodes []*Node) *Node {
		return value
	}

	return value
}

func (e *Env) NewChild() *Env {
	env := NewEnv()
	env.parent = e
	return env
}
