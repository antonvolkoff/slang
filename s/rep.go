package s

// import "github.com/k0kubun/pp"

var environment = NewEnv()

func read(input string) (Item, error) {
	r := NewReader()
	node, err := r.Parse(input)
	return node, err
}

// Eval executes code
func Eval(root Item, env *Env) (Item, error) {

	switch v := root.(type) {
	case List:
		// Return empty list
		if len(v.Value) == 0 {
			return v, nil
		}

	case Symbol:
		val, err := env.Get(v.Value)
		return val, err

	default:
		return v, nil
	}

	return nil, nil
}

// func eval(ast Item, env *Env) (Item, error) {
// 	pp.Println("Eval", ast)
// 	var result Item
// 	var err error
//
// 	switch v := ast.(type) {
// 	case List:
// 		var symbol Symbol
// 		if v.Value[0].IsList() {
// 			item, err := eval(v.Value[0], env)
// 			if err != nil {
// 				return nil, err
// 			}
// 			return eval(item, env)
// 		}
//
// 		symbol = v.Value[0].(Symbol)
// 		nodes := v.Value[1:]
//
// 		if symbol.Value == "def" {
// 			if nodes[1].IsList() {
// 				newNode, err := eval(nodes[1], env)
// 				if err != nil {
// 					return nil, err
// 				}
// 				nodes[1] = newNode
// 			}
// 			result = env.Define(nodes[0], nodes[1])
// 			break
// 		}
//
// 		if symbol.Value == "let" {
// 			childEnv := env.NewChild()
//
// 			defs := nodes[0].(Hash)
// 			for _, kv := range defs.Value {
// 				var value Item
// 				switch {
// 				case kv.Value.IsList():
// 					value, err = eval(kv.Value, childEnv)
// 					if err != nil {
// 						return nil, err
// 					}
// 				case kv.Value.IsSymbol():
// 					value, err = eval(kv.Value, childEnv)
// 					if err != nil {
// 						return nil, err
// 					}
// 				default:
// 					value = kv.Value
// 				}
//
// 				childEnv.Define(kv.Key, value)
// 			}
//
// 			exps := nodes[1]
// 			newNode, err := eval(exps, childEnv)
// 			if err != nil {
// 				return nil, err
// 			}
//
// 			result = newNode
// 			break
// 		}
//
// 		if symbol.Value == "fn" {
// 			fnName := Symbol{Value: "__fn__"}
// 			env.DefineFn(fnName, func(args []Item) Item {
// 				fnEnv := env.NewChild()
//
// 				if len(args) > 0 {
// 					names := nodes[0].(Vector)
// 					for idx, item := range args {
// 						fnEnv.Define(names.Value[idx], item)
// 					}
// 				}
//
// 				ret, _ := eval(nodes[1], fnEnv)
// 				return ret
// 			})
//
// 			result = fnName
// 			break
// 		}
//
// 		for idx, node := range nodes {
// 			if node.IsList() {
// 				newNode, err := eval(node, env)
// 				if err != nil {
// 					return nil, err
// 				}
//
// 				nodes[idx] = newNode
// 			}
// 			if node.IsSymbol() {
// 				newNode, err := eval(node, env)
// 				if err != nil {
// 					return nil, err
// 				}
//
// 				nodes[idx] = newNode
// 			}
// 		}
//
// 		result, err = env.Call(symbol.Value, nodes)
// 		if err != nil {
// 			return nil, err
// 		}
//
// 	case Symbol:
// 		result, err = env.Call(v.Value, []Item{})
// 		if err != nil {
// 			return nil, err
// 		}
//
// 	default:
// 		result = ast
// 	}
//
// 	return result, nil
// }

func print(exp Item) (string, error) {
	p := NewPrinter(exp)
	output, err := p.ToString()
	if err != nil {
		return "", err
	}
	return output, nil
}

// Rep is an read-eval-print implementation
func Rep(input string) (string, error) {
	environment.Init()
	ast, err := read(input)
	if err != nil {
		return "", err
	}

	exp, err := Eval(ast, environment)
	if err != nil {
		return "", err
	}

	output, err := print(exp)
	if err != nil {
		return "", err
	}

	return output, nil
}
