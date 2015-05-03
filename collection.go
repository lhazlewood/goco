package goco

type Collection interface {
	Iterable

	Add(element interface{}) bool

	AddAll(elements ...interface{}) bool

	AddAllFrom(c Collection) bool

	Clear()

	//Collect(f func(interface{}) interface{} ) Collection

	Contains(element interface{}) bool

	//Each(f func(interface{}) )

	Empty() bool

	Len() int

	Remove(element interface{}) bool

	//Slice() []interface{}
}
