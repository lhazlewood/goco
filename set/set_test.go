package set

import (
    "testing"
    . "launchpad.net/gocheck"
)

func Test(t *testing.T) { TestingT(t) }
type suite struct{}
var _ = Suite(&suite{}) 

func (s *suite) TestNew(c *C) {
    set := NewSet()
    c.Assert(set.Empty(), Equals, true)
    c.Assert(set.Size(), Equals, 0)
}

func (s *suite) TestAdd(c *C) {
    set := NewSet()
    added := set.Add(42)
    c.Assert(added, Equals, true)
}

func (s *suite) TestAddWhenExists(c *C) {
    set := NewSet()
    added := set.Add(42)
    c.Assert(added, Equals, true)
    added = set.Add(42)
    c.Assert(added, Equals, false)
}

func (s *suite) TestContains(c *C) {
    set := NewSet()
    set.Add("hello")
    c.Assert(set.Contains("hello"), Equals, true)
}

func (s *suite) TestSize(c *C) {
    set := NewSet()
    c.Assert(set.Size(), Equals, 0)
    set.Add("hello")
    c.Assert(set.Size(), Equals, 1)
}

func (s *suite) TestIterator(c *C) {
    set := NewSet()
    
    set.Add("hello")
    set.Add("world")
    
    i := set.Iterator()
    
    c.Assert(i.HasNext(), Equals, true)
    c.Assert(i.Next(), Equals, "hello")
    c.Assert(i.HasNext(), Equals, true)
    c.Assert(i.Next(), Equals, "world")
    c.Assert(i.HasNext(), Equals, false)
}








