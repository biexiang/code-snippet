package courses

import (
	. "gopkg.in/check.v1"
	"testing"
)

func Test(t *testing.T) {
	TestingT(t)
}

type MyTestSuite struct{}

var _ = Suite(&MyTestSuite{})

func (s *MyTestSuite) TestCanFinish(c *C) {
	c.Assert(canFinishDFS(2, [][]int{{0, 1}, {1, 0}}), Equals, false)
	c.Assert(canFinishBFS(2, [][]int{{0, 1}, {1, 0}}), Equals, false)
	c.Log(findOrder(2, [][]int{{1, 0}}))
	c.Log(findOrder(4, [][]int{{1, 0}, {2, 0}, {3, 1}, {3, 2}}))
	c.Assert(scheduleCourse([][]int{{1, 2}}), Equals, 1)
	c.Assert(scheduleCourse([][]int{{3, 2}, {4, 3}}), Equals, 0)
}

// https://leetcode-cn.com/problems/course-schedule/
// state: 0 means not begin
// state: 1 means begin but not finish
// state: 2 means finish
func canFinishDFS(numCourses int, prerequisites [][]int) bool {
	var (
		stateRec map[int]int
		edgeRec  map[int][]int
		dfs      func(courseID int) bool
	)
	stateRec, edgeRec = make(map[int]int, numCourses), make(map[int][]int)
	for _, item := range prerequisites {
		if edgeRec[item[1]] == nil {
			edgeRec[item[1]] = append([]int{}, item[0])
		} else {
			edgeRec[item[1]] = append(edgeRec[item[1]], item[0])
		}
	}

	dfs = func(courseID int) bool {
		stateRec[courseID] = 1
		for _, nextCourseID := range edgeRec[courseID] {
			if stateRec[nextCourseID] == 0 {
				if dfs(nextCourseID) == false {
					return false
				}
			}
			if stateRec[nextCourseID] == 1 {
				return false
			}
		}
		stateRec[courseID] = 2
		return true
	}

	for i := 0; i < numCourses; i++ {
		if stateRec[i] == 0 {
			if dfs(i) == false {
				return false
			}
		}
	}
	return true
}

// https://leetcode-cn.com/problems/course-schedule/
func canFinishBFS(numCourses int, prerequisites [][]int) bool {
	var (
		edgeRec        map[int][]int
		verticesWeight map[int]int
		cleanCourses   []int
		result         []int
	)
	edgeRec, verticesWeight = make(map[int][]int, numCourses), make(map[int]int, numCourses)
	for _, item := range prerequisites {
		// 前置课程对应的其他课程
		if edgeRec[item[1]] == nil {
			edgeRec[item[1]] = append([]int{}, item[0])
		} else {
			edgeRec[item[1]] = append(edgeRec[item[1]], item[0])
		}
		// 若课程有其他前置课程依赖，则加一
		verticesWeight[item[0]]++
	}
	for courseID, weight := range verticesWeight {
		if weight == 0 {
			cleanCourses = append(cleanCourses, courseID)
		}
	}

	// 因为第一门课程总是没有前置依赖的
	for len(cleanCourses) > 0 {
		courseID := cleanCourses[0]
		cleanCourses = cleanCourses[1:]
		result = append(result, courseID)
		for _, otherCourseID := range edgeRec[courseID] {
			if verticesWeight[otherCourseID] > 0 {
				verticesWeight[otherCourseID]--
			}
			if verticesWeight[otherCourseID] == 0 {
				cleanCourses = append(cleanCourses, otherCourseID)
			}
		}
	}
	return len(result) == numCourses
}

// https://leetcode-cn.com/problems/course-schedule-ii/
func findOrder(numCourses int, prerequisites [][]int) []int {
	var (
		edgeRec        map[int][]int
		verticesWeight map[int]int
		cleanCourses   []int
		result         []int
	)
	edgeRec, verticesWeight = make(map[int][]int, numCourses), make(map[int]int, numCourses)
	for _, item := range prerequisites {
		// 前置课程对应的其他课程
		if edgeRec[item[1]] == nil {
			edgeRec[item[1]] = append([]int{}, item[0])
		} else {
			edgeRec[item[1]] = append(edgeRec[item[1]], item[0])
		}
		// 若课程有其他前置课程依赖，则加一
		verticesWeight[item[0]]++
	}

	for i := 0; i < numCourses; i++ {
		if verticesWeight[i] == 0 {
			cleanCourses = append(cleanCourses, i)
		}
	}

	// 因为第一门课程总是没有前置依赖的
	for len(cleanCourses) > 0 {
		courseID := cleanCourses[0]
		cleanCourses = cleanCourses[1:]
		result = append(result, courseID)
		for _, otherCourseID := range edgeRec[courseID] {
			if verticesWeight[otherCourseID] > 0 {
				verticesWeight[otherCourseID]--
			}
			if verticesWeight[otherCourseID] == 0 {
				cleanCourses = append(cleanCourses, otherCourseID)
			}
		}
	}

	if len(result) != numCourses {
		return []int{}
	}
	return result
}
