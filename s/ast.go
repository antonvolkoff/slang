package s

type Item interface {
	Equal(Item) Item
}

////////////////////////////////////////////////////////////////////////////////

// True is a `true` type of slang
type True struct{}

func (self True) Equal(i Item) Item {
	switch i.(type) {
	case True:
		return True{}
	default:
		return False{}
	}
}

////////////////////////////////////////////////////////////////////////////////

// False is a `false` type of slang
type False struct{}

func (self False) Equal(i Item) Item {
	switch i.(type) {
	case False:
		return True{}
	default:
		return False{}
	}
}

////////////////////////////////////////////////////////////////////////////////

// Nil is a `nil` type of slang
type Nil struct{}

func (self Nil) Equal(i Item) Item {
	switch i.(type) {
	case Nil:
		return True{}
	default:
		return False{}
	}
}

////////////////////////////////////////////////////////////////////////////////

type Integer struct {
	Value int64
}

func (self Integer) Equal(i Item) Item {
	switch v := i.(type) {
	case Integer:
		if self.Value != v.Value {
			return False{}
		}
		return True{}

	default:
		return False{}
	}
}

////////////////////////////////////////////////////////////////////////////////

type String struct {
	Value string
}

func (self String) Equal(i Item) Item {
	switch v := i.(type) {
	case String:
		if self.Value != v.Value {
			return False{}
		}
		return True{}

	default:
		return False{}
	}
}

////////////////////////////////////////////////////////////////////////////////

type Symbol struct {
	Value string
}

func (self Symbol) Equal(i Item) Item {
	switch v := i.(type) {
	case Symbol:
		if self.Value != v.Value {
			return False{}
		}
		return True{}

	default:
		return False{}
	}
}

////////////////////////////////////////////////////////////////////////////////

type Keyword struct {
	Value string
}

func (self Keyword) Equal(i Item) Item {
	switch v := i.(type) {
	case Keyword:
		if self.Value != v.Value {
			return False{}
		}
		return True{}

	default:
		return False{}
	}
}

////////////////////////////////////////////////////////////////////////////////

type List struct {
	Value []Item
}

func (self List) Equal(i Item) Item {
	switch v := i.(type) {
	case List:
		if len(v.Value) != len(self.Value) {
			return False{}
		}

		for i, elem := range self.Value {
			if _, ok := elem.Equal(v.Value[i]).(False); ok {
				return False{}
			}
		}

		return True{}

	default:
		return False{}
	}
}

func (self List) Add(i Item) {
	self.Value = append(self.Value, i)
}

////////////////////////////////////////////////////////////////////////////////

type KeyValue struct {
	Key   Item
	Value Item
}

func (self KeyValue) Equal(i Item) Item {
	switch v := i.(type) {
	case KeyValue:
		if _, ok := self.Key.Equal(v.Key).(False); ok {
			return False{}
		}
		if _, ok := self.Value.Equal(v.Value).(False); ok {
			return False{}
		}
		return True{}

	default:
		return False{}
	}
}

////////////////////////////////////////////////////////////////////////////////

type Map struct {
	Value []KeyValue
}

func (self Map) Equal(i Item) Item {
	switch v := i.(type) {
	case Map:
		if len(v.Value) != len(self.Value) {
			return False{}
		}

		for i, elem := range self.Value {
			if _, ok := elem.Equal(v.Value[i]).(False); ok {
				return False{}
			}
		}

		return True{}

	default:
		return False{}
	}
}

func (self Map) Add(kv KeyValue) {
	self.Value = append(self.Value, kv)
}

////////////////////////////////////////////////////////////////////////////////

type Vector struct {
	Value []Item
}

func (self Vector) Equal(i Item) Item {
	switch v := i.(type) {
	case Vector:
		if len(v.Value) != len(self.Value) {
			return False{}
		}

		for i, elem := range self.Value {
			if _, ok := elem.Equal(v.Value[i]).(False); ok {
				return False{}
			}
		}

		return True{}

	default:
		return False{}
	}
}

func (self Vector) Add(i Item) {
	self.Value = append(self.Value, i)
}
