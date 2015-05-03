package goco

type mapSet struct {
    m map[interface{}]bool
}

type Set interface {
    Collection
}

func NewSet() Set {
    return &mapSet{ make( map[interface{}]bool ) }
}

var EmptySet Set = NewSet()

func SetOf(elements ...interface{}) Set {
    if (elements == nil) {
        return EmptySet
    }
    set := NewSet()
    for _, e := range elements {
        set.Add(e);
    }
    return set;
}

func SetFrom(c Collection) Set {
    if (c == nil) {
        return EmptySet
    }
    set := NewSet()
    for i := c.Iterator(); i.HasNext(); {
        set.Add(i.Next())
    }
    return set
}

func (s *mapSet) Add(e interface{}) bool {
    _, found := s.m[e]
    s.m[e] = true
    return !found
}

func (s *mapSet) AddAll(elements ...interface{}) bool {
    if (elements == nil) {
        return false
    }
    changed := false
    for _, e := range elements {
        added := s.Add(e)
        if (!changed && added) {
            changed = true
        }
    }
    return changed
}

func (s *mapSet) AddAllFrom(c Collection) bool {
    if (c == nil) {
        return false
    }
    changed := false
    for i := c.Iterator(); i.HasNext(); {
        added := s.Add(i.Next())
        if (!changed && added) {
            changed = true
        }
    }
    return changed
}

func (s *mapSet) Clear() {
    s.m = make(map[interface{}]bool)
}

func (s *mapSet) Contains(e interface{}) bool {
    _, found := s.m[e]
    return found
}

func (s *mapSet) Empty() bool {
    return len(s.m) == 0
}

func (s *mapSet) Iterator() Iterator {
    m := s.m
    l := len(m)
    a := make([]interface{}, l)
    i := 0
    for k, _ := range m {
        a[i] = k
        i++
    }

    iter, _ := newIterator(a)
    return iter;
}

func (s *mapSet) Len() int {
    return len(s.m)
}

func (s *mapSet) Remove(e interface{}) bool {
    _, found := s.m[e]
    delete(s.m, e)
    return found
}
