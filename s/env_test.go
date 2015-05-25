package s

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEnv_Define(t *testing.T) {
	e := NewEnv()

	val := Func{Value: func(items []Item) (Item, error) {
		return Integer{Value: 3}, nil
	}}
	e.Define("x", val)

	assert.Equal(t, val, e.defs["x"])
}

func TestEnv_Get(t *testing.T) {
	e := NewEnv()

	val := Func{Value: func(items []Item) (Item, error) {
		return Integer{Value: 3}, nil
	}}
	e.Define("x", val)

	item, err := e.Get("x")

	assert.NoError(t, err)
	assert.Equal(t, val, item)
}

func TestEnv_NewChild(t *testing.T) {
	parent := NewEnv()
	child := parent.NewChild()

	assert.Equal(t, parent, child.parent)
}
