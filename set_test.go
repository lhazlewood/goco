package goco

import (
	//"fmt"
	. "launchpad.net/gocheck"
)

func (s *suite) TestNewSet(c *C) {
	set := NewSet()
	c.Assert(set.Empty(), Equals, true)
	c.Assert(set.Len(), Equals, 0)
}

func (s *suite) TestEmptySet(c *C) {
	set := EmptySet
	c.Assert(set.Empty(), Equals, true)
	c.Assert(set.Len(), Equals, 0)
}

func (s *suite) TestSetOf(c *C) {

	set := SetOf(1, 2, 3, 3)

	c.Assert(set.Empty(), Equals, false)
	c.Assert(set.Len(), Equals, 3)

	c.Assert(set.Contains(1), Equals, true)
	c.Assert(set.Contains(2), Equals, true)
	c.Assert(set.Contains(3), Equals, true)
}

func (s *suite) TestSetFrom(c *C) {

	seed := SetOf(1, 2, 3)

	set := SetFrom(seed);

	set.Add(4)

	c.Assert(set.Empty(), Equals, false)
	c.Assert(set.Len(), Equals, 4)

	c.Assert(set.Contains(1), Equals, true)
	c.Assert(set.Contains(2), Equals, true)
	c.Assert(set.Contains(3), Equals, true)
	c.Assert(set.Contains(4), Equals, true)
}

func (s *suite) TestSetAdd(c *C) {

	set := NewSet()

	added := set.Add(42)
	c.Assert(added, Equals, true)
}

func (s *suite) TestSetAddWhenExists(c *C) {

	set := NewSet()

	added := set.Add(42)
	c.Assert(added, Equals, true)

	added = set.Add(42)
	c.Assert(added, Equals, false)
}

func (s *suite) TestSetAddAll(c *C) {

	set := NewSet()

	added := set.AddAll(1, 2, 3)
	c.Assert(added, Equals, true)
	c.Assert(set.Len(), Equals, 3)
	c.Assert(set.Contains(1), Equals, true)
	c.Assert(set.Contains(2), Equals, true)
	c.Assert(set.Contains(3), Equals, true)
}

func (s *suite) TestSetClear(c *C) {

	set := NewSet()

	c.Assert(set.Len(), Equals, 0)
	c.Assert(set.Empty(), Equals, true)

	set.Add(1)
	set.Add(2)
	set.Add(3)
	c.Assert(set.Len(), Equals, 3)
	c.Assert(set.Empty(), Equals, false)

	set.Clear()
	c.Assert(set.Len(), Equals, 0)
	c.Assert(set.Empty(), Equals, true)
}

func (s *suite) TestSetRemove(c *C) {

	set := NewSet()

	set.Add(42)
	removed := set.Remove(42)
	c.Assert(removed, Equals, true)
}

func (s *suite) TestSetRemoveWhenMissing(c *C) {

	set := NewSet()

	removed := set.Remove(42)
	c.Assert(removed, Equals, false)
}

func (s *suite) TestSetContains(c *C) {
	set := NewSet()
	set.Add("hello")
	c.Assert(set.Contains("hello"), Equals, true)
}

func (s *suite) TestSetContainsWhenMissing(c *C) {
	set := NewSet()
	c.Assert(set.Contains("hello"), Equals, false)
}

func (s *suite) TestSetLen(c *C) {

	set := NewSet()
	c.Assert(set.Len(), Equals, 0)

	set.Add("hello")
	c.Assert(set.Len(), Equals, 1)

	set.Remove("hello")
	c.Assert(set.Len(), Equals, 0)
}

func (s *suite) TestSetEmpty(c *C) {

	set := NewSet()
	c.Assert(set.Empty(), Equals, true)

	set.Add("hello")
	c.Assert(set.Empty(), Equals, false)

	set.Remove("hello")
	c.Assert(set.Empty(), Equals, true)
}

func (s *suite) TestSetIterator(c *C) {

	set := NewSet()

	set.Add("hello")
	set.Add("world")

	i := set.Iterator()

	c.Assert(i.HasNext(), Equals, true)
	str := i.Next().(string)
	//Default set implementation does not guarantee a deterministic iteration order, so we must check for either of the known values:
	c.Assert(str == "hello" || str == "world", Equals, true)

	c.Assert(i.HasNext(), Equals, true)
	str = i.Next().(string)
	//Default set implementation does not guarantee a deterministic iteration order, so we must check for either of the known values:
	c.Assert(str == "hello" || str == "world", Equals, true)

	c.Assert(i.HasNext(), Equals, false)
}
