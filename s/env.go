package s

import (
	"fmt"
	// "github.com/k0kubun/pp"
)

// Env is a structure which holds environment data
type Env struct {
	defs   map[string]Item
	parent *Env
}

// NewEnv returns new environment data struct
func NewEnv() *Env {
	env := &Env{
		defs:   make(map[string]Item),
		parent: nil,
	}

	return env
}

// Init sets up main environment functions which can be executed
func (e *Env) Init() {
	e.Define("+", Func{Value: func(args []Item) (Item, error) {
		var result int64
		for _, item := range args {
			num := item.(Integer)
			result += num.Value
		}

		return Integer{Value: result}, nil
	}})

	e.Define("-", Func{Value: func(args []Item) (Item, error) {
		result := args[0].(Integer).Value
		for _, item := range args[1:] {
			result -= item.(Integer).Value
		}

		return Integer{Value: result}, nil
	}})

	e.Define("*", Func{Value: func(args []Item) (Item, error) {
		result := args[0].(Integer).Value
		for _, item := range args[1:] {
			result *= item.(Integer).Value
		}

		return Integer{Value: result}, nil
	}})

	e.Define("/", Func{Value: func(args []Item) (Item, error) {
		result := args[0].(Integer).Value
		for _, item := range args[1:] {
			result /= item.(Integer).Value
		}

		return Integer{Value: result}, nil
	}})

	e.Define("list", Func{Value: func(args []Item) (Item, error) {
		var value []Item
		if args == nil {
			value = []Item{}
		} else {
			value = args
		}

		return List{Value: value}, nil
	}})

	e.Define("list?", Func{Value: func(args []Item) (Item, error) {
		if _, ok := args[0].(List); ok {
			return True{}, nil
		}
		return False{}, nil
	}})

	e.Define("empty?", Func{Value: func(args []Item) (Item, error) {
		list := args[0].(List)
		if len(list.Value) == 0 {
			return True{}, nil
		}

		return False{}, nil
	}})

	e.Define("count", Func{Value: func(args []Item) (Item, error) {
		if !args[0].IsList() {
			return Integer{Value: 0}, nil
		}

		list := args[0].(List)
		count := int64(len(list.Value))
		return Integer{Value: count}, nil
	}})

	// Basic cond

	e.Define("=", Func{Value: func(args []Item) (Item, error) {
		left := args[0]
		right := args[1]

		if left.Equal(right).IsFalse() {
			return False{}, nil
		}

		return True{}, nil
	}})

	e.Define(">", Func{Value: func(args []Item) (Item, error) {
		left := args[0].(Integer).Value
		right := args[1].(Integer).Value
		if left > right {
			return True{}, nil
		}
		return False{}, nil
	}})

	e.Define(">=", Func{Value: func(args []Item) (Item, error) {
		left := args[0].(Integer).Value
		right := args[1].(Integer).Value
		if left >= right {
			return True{}, nil
		}
		return False{}, nil
	}})

	e.Define("<=", Func{Value: func(args []Item) (Item, error) {
		left := args[0].(Integer).Value
		right := args[1].(Integer).Value
		if left <= right {
			return True{}, nil
		}
		return False{}, nil
	}})

	e.Define("<", Func{Value: func(args []Item) (Item, error) {
		left := args[0].(Integer).Value
		right := args[1].(Integer).Value
		if left < right {
			return True{}, nil
		}

		return False{}, nil
	}})

	e.Define("not", Func{Value: func(args []Item) (Item, error) {
		val := args[0]
		if val.IsFalse() {
			return True{}, nil
		}

		return False{}, nil
	}})
}

// Define adds new function to an environment
func (e *Env) Define(name string, val Item) Item {
	e.defs[name] = val
	return val
}

func (e *Env) getRef(name string) (Item, error) {
	if item, ok := e.defs[name]; ok {
		return item, nil
	}
	return nil, fmt.Errorf("%s is undefined", name)
}

// Get return environment function
func (e *Env) Get(name string) (Item, error) {
	var item Item
	var err error

	item, err = e.getRef(name)
	if err != nil {
		if e.parent != nil {
			item, err = e.parent.Get(name)
			if err != nil {
				return nil, err
			}
		} else {
			return nil, err
		}
	}

	return item, nil
}

// NewChild creates empty child environment
func (e *Env) NewChild() *Env {
	env := NewEnv()
	env.parent = e
	return env
}
