package sortalgo

import (
	"container/heap"
	"sort"
)

type Shit struct {
	sort.IntSlice
}

func (s *Shit) Push(x interface{}) {
	s.IntSlice = append(s.IntSlice, x.(int))
}

func (s *Shit) Pop() interface{} {
	last := s.IntSlice[len(s.IntSlice)-1]
	s.IntSlice = s.IntSlice[:len(s.IntSlice)-1]
	return last
}

// 官方方法
func sortArrayInHeapSort1(nums []int) []int {
	shit := &Shit{}
	for _, num := range nums {
		heap.Push(shit, num)
	}
	for i := 0; i < len(nums); i++ {
		nums[i] = heap.Pop(shit).(int)
	}
	return nums
}

// 简单堆，外层往上收，内层向下堆化
func sortArrayInHeapSort2(nums []int) (ret []int) {
	var heapify func(nums []int, root, end int)
	heapify = func(nums []int, root, end int) {
		for {
			leftChild, rightChild := root*2+1, root*2+2
			if leftChild > end {
				break
			}
			curr := root
			if leftChild <= end && nums[root] > nums[leftChild] {
				root = leftChild
			}
			if rightChild <= end && nums[root] > nums[rightChild] {
				root = rightChild
			}
			if curr == root {
				break
			}
			nums[curr], nums[root] = nums[root], nums[curr]
		}
	}
	end := len(nums) - 1
	for i := end / 2; i >= 0; i-- {
		heapify(nums, i, end)
	}
	for i := end; i >= 0; i-- {
		ret = append(ret, nums[0])
		nums[i], nums[0] = nums[0], nums[i]
		end--
		heapify(nums, 0, end)
	}
	return ret
}
