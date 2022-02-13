package courses

import (
	"container/heap"
	"sort"
)

type Heap struct {
	sort.IntSlice
}

func (h *Heap) Less(i, j int) bool {
	return h.IntSlice[i] > h.IntSlice[j]
}

func (h *Heap) Push(x interface{}) {
	h.IntSlice = append(h.IntSlice, x.(int))
}

func (h *Heap) Pop() interface{} {
	tmp := h.IntSlice
	last := tmp[len(tmp)-1]
	h.IntSlice = tmp[:len(tmp)-1]
	return last
}

// https://leetcode-cn.com/problems/course-schedule-iii/
// Greedy Algorithm
func scheduleCourse(courses [][]int) int {
	// 按照deadline排个序
	sort.Slice(courses, func(i, j int) bool {
		return courses[i][1] < courses[j][1]
	})

	h := &Heap{}
	total := 0
	// 遍历有序courses，不满足deadline的course和用时最长的比较，如果比较长，则替换
	for _, course := range courses {
		if total+course[0] <= course[1] {
			total += course[0]
			heap.Push(h, course[0])
		} else if h.Len() > 0 && course[0] < h.IntSlice[0] {
			total = total - h.IntSlice[0] + course[0]
			h.IntSlice[0] = course[0]
			heap.Fix(h, 0)
		}
	}
	return h.Len()
}
