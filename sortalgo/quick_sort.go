package sortalgo

// 升序排列
// 找一个数，比它小的放在前面，大的放在后面

func middleOfThree(nums []int) {
	low, high := 0, len(nums)-1
	mid := low + (high-low)>>1
	if nums[low] > nums[high] {
		nums[low], nums[high] = nums[high], nums[low]
	}
	if nums[mid] > nums[high] {
		nums[mid], nums[high] = nums[high], nums[mid]
	}
	if nums[mid] < nums[low] {
		nums[mid], nums[low] = nums[low], nums[mid]
	}
}

func sortArrayInQuickSort1(nums []int) []int {
	return quickSort1(nums, 0, len(nums)-1)
}

// QuickSort1
func quickSort1(nums []int, start, end int) []int {
	if start < end {
		pos := partition1(nums, start, end)
		_ = quickSort1(nums, start, pos-1)
		_ = quickSort1(nums, pos+1, end)
	}
	return nums
}

// 按照pivot分区 挖坑法
// 因为pivot为左侧第一个，可以被覆盖，所以从右侧找第一个应该放到左侧的元素
func partition1(nums []int, low, high int) int {
	middleOfThree(nums)
	pivot := nums[low]
	for low < high {
		for low < high && nums[high] >= pivot {
			high--
		}
		if low < high {
			nums[low] = nums[high]
		}
		for low < high && nums[low] <= pivot {
			low++
		}
		if low < high {
			nums[high] = nums[low]
		}
	}
	nums[low] = pivot
	return low
}

func sortArrayInQuickSort2(nums []int) []int {
	return quickSort1(nums, 0, len(nums)-1)
}

// QuickSort2
func quickSort2(nums []int, start, end int) []int {
	if start < end {
		pos := partition2(nums, start, end)
		_ = quickSort1(nums, start, pos-1)
		_ = quickSort1(nums, pos+1, end)
	}
	return nums
}

// 按照pivot分区 swap法
func partition2(nums []int, low, high int) int {
	middleOfThree(nums)
	pivot := nums[low]
	first := low
	for low < high {
		for low < high && nums[high] >= pivot {
			high--
		}
		for low < high && nums[low] <= pivot {
			low++
		}
		if low < high {
			// 如果high和low都找到了，swap
			// 如果high和low都没找到，not swap
			// 如果high找到了，low没找到，not swap
			// 如果low找到，high没找到，not possible
			nums[low], nums[high] = nums[high], nums[low]
		}
	}
	nums[first], nums[low] = nums[low], nums[first]
	return low
}

// 非递归，借助栈进行迭代
func sortArrayInQuickSort3(nums []int) []int {
	// 懒得实现栈，拿切片模拟一下
	var stack []int
	stack = append(stack, 0, len(nums)-1)
	for len(stack) >= 2 {
		low, high := stack[len(stack)-2], stack[len(stack)-1]
		stack = stack[:len(stack)-2]
		if low < high {
			pos := partition1(nums, low, high)
			stack = append(stack, low, pos-1)
			stack = append(stack, pos+1, high)
		}
	}
	return nums
}

// InsertSortMaxLength 插入排序支持的最大长度
const InsertSortMaxLength = 7

// 快排+插入排序
// 当元素数量较少时，快速排序不如插入排序
func sortArrayInQuickSort4(nums []int) []int {
	return quickSort4(nums, 0, len(nums)-1)
}

func quickSort4(nums []int, start, end int) []int {
	if end-start <= InsertSortMaxLength {
		return insertSort(nums, start, end)
	}
	if start < end {
		pos := partition1(nums, start, end)
		_ = quickSort1(nums, start, pos-1)
		_ = quickSort1(nums, pos+1, end)
	}
	return nums
}
