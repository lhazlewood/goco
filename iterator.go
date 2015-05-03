package goco

// ==================================
// Iterator support:
// ==================================

import (
	"reflect"
	"errors"
)

type Iterator interface {

	HasNext() bool

	Next() interface{}

	MaybeNext() (element interface{}, err error)
}

type Iterable interface {

	Iterator() Iterator
}

type iterator struct {
	index int          //index of the current element in the data
	data reflect.Value //a slice or array
}

var (
	emptyArray = []int{}
	EmptyIterator Iterator = &iterator{0, reflect.ValueOf(emptyArray)}

	ErrIllegalArgument error = errors.New("iterator: argument must be a slice or array")
	ErrExhaustedElements error = errors.New("iterator: exhausted collection elements.  Check HasNext() before calling Next() to avoid this error.")
)

func newIterator(data interface{}) (i Iterator, err error) {
	if (data == nil) {
		return EmptyIterator, nil
	}

	var val reflect.Value

	aType := reflect.TypeOf(data)

	switch aType.Kind() {
	case reflect.Array: fallthrough
	case reflect.Slice:
		val = reflect.ValueOf(data)
	default:
		return nil, ErrIllegalArgument
	}

	return &iterator{0, val}, nil
}

func (self *iterator) HasNext() bool {
	l := self.data.Len()
	if self.index >= l {
		return false
	}
	return true
}

func (self *iterator) MaybeNext() (interface{}, error) {
	i := self.index
	l := self.data.Len()

	if i >= l {
		return nil, ErrExhaustedElements
	}

	val := self.data.Index(i).Interface()

	self.index++

	return val, nil
}

func (self *iterator) Next() interface{} {
	next, err := self.MaybeNext()
	if (err != nil) {
		panic(err)
	}
	return next
}