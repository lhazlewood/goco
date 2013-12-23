package goco

type Iterator interface {    
    HasNext() bool
    Next() interface{}
}
