package s

var environment = NewEnv()

func read(input string) (*Node, error) {
	r := NewReader()
	node, err := r.Parse(input)
	return node, err
}

func eval(ast *Node, env *Env) (*Node, error) {
	var result *Node
	var err error

	switch ast.Type {
	case "list":
		symbol := ast.Children[0].Value.(string)
		nodes := ast.Children[1:]

		if symbol == "def" {
			if nodes[1].Type == "list" {
				newNode, err := eval(nodes[1], env)
				if err != nil {
					return nil, err
				}
				nodes[1] = newNode
			}
			result = env.Define(nodes[0], nodes[1])
			break
		}

		for idx, node := range nodes {
			if node.Type == "list" {
				newNode, err := eval(node, env)
				if err != nil {
					return nil, err
				}

				nodes[idx] = newNode
			}
			if node.Type == "symbol" {
				newNode, err := eval(node, env)
				if err != nil {
					return nil, err
				}

				nodes[idx] = newNode
			}
		}

		result, err = env.Call(symbol, nodes)
		if err != nil {
			return nil, err
		}

	case "symbol":
		result, err = env.Call(ast.Value.(string), []*Node{})
		if err != nil {
			return nil, err
		}

	default:
		result = ast
	}

	return result, nil
}

func print(exp *Node) (string, error) {
	p := NewPrinter(exp)
	output, err := p.ToString()
	if err != nil {
		return "", err
	}
	return output, nil
}

// Read Eval Print
func Rep(input string) (string, error) {
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
