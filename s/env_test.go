package s

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSum(t *testing.T) {
	e := NewEnv()
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
	nodes := []*Node{
		&Node{Type: "number", Value: 50},
		&Node{Type: "number", Value: 5},
		&Node{Type: "number", Value: 2},
	}

	result, err := e.Call("/", nodes)

	assert.NoError(t, err)
	assert.Equal(t, &Node{Type: "number", Value: 5}, result)
}
