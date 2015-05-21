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

	env.defs["+"] = func(nodes []*Node) *Node {
		result := 0
		for _, node := range nodes {
			result += node.Value.(int)
		}

		return &Node{Type: "number", Value: result}
	}
	env.defs["-"] = func(nodes []*Node) *Node {
		result := nodes[0].Value.(int)
		for _, node := range nodes[1:] {
			result -= node.Value.(int)
		}

		return &Node{Type: "number", Value: result}
	}
	env.defs["*"] = func(nodes []*Node) *Node {
		result := nodes[0].Value.(int)
		for _, node := range nodes[1:] {
			result *= node.Value.(int)
		}

		return &Node{Type: "number", Value: result}
	}
	env.defs["/"] = func(nodes []*Node) *Node {
		result := nodes[0].Value.(int)
		for _, node := range nodes[1:] {
			result /= node.Value.(int)
		}

		return &Node{Type: "number", Value: result}
	}

	return env
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
