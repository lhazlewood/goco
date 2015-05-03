package goco

import (
    . "launchpad.net/gocheck"
)

func assertNextExhausted(i Iterator, c *C) {
    next, err := i.MaybeNext()
    c.Assert(next, Equals, nil)
    c.Assert(err, Equals, ErrExhaustedElements)
    c.Assert(func(){i.Next()}, Panics, ErrExhaustedElements)
}

func (s *suite) TestNewWithNilArgument(c *C) {
    i, _ := newIterator(nil)
    c.Assert(i, NotNil)
    c.Assert(i.HasNext(), Equals, false)
    assertNextExhausted(i, c)
}

func (s *suite) TestEmptyIterator(c *C) {
    i := EmptyIterator
    c.Assert(i.HasNext(), Equals, false)
    assertNextExhausted(i, c)
}

func (s *suite) TestNewArrayIterator(c *C) {
    
    a := []string{"hello", "world"}
    
    i, _ := newIterator(a)

    c.Assert(i.Next(), Equals, "hello")
    c.Assert(i.Next(), Equals, "world")
    c.Assert(i.HasNext(), Equals, false)

    assertNextExhausted(i, c)
}

func (s *suite) TestNewSliceIterator(c *C) {
    
    a := []string{"hello", "world"}
    slice := a[0:len(a)]
    
    i, _ := newIterator(slice)

    c.Assert(i.Next(), Equals, "hello")
    c.Assert(i.Next(), Equals, "world")
    c.Assert(i.HasNext(), Equals, false)

    assertNextExhausted(i, c)
}
