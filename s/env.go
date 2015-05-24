package s

import (
	"fmt"
)

// EnvFunc is a type definition of environment function
type EnvFunc func([]Item) Item

// Env is a structure which holds environment data
type Env struct {
	defs   map[string]EnvFunc
	parent *Env
}

// NewEnv returns new environment data struct
func NewEnv() *Env {
	env := &Env{
		defs:   make(map[string]EnvFunc),
		parent: nil,
	}

	return env
}

// Init sets up main environment functions which can be executed
func (e *Env) Init() {
	e.defs["+"] = func(items []Item) Item {
		var result int64
		for _, item := range items {
			num := item.(Integer)
			result += num.Value
		}

		return Integer{Value: result}
	}
	e.defs["-"] = func(items []Item) Item {
		result := items[0].(Integer).Value
		for _, item := range items[1:] {
			result -= item.(Integer).Value
		}

		return Integer{Value: result}
	}
	e.defs["*"] = func(items []Item) Item {
		result := items[0].(Integer).Value
		for _, item := range items[1:] {
			result *= item.(Integer).Value
		}

		return Integer{Value: result}
	}
	e.defs["/"] = func(items []Item) Item {
		result := items[0].(Integer).Value
		for _, item := range items[1:] {
			result /= item.(Integer).Value
		}

		return Integer{Value: result}
	}

	// List functions
	e.defs["list"] = func(items []Item) Item {
		var value []Item
		if items == nil {
			value = []Item{}
		} else {
			value = items
		}

		return List{Value: value}
	}
	e.defs["list?"] = func(items []Item) Item {
		if _, ok := items[0].(List); ok {
			return True{}
		}
		return False{}
	}
	e.defs["empty?"] = func(items []Item) Item {
		list := items[0].(List)
		if len(list.Value) == 0 {
			return True{}
		}

		return False{}
	}
	e.defs["count"] = func(items []Item) Item {
		if !items[0].IsList() {
			return Integer{Value: 0}
		}

		list := items[0].(List)
		count := int64(len(list.Value))
		return Integer{Value: count}
	}

	// If condition
	e.defs["if"] = func(nodes []Item) Item {
		cond := nodes[0]
		ifTrue := nodes[1]
		var ifFalse Item
		if len(nodes) == 3 {
			ifFalse = nodes[2]
		} else {
			ifFalse = Nil{}
		}

		if cond.IsFalse() || cond.IsNil() {
			return ifFalse
		}

		return ifTrue
	}

	// Basic cond
	e.defs["="] = func(nodes []Item) Item {
		left := nodes[0]
		right := nodes[1]

		if left.Equal(right).IsFalse() {
			return False{}
		}

		return True{}
	}
	e.defs[">"] = func(items []Item) Item {
		left := items[0].(Integer).Value
		right := items[1].(Integer).Value
		if left > right {
			return True{}
		}
		return False{}
	}
	e.defs[">="] = func(items []Item) Item {
		left := items[0].(Integer).Value
		right := items[1].(Integer).Value
		if left >= right {
			return True{}
		}
		return False{}
	}
	e.defs["<="] = func(items []Item) Item {
		left := items[0].(Integer).Value
		right := items[1].(Integer).Value
		if left <= right {
			return True{}
		}
		return False{}
	}
	e.defs["<"] = func(items []Item) Item {
		left := items[0].(Integer).Value
		right := items[1].(Integer).Value
		if left < right {
			return True{}
		}

		return False{}
	}
}

// Call executes environment function by given name and arguments
func (e *Env) Call(sym string, nodes []Item) (Item, error) {
	fn, ok := e.defs[sym]
	if !ok {
		if e.parent != nil {
			return e.parent.Call(sym, nodes)
		}
		return nil, fmt.Errorf("Undefined call to %s", sym)
	}
	return fn(nodes), nil
}

// Define adds new function to an environment
func (e *Env) Define(symbol Item, value Item) Item {
	e.defs[symbol.(Symbol).Value] = func(items []Item) Item {
		return value
	}

	return value
}

// DefineFn creates new function
func (e *Env) DefineFn(symbol Symbol, fn EnvFunc) {
	e.defs[symbol.Value] = fn
}

// NewChild creates empty child environment
func (e *Env) NewChild() *Env {
	env := NewEnv()
	env.parent = e
	return env
}
