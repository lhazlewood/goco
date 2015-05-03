package goco

import (
	"testing"
	. "launchpad.net/gocheck"
)

func Test(t *testing.T) { TestingT(t) }
type suite struct{}
var _ = Suite(&suite{})