package s

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSum(t *testing.T) {
	e := NewEnv()
	e.Init()

	nodes := []*Node{
		&Node{Type: "number", Value: 2},
		&Node{Type: "number", Value: 3},
		&Node{Type: "number", Value: 5},
	}

	result, err := e.Call("+", nodes)

	assert.NoError(t, err)
	assert.Equal(t, &Node{Type: "number", Value: 10}, result)
}

func TestSub(t *testing.T) {
	e := NewEnv()
	e.Init()

	nodes := []*Node{
		&Node{Type: "number", Value: 10},
		&Node{Type: "number", Value: 3},
		&Node{Type: "number", Value: 2},
	}

	result, err := e.Call("-", nodes)

	assert.NoError(t, err)
	assert.Equal(t, &Node{Type: "number", Value: 5}, result)
}

func TestMult(t *testing.T) {
	e := NewEnv()
	e.Init()

	nodes := []*Node{
		&Node{Type: "number", Value: 10},
		&Node{Type: "number", Value: 5},
		&Node{Type: "number", Value: 2},
	}

	result, err := e.Call("*", nodes)

	assert.NoError(t, err)
	assert.Equal(t, &Node{Type: "number", Value: 100}, result)
}

func TestDiv(t *testing.T) {
	e := NewEnv()
	e.Init()

	nodes := []*Node{
		&Node{Type: "number", Value: 50},
		&Node{Type: "number", Value: 5},
		&Node{Type: "number", Value: 2},
	}

	result, err := e.Call("/", nodes)

	assert.NoError(t, err)
	assert.Equal(t, &Node{Type: "number", Value: 5}, result)
}

func TestEnv_Call_ListFunctions(t *testing.T) {
	e := NewEnv()
	e.Init()

	result1, err1 := e.Call("list", []*Node{})
	assert.NoError(t, err1)
	assert.Equal(t, &Node{Type: "list"}, result1)

	result2, err2 := e.Call("list?", []*Node{&Node{Type: "list"}})
	assert.NoError(t, err2)
	assert.Equal(t, &Node{Type: "true"}, result2)

	result3, err3 := e.Call("list?", []*Node{&Node{Type: "number"}})
	assert.NoError(t, err3)
	assert.Equal(t, &Node{Type: "false"}, result3)
}

func TestEnv_Define(t *testing.T) {
	e := NewEnv()
	e.Init()

	result1 := e.Define(
		&Node{Type: "symbol", Value: "test"},
		&Node{Type: "number", Value: 42},
	)

	assert.Equal(t, &Node{Type: "number", Value: 42}, result1)

	result2, err := e.Call("test", []*Node{})

	assert.NoError(t, err)
	assert.Equal(t, &Node{Type: "number", Value: 42}, result2)
}

func TestEnv_NewChild(t *testing.T) {
	parent := NewEnv()
	parent.Init()
	child := parent.NewChild()

	assert.Equal(t, parent, child.parent)

	parent.Define(&Node{Type: "symbol", Value: "x"}, &Node{Type: "number", Value: 3})

	node, err := child.Call("x", []*Node{})

	assert.NoError(t, err)
	assert.Equal(t, &Node{Type: "number", Value: 3}, node)
}
