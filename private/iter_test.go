package iterator

import (
    "testing"
    . "launchpad.net/gocheck"
)

// Hook up gocheck into the "go test" runner.
func Test(t *testing.T) { TestingT(t) }
type suite struct{}
var _ = Suite(&suite{}) 

func (s *suite) TestEmptyIterator(c *C) {
    i := EmptyIterator
    c.Assert(i.HasNext(), Equals, false)
    c.Assert(func(){i.Next()}, PanicMatches, "Iterator has exhausted its element set.  Check HasNext before attempting to call Next to avoid this condition.")
}

func (s *suite) TestNewArrayIterator(c *C) {
    
    a := []string{"hello", "world"}
    
    i := NewIterator(a)
    
    for _, element := range a {
        c.Assert(i.HasNext(), Equals, true)
        c.Assert(i.Next(), Equals, element)
    }
}

func (s *suite) TestNewSliceIterator(c *C) {
    
    a := []string{"hello", "world"}
    slice := a[0:len(a)]
    
    i := NewIterator(slice)
    
    for _, element := range a {
        c.Assert(i.HasNext(), Equals, true)
        c.Assert(i.Next(), Equals, element)
    }
}
