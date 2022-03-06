package sortalgo

// 升序排序

func sortArrayInInsertSort(nums []int) []int {
	return insertSort(nums, 0, len(nums)-1)
}

func sortArrayInShellSort(nums []int) []int {
	return shellSort(nums, 0, len(nums)-1)
}

// 稳定算法，比较交换，时间O(n^2)，空间O(1)
// 数据规模小的时候，或基本有序，效率高
func insertSort(nums []int, low, high int) []int {
	for i := low; i <= high; i++ {
		for j := i; j > low; j-- {
			if nums[j] < nums[j-1] {
				nums[j], nums[j-1] = nums[j-1], nums[j]
			}
		}
	}
	return nums
}

// 不稳定算法，比较交换，时间O(n^2)，空间O(1)
// 有点不好理解，待排序列表分段，每段按照索引和前面的段进行插入排序
func shellSort(nums []int, low, high int) []int {
	gap := (high - low + 1) / 2
	for gap > 0 {
		for i := gap; i <= high; i++ {
			j := i
			for j-gap >= low && nums[j] < nums[j-gap] {
				nums[j], nums[j-gap] = nums[j-gap], nums[j]
				j = j - gap
			}
		}
		gap = gap >> 1
	}
	return nums
}
