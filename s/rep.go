package s

var environment = NewEnv()

func read(input string) (*Node, error) {
	r := NewReader()
	node, err := r.Parse(input)
	return node, err
}

func eval(ast *Node, env *Env) (*Node, error) {
	var result *Node
	switch ast.Type {
	case "list":
		symbol := ast.Children[0].Value.(string)
		nodes := ast.Children[1:]

		for idx, node := range nodes {
			if node.Type == "list" {
				newNode, err := eval(node, env)
				if err != nil {
					return nil, err
				}

				nodes[idx] = newNode
			}
		}

		var err error
		result, err = env.Call(symbol, nodes)
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
