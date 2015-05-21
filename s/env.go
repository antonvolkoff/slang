package s

import (
	"fmt"
)

type EnvFunc func([]*Node) *Node

type Env struct {
	defs map[string]EnvFunc
}

func NewEnv() *Env {
	env := &Env{
		defs: make(map[string]EnvFunc),
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
