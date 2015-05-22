package s

type Item interface {
	Equal(Item) Item
	IsTrue() bool
	IsFalse() bool
	IsNil() bool
	IsInteger() bool
	IsString() bool
	IsSymbol() bool
	IsKeyword() bool
	IsList() bool
	IsHash() bool
	IsVector() bool
}

type DefaultItem struct{}

func (self DefaultItem) IsTrue() bool {
	return false
}

func (self DefaultItem) IsFalse() bool {
	return false
}

func (self DefaultItem) IsNil() bool {
	return false
}

func (self DefaultItem) IsInteger() bool {
	return false
}

func (self DefaultItem) IsString() bool {
	return false
}

func (self DefaultItem) IsSymbol() bool {
	return false
}

func (self DefaultItem) IsKeyword() bool {
	return false
}

func (self DefaultItem) IsList() bool {
	return false
}

func (self DefaultItem) IsHash() bool {
	return false
}

func (self DefaultItem) IsVector() bool {
	return false
}

////////////////////////////////////////////////////////////////////////////////

// True is a `true` type of slang
type True struct {
	DefaultItem
}

func (self True) IsTrue() bool {
	return true
}

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
type False struct {
	DefaultItem
}

func (self False) IsFalse() bool {
	return true
}

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
type Nil struct {
	DefaultItem
}

func (self Nil) IsNil() bool {
	return true
}

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
	DefaultItem
	Value int64
}

func (self Integer) IsInteger() bool {
	return true
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
	DefaultItem
	Value string
}

func (self String) IsString() bool {
	return true
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
	DefaultItem
	Value string
}

func (self Symbol) IsSymbol() bool {
	return true
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
	DefaultItem
	Value string
}

func (self Keyword) IsKeyword() bool {
	return true
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
	DefaultItem
	Value []Item
}

func (self List) IsList() bool {
	return true
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

func (self List) Add(i Item) List {
	self.Value = append(self.Value, i)
	return self
}

////////////////////////////////////////////////////////////////////////////////

type KeyValue struct {
	DefaultItem
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

type Hash struct {
	DefaultItem
	Value []KeyValue
}

func (self Hash) IsHash() bool {
	return true
}

func (self Hash) Equal(i Item) Item {
	switch v := i.(type) {
	case Hash:
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

func (self Hash) Add(kv KeyValue) Hash {
	self.Value = append(self.Value, kv)
	return self
}

////////////////////////////////////////////////////////////////////////////////

type Vector struct {
	DefaultItem
	Value []Item
}

func (self Vector) IsVector() bool {
	return true
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

func (self Vector) Add(i Item) Vector {
	self.Value = append(self.Value, i)
	return self
}
