package lock_free

import (
	"testing"

	. "gopkg.in/check.v1"
)

func Test(t *testing.T) {
	TestingT(t)
}

type MyTestSuite struct{}

var _ = Suite(&MyTestSuite{})

func (s *MyTestSuite) TestRingBuffer(c *C) {
	rb := Constructor(3)
	c.Assert(rb.EnQueue(1), Equals, true)
	c.Assert(rb.EnQueue(2), Equals, true)
	c.Assert(rb.EnQueue(3), Equals, true)
	c.Assert(rb.EnQueue(4), Equals, false)
	c.Assert(rb.EnQueue(5), Equals, false)

	value, success := rb.DeQueue()
	c.Assert(value, Equals, 1)
	c.Assert(success, Equals, true)

	value, success = rb.DeQueue()
	c.Assert(value, Equals, 2)
	c.Assert(success, Equals, true)

	value, success = rb.DeQueue()
	c.Assert(value, Equals, 3)
	c.Assert(success, Equals, true)

	value, success = rb.DeQueue()
	c.Assert(value, Equals, nil)
	c.Assert(success, Equals, false)

	value, success = rb.DeQueue()
	c.Assert(value, Equals, nil)
	c.Assert(success, Equals, false)
}
