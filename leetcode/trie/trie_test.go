package trie

import (
	"testing"

	. "gopkg.in/check.v1"
)

func Test(t *testing.T) {
	TestingT(t)
}

type MyTestSuite struct{}

var _ = Suite(&MyTestSuite{})

func (s *MyTestSuite) TestTrie(c *C) {
	t := Constructor()
	t.Insert("ileopold")
	c.Assert(t.Search("ileopold"), Equals, true)
	c.Assert(t.StartsWith("ileo"), Equals, true)
	c.Assert(t.Search("ileopolds"), Equals, false)
	c.Assert(t.Search("ileopo"), Equals, false)
	c.Assert(t.StartsWith("ileopold"), Equals, false)
	c.Assert(t.StartsWith("ileopolds"), Equals, false)
}
