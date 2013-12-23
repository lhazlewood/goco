package set

import (
    . "github.com/lhazlewood/goco"
    . "github.com/lhazlewood/goco/private"
)

type Set struct {
    m map[interface{}]bool
}

func NewSet() *Set {
    return &Set{ make( map[interface{}]bool ) }
}

func (s *Set) Add(e interface{}) bool {
    _, found := s.m[e]
    s.m[e] = true
    return !found
}

func (s *Set) Remove(e interface{}) bool {
    _, found := s.m[e]
    delete(s.m, e)
    return found
}

func (s *Set) Contains(e interface{}) bool {
    _, found := s.m[e]
    return found
}

func (s *Set) Size() int {
    return len(s.m)
}

func (s *Set) Empty() bool {
    return s.Size() == 0
}

func (s *Set) Iterator() Iterator {
    l := len(s.m)
    a := make([]interface{}, l)
    i := 0
    for k, _ := range s.m {
        a[i] = k
        i++
    }
    return NewIterator(a);
}


