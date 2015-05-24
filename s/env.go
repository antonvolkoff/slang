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

	// e.defs["-"] = func(items []Item) Item {
	// 	result := items[0].(Integer).Value
	// 	for _, item := range items[1:] {
	// 		result -= item.(Integer).Value
	// 	}
	//
	// 	return Integer{Value: result}
	// }
	// e.defs["*"] = func(items []Item) Item {
	// 	result := items[0].(Integer).Value
	// 	for _, item := range items[1:] {
	// 		result *= item.(Integer).Value
	// 	}
	//
	// 	return Integer{Value: result}
	// }
	// e.defs["/"] = func(items []Item) Item {
	// 	result := items[0].(Integer).Value
	// 	for _, item := range items[1:] {
	// 		result /= item.(Integer).Value
	// 	}
	//
	// 	return Integer{Value: result}
	// }
	//
	// // List functions
	// e.defs["list"] = func(items []Item) Item {
	// 	var value []Item
	// 	if items == nil {
	// 		value = []Item{}
	// 	} else {
	// 		value = items
	// 	}
	//
	// 	return List{Value: value}
	// }
	// e.defs["list?"] = func(items []Item) Item {
	// 	if _, ok := items[0].(List); ok {
	// 		return True{}
	// 	}
	// 	return False{}
	// }
	// e.defs["empty?"] = func(items []Item) Item {
	// 	list := items[0].(List)
	// 	if len(list.Value) == 0 {
	// 		return True{}
	// 	}
	//
	// 	return False{}
	// }
	// e.defs["count"] = func(items []Item) Item {
	// 	if !items[0].IsList() {
	// 		return Integer{Value: 0}
	// 	}
	//
	// 	list := items[0].(List)
	// 	count := int64(len(list.Value))
	// 	return Integer{Value: count}
	// }
	//
	// // If condition
	// e.defs["if"] = func(nodes []Item) Item {
	// 	cond := nodes[0]
	// 	ifTrue := nodes[1]
	// 	var ifFalse Item
	// 	if len(nodes) == 3 {
	// 		ifFalse = nodes[2]
	// 	} else {
	// 		ifFalse = Nil{}
	// 	}
	//
	// 	if cond.IsFalse() || cond.IsNil() {
	// 		return ifFalse
	// 	}
	//
	// 	return ifTrue
	// }
	//
	// // Basic cond
	// e.defs["="] = func(nodes []Item) Item {
	// 	left := nodes[0]
	// 	right := nodes[1]
	//
	// 	if left.Equal(right).IsFalse() {
	// 		return False{}
	// 	}
	//
	// 	return True{}
	// }
	// e.defs[">"] = func(items []Item) Item {
	// 	left := items[0].(Integer).Value
	// 	right := items[1].(Integer).Value
	// 	if left > right {
	// 		return True{}
	// 	}
	// 	return False{}
	// }
	// e.defs[">="] = func(items []Item) Item {
	// 	left := items[0].(Integer).Value
	// 	right := items[1].(Integer).Value
	// 	if left >= right {
	// 		return True{}
	// 	}
	// 	return False{}
	// }
	// e.defs["<="] = func(items []Item) Item {
	// 	left := items[0].(Integer).Value
	// 	right := items[1].(Integer).Value
	// 	if left <= right {
	// 		return True{}
	// 	}
	// 	return False{}
	// }
	// e.defs["<"] = func(items []Item) Item {
	// 	left := items[0].(Integer).Value
	// 	right := items[1].(Integer).Value
	// 	if left < right {
	// 		return True{}
	// 	}
	//
	// 	return False{}
	// }
}

// Define adds new function to an environment
func (e *Env) Define(name string, val Item) Item {
	e.defs[name] = val
	return val
}

// Get return environment function
func (e *Env) Get(name string) (Item, error) {
	var item Item
	var found bool

	item, found = e.defs[name]
	if !found {
		item, found = e.parent.defs[name]
		if !found {
			return nil, fmt.Errorf("%s is undefined", name)
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
