package s

import (
	"fmt"
)

var environment = NewEnv()

func read(input string) (Item, error) {
	r := NewReader()
	node, err := r.Parse(input)
	return node, err
}

func eval(ast Item, env *Env) (Item, error) {
	var result Item
	var err error

	fmt.Println("Eval", ast)

	switch v := ast.(type) {
	case List:
		symbol := v.Value[0].(Symbol)
		nodes := v.Value[1:]

		if symbol.Value == "def" {
			if nodes[1].IsList() {
				newNode, err := eval(nodes[1], env)
				if err != nil {
					return nil, err
				}
				nodes[1] = newNode
			}
			result = env.Define(nodes[0], nodes[1])
			break
		}

		if symbol.Value == "let" {
			childEnv := env.NewChild()

			defs := nodes[0].(Hash)
			for _, kv := range defs.Value {
				var value Item
				switch {
				case kv.Value.IsList():
					value, err = eval(kv.Value, childEnv)
					if err != nil {
						return nil, err
					}
				case kv.Value.IsSymbol():
					value, err = eval(kv.Value, childEnv)
					if err != nil {
						return nil, err
					}
				default:
					value = kv.Value
				}

				fmt.Println("Define", kv.Key, value)
				childEnv.Define(kv.Key, value)
			}

			exps := nodes[1]
			newNode, err := eval(exps, childEnv)
			if err != nil {
				return nil, err
			}

			result = newNode
			break
		}

		for idx, node := range nodes {
			if node.IsList() {
				newNode, err := eval(node, env)
				if err != nil {
					return nil, err
				}

				nodes[idx] = newNode
			}
			if node.IsSymbol() {
				newNode, err := eval(node, env)
				if err != nil {
					return nil, err
				}

				nodes[idx] = newNode
			}
		}

		fmt.Println("Call", symbol.Value, nodes)
		result, err = env.Call(symbol.Value, nodes)
		if err != nil {
			return nil, err
		}

	case Symbol:
		result, err = env.Call(v.Value, []Item{})
		if err != nil {
			return nil, err
		}

	default:
		result = ast
	}

	return result, nil
}

func print(exp Item) (string, error) {
	p := NewPrinter(exp)
	output, err := p.ToString()
	if err != nil {
		return "", err
	}
	return output, nil
}

// Read Eval Print
func Rep(input string) (string, error) {
	environment.Init()
	ast, err := read(input)
	if err != nil {
		return "", err
	}

	exp, err := eval(ast, environment)
	if err != nil {
		return "", err
	}

	output, err := print(exp)
	if err != nil {
		return "", err
	}

	return output, nil
}
