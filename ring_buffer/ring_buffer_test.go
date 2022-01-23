package ring_buffer

import (
	"github.com/biexiang/code-snippet/ring_buffer/lock_free_1"
	"github.com/biexiang/code-snippet/ring_buffer/lock_free_2"
	"github.com/biexiang/code-snippet/ring_buffer/no_lock"
	"testing"

	. "gopkg.in/check.v1"
)

func Test(t *testing.T) {
	TestingT(t)
}

type MyTestSuite struct{}

var _ = Suite(&MyTestSuite{})

func (s *MyTestSuite) TestNoLockRingBuffer(c *C) {
	rb := no_lock.Constructor(3)
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

func (s *MyTestSuite) TestLockFree1(c *C) {
	rb := lock_free_1.Constructor(3)
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

func (s *MyTestSuite) TestLockFree2(c *C) {
	rb := lock_free_2.Constructor(3)
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
