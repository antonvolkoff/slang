package s

import "github.com/k0kubun/pp"
import "fmt"

var environment = NewEnv()

func read(input string) (Item, error) {
	r := NewReader()
	node, err := r.Parse(input)
	return node, err
}

func evalSet(args []Item, env *Env) (Item, error) {
	name := args[0].(Symbol)
	value := args[1]

	if value.IsList() {
		var err error
		value, err = Eval(value, env)
		if err != nil {
			return nil, err
		}
	}

	env.Define(name.Value, value)

	return value, nil
}

func evalLet(args []Item, env *Env) (Item, error) {
	childEnv := env.NewChild()

	// Set env variables
	vars := args[0].(Hash)

	for _, kv := range vars.Value {
		var value Item
		var err error

		switch v := kv.Value.(type) {
		case List:
			value, err = Eval(v, childEnv)
			if err != nil {
				return nil, err
			}
		case Symbol:
			value, err = Eval(v, childEnv)
			if err != nil {
				return nil, err
			}
		default:
			value = kv.Value
		}

		name := kv.Key.(Symbol).Value
		childEnv.Define(name, value)
	}

	// Eval code inside of let
	exps := args[1]
	result, err := Eval(exps, childEnv)
	if err != nil {
		return nil, err
	}

	return result, nil
}

// Eval executes code
func Eval(root Item, env *Env) (Item, error) {
	pp.Println(root)

	switch v := root.(type) {
	case List:
		// Return empty list
		if len(v.Value) == 0 {
			return v, nil
		}

		head := v.Value[0]
		rest := v.Value[1:]

		name := "-::fn::-"
		if head.IsSymbol() {
			name = head.(Symbol).Value
		}

		switch name {
		case "fn":
			fn := Func{Value: func(args []Item) (Item, error) {
				fnEnv := env.NewChild()
				defs := rest[0].(Vector).Value
				for i, arg := range args {
					fnEnv.Define(defs[i].(Symbol).Value, arg)
				}

				return Eval(rest[1], fnEnv)
			}}

			return fn, nil

		case "set":
			return evalSet(rest, env)

		case "let":
			return evalLet(rest, env)

		default:
			fn, err := Eval(head, env)
			if err != nil {
				return nil, err
			}

			if !fn.IsFunc() {
				return nil, fmt.Errorf("Unexpected type of %v", fn)
			}

			// Transform everything to Item value
			for i, item := range rest {
				output, err := Eval(item, env)
				if err != nil {
					return nil, err
				}

				rest[i] = output
			}

			val, err := fn.(Func).Value(rest)
			if err != nil {
				return nil, err
			}

			return val, nil
		}

	case Symbol:
		return env.Get(v.Value)

	default:
		return v, nil
	}
}

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
