package sortalgo

import (
	"strconv"
	"testing"

	. "gopkg.in/check.v1"
)

func Test(t *testing.T) {
	TestingT(t)
}

type MyTestSuite struct{}

var _ = Suite(&MyTestSuite{})

func joinInt(sl []int) (s string) {
	for _, i := range sl {
		s = s + strconv.Itoa(i)
	}
	return
}

func (s *MyTestSuite) TestQuickSort(c *C) {
	c.Assert(joinInt(sortArrayInQuickSort1([]int{5, 2, 3, 1})), Equals, joinInt([]int{1, 2, 3, 5}))
	c.Assert(joinInt(sortArrayInQuickSort1([]int{5, 1, 1, 2, 0, 0})), Equals, joinInt([]int{0, 0, 1, 1, 2, 5}))

	c.Assert(joinInt(sortArrayInQuickSort2([]int{5, 2, 3, 1})), Equals, joinInt([]int{1, 2, 3, 5}))
	c.Assert(joinInt(sortArrayInQuickSort2([]int{5, 1, 1, 2, 0, 0})), Equals, joinInt([]int{0, 0, 1, 1, 2, 5}))

	c.Assert(joinInt(sortArrayInQuickSort3([]int{5, 2, 3, 1})), Equals, joinInt([]int{1, 2, 3, 5}))
	c.Assert(joinInt(sortArrayInQuickSort3([]int{5, 1, 1, 2, 0, 0})), Equals, joinInt([]int{0, 0, 1, 1, 2, 5}))

	c.Assert(joinInt(sortArrayInQuickSort4([]int{5, 2, 3, 1})), Equals, joinInt([]int{1, 2, 3, 5}))
	c.Assert(joinInt(sortArrayInQuickSort4([]int{5, 1, 1, 2, 0, 0})), Equals, joinInt([]int{0, 0, 1, 1, 2, 5}))
	c.Assert(joinInt(sortArrayInQuickSort4([]int{5, 1, 1, 2, 0, 0, 10, 9, 8, 7, 6})), Equals, joinInt([]int{0, 0, 1, 1, 2, 5, 6, 7, 8, 9, 10}))
}

func (s *MyTestSuite) TestInsertSort(c *C) {
	c.Assert(joinInt(sortArrayInInsertSort([]int{5, 2, 3, 1})), Equals, joinInt([]int{1, 2, 3, 5}))
	c.Assert(joinInt(sortArrayInInsertSort([]int{5, 1, 1, 2, 0, 0})), Equals, joinInt([]int{0, 0, 1, 1, 2, 5}))

	c.Assert(joinInt(sortArrayInShellSort([]int{5, 2, 3, 1})), Equals, joinInt([]int{1, 2, 3, 5}))
	c.Assert(joinInt(sortArrayInShellSort([]int{5, 1, 1, 2, 0, 0})), Equals, joinInt([]int{0, 0, 1, 1, 2, 5}))
}

func (s *MyTestSuite) TestHeapSort(c *C) {
	c.Assert(joinInt(sortArrayInHeapSort1([]int{5, 2, 3, 1})), Equals, joinInt([]int{1, 2, 3, 5}))
	c.Assert(joinInt(sortArrayInHeapSort1([]int{5, 1, 1, 2, 0, 0})), Equals, joinInt([]int{0, 0, 1, 1, 2, 5}))

	c.Assert(joinInt(sortArrayInHeapSort2([]int{5, 2, 3, 1})), Equals, joinInt([]int{1, 2, 3, 5}))
	c.Assert(joinInt(sortArrayInHeapSort2([]int{5, 1, 1, 2, 0, 0})), Equals, joinInt([]int{0, 0, 1, 1, 2, 5}))
}
