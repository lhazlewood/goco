package iterator

// ==================================
// Iterator support:
// ==================================

import (
    "reflect"
    . "github.com/lhazlewood/goco"   
)

type iterator struct {
    index int          //index of the current element in the data
    data reflect.Value //a slice or array
}


//global var - may be used by collections when they are empty
var emptyArray = []int{}
var EmptyIterator Iterator = NewIterator(emptyArray)

func NewIterator(data interface{}) Iterator {
    if (data == nil) {
        panic("Argument cannot be nil")
    }
    var val reflect.Value
    
    aType := reflect.TypeOf(data)
    
    switch aType.Kind() {
    case reflect.Array: fallthrough
    case reflect.Slice:
        val = reflect.ValueOf(data)
    default:
        panic("Argument must be a slice or array.")
    }
    
    return &iterator{0,val}
}

func (self *iterator) HasNext() bool {
    l := self.data.Len()
    if self.index >= l {
        return false
    }
    return true
}

func (self *iterator) Next() interface{} {
    i := self.index
    l := self.data.Len()
    
    if i >= l {
        panic("Iterator has exhausted its element set.  Check HasNext before attempting to call Next to avoid this condition.")
    }
    
    self.index++
    
    return self.data.Index(i).Interface()
}