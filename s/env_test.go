package s

// import (
// 	"testing"
//
// 	"github.com/stretchr/testify/assert"
// )
//
// func TestSum(t *testing.T) {
// 	e := NewEnv()
// 	e.Init()
//
// 	nodes := []Item{
// 		Integer{Value: 2},
// 		Integer{Value: 3},
// 		Integer{Value: 5},
// 	}
//
// 	result, err := e.Call("+", nodes)
//
// 	assert.NoError(t, err)
// 	assert.Equal(t, Integer{Value: 10}, result)
// }
//
// func TestSub(t *testing.T) {
// 	e := NewEnv()
// 	e.Init()
//
// 	nodes := []Item{
// 		Integer{Value: 10},
// 		Integer{Value: 3},
// 		Integer{Value: 2},
// 	}
//
// 	result, err := e.Call("-", nodes)
//
// 	assert.NoError(t, err)
// 	assert.Equal(t, Integer{Value: 5}, result)
// }
//
// func TestMult(t *testing.T) {
// 	e := NewEnv()
// 	e.Init()
//
// 	nodes := []Item{
// 		Integer{Value: 10},
// 		Integer{Value: 5},
// 		Integer{Value: 2},
// 	}
//
// 	result, err := e.Call("*", nodes)
//
// 	assert.NoError(t, err)
// 	assert.Equal(t, Integer{Value: 100}, result)
// }
//
// func TestDiv(t *testing.T) {
// 	e := NewEnv()
// 	e.Init()
//
// 	nodes := []Item{
// 		Integer{Value: 50},
// 		Integer{Value: 5},
// 		Integer{Value: 2},
// 	}
//
// 	result, err := e.Call("/", nodes)
//
// 	assert.NoError(t, err)
// 	assert.Equal(t, Integer{Value: 5}, result)
// }
//
// func TestEnv_Call_ListFunctions(t *testing.T) {
// 	e := NewEnv()
// 	e.Init()
//
// 	result1, err1 := e.Call("list", []Item{})
// 	assert.NoError(t, err1)
// 	assert.Equal(t, List{Value: []Item{}}, result1)
//
// 	result2, err2 := e.Call("list", []Item{
// 		Integer{Value: 1}, Integer{Value: 2}})
// 	assert.NoError(t, err2)
// 	expected := List{Value: []Item{
// 		Integer{Value: 1}, Integer{Value: 2}}}
// 	assert.Equal(t, expected, result2)
//
// 	result3, err3 := e.Call("list?", []Item{List{}})
// 	assert.NoError(t, err3)
// 	assert.Equal(t, True{}, result3)
//
// 	result4, err4 := e.Call("list?", []Item{Integer{}})
// 	assert.NoError(t, err4)
// 	assert.Equal(t, False{}, result4)
// }
//
// func TestEnv_Call_Empty(t *testing.T) {
// 	e := NewEnv()
// 	e.Init()
//
// 	result1, err1 := e.Call("empty?", []Item{
// 		List{Value: []Item{}},
// 	})
// 	assert.NoError(t, err1)
// 	assert.Equal(t, True{}, result1)
//
// 	result2, err2 := e.Call("empty?", []Item{
// 		List{Value: []Item{Integer{Value: 1}}},
// 	})
// 	assert.NoError(t, err2)
// 	assert.Equal(t, False{}, result2)
// }
//
// func TestEnv_Call_Count(t *testing.T) {
// 	e := NewEnv()
// 	e.Init()
//
// 	result1, err1 := e.Call("count", []Item{
// 		List{Value: []Item{}},
// 	})
// 	assert.NoError(t, err1)
// 	assert.Equal(t, Integer{Value: 0}, result1)
//
// 	result2, err2 := e.Call("count", []Item{
// 		List{Value: []Item{Integer{Value: 1}}},
// 	})
// 	assert.NoError(t, err2)
// 	assert.Equal(t, Integer{Value: 1}, result2)
// }
//
// func TestEnv_Call_If(t *testing.T) {
// 	e := NewEnv()
// 	e.Init()
//
// 	result1, err1 := e.Call("if", []Item{
// 		True{},
// 		Integer{Value: 1},
// 		Integer{Value: 2},
// 	})
// 	assert.NoError(t, err1)
// 	assert.Equal(t, Integer{Value: 1}, result1)
//
// 	result2, err2 := e.Call("if", []Item{
// 		False{},
// 		Integer{Value: 1},
// 		Integer{Value: 2},
// 	})
// 	assert.NoError(t, err2)
// 	assert.Equal(t, Integer{Value: 2}, result2)
//
// 	result3, err3 := e.Call("if", []Item{
// 		Nil{},
// 		Integer{Value: 1},
// 		Integer{Value: 2},
// 	})
// 	assert.NoError(t, err3)
// 	assert.Equal(t, Integer{Value: 2}, result3)
//
// 	result4, err4 := e.Call("if", []Item{
// 		Nil{},
// 		Integer{Value: 1},
// 	})
// 	assert.NoError(t, err4)
// 	assert.Equal(t, Nil{}, result4)
// }
//
// func TestEnv_Call_Equal(t *testing.T) {
// 	e := NewEnv()
// 	e.Init()
//
// 	result1, err := e.Call("=", []Item{
// 		Integer{Value: 1},
// 		Integer{Value: 2},
// 	})
// 	assert.NoError(t, err)
// 	assert.Equal(t, False{}, result1)
//
// 	result2, err := e.Call("=", []Item{
// 		Integer{Value: 1},
// 		Integer{Value: 1},
// 	})
// 	assert.NoError(t, err)
// 	assert.Equal(t, True{}, result2)
// }
//
// func TestEnv_Call_More(t *testing.T) {
// 	e := NewEnv()
// 	e.Init()
//
// 	result1, err := e.Call(">", []Item{
// 		Integer{Value: 1},
// 		Integer{Value: 2},
// 	})
// 	assert.NoError(t, err)
// 	assert.Equal(t, False{}, result1)
//
// 	result2, err := e.Call(">", []Item{
// 		Integer{Value: 2},
// 		Integer{Value: 1},
// 	})
// 	assert.NoError(t, err)
// 	assert.Equal(t, True{}, result2)
// }
//
// func TestEnv_Call_MoreOrEqual(t *testing.T) {
// 	e := NewEnv()
// 	e.Init()
//
// 	result1, err := e.Call(">=", []Item{
// 		Integer{Value: 1},
// 		Integer{Value: 2},
// 	})
// 	assert.NoError(t, err)
// 	assert.Equal(t, False{}, result1)
//
// 	result2, err := e.Call(">=", []Item{
// 		Integer{Value: 2},
// 		Integer{Value: 1},
// 	})
// 	assert.NoError(t, err)
// 	assert.Equal(t, True{}, result2)
//
// 	result3, err := e.Call(">=", []Item{
// 		Integer{Value: 1},
// 		Integer{Value: 1},
// 	})
// 	assert.NoError(t, err)
// 	assert.Equal(t, True{}, result3)
// }
//
// func TestEnv_Call_LessOrEqual(t *testing.T) {
// 	e := NewEnv()
// 	e.Init()
//
// 	result1, err := e.Call("<=", []Item{
// 		Integer{Value: 1},
// 		Integer{Value: 2},
// 	})
// 	assert.NoError(t, err)
// 	assert.Equal(t, True{}, result1)
//
// 	result2, err := e.Call("<=", []Item{
// 		Integer{Value: 2},
// 		Integer{Value: 1},
// 	})
// 	assert.NoError(t, err)
// 	assert.Equal(t, False{}, result2)
//
// 	result3, err := e.Call("<=", []Item{
// 		Integer{Value: 1},
// 		Integer{Value: 1},
// 	})
// 	assert.NoError(t, err)
// 	assert.Equal(t, True{}, result3)
// }
//
// func TestEnv_Call_Less(t *testing.T) {
// 	e := NewEnv()
// 	e.Init()
//
// 	result1, err := e.Call("<", []Item{
// 		Integer{Value: 1},
// 		Integer{Value: 2},
// 	})
// 	assert.NoError(t, err)
// 	assert.Equal(t, True{}, result1)
//
// 	result2, err := e.Call("<", []Item{
// 		Integer{Value: 2},
// 		Integer{Value: 1},
// 	})
// 	assert.NoError(t, err)
// 	assert.Equal(t, False{}, result2)
//
// 	result3, err := e.Call("<", []Item{
// 		Integer{Value: 1},
// 		Integer{Value: 1},
// 	})
// 	assert.NoError(t, err)
// 	assert.Equal(t, False{}, result3)
// }
//
// func TestEnv_Define(t *testing.T) {
// 	e := NewEnv()
// 	e.Init()
//
// 	result1 := e.Define(
// 		Symbol{Value: "test"},
// 		Integer{Value: 42},
// 	)
//
// 	assert.Equal(t, Integer{Value: 42}, result1)
//
// 	result2, err := e.Call("test", []Item{})
//
// 	assert.NoError(t, err)
// 	assert.Equal(t, Integer{Value: 42}, result2)
// }
//
// func TestEnv_NewChild(t *testing.T) {
// 	parent := NewEnv()
// 	parent.Init()
// 	child := parent.NewChild()
//
// 	assert.Equal(t, parent, child.parent)
//
// 	parent.Define(Symbol{Value: "x"}, Integer{Value: 3})
//
// 	node, err := child.Call("x", []Item{})
//
// 	assert.NoError(t, err)
// 	assert.Equal(t, Integer{Value: 3}, node)
// }
