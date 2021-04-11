package q35

// 思路
//	根据观察，查找等于或跟target值最近的大于target的值，所在的位置即为插入位置
//	二分查找
func searchInsert(nums []int, target int) int {
	n := len(nums)
	if n == 0 {
		return 0
	} else if n == 1 {
		if target > nums[0] {
			return 1
		}
		return 0
	}

	i, j := 0, len(nums)-1
	for i != j {
		m := i + (j-i)>>1
		if nums[m] > target {
			j = m
		} else if nums[m] < target {
			i = m + 1
		} else {
			return m
		}
	}

	if nums[j] < target {
		return j + 1
	}

	return j
}

// 对比官方解法
// 先给ans一个最大值
// 然后逐渐缩小ans
func searchInsert1(nums []int, target int) int {
	n := len(nums)
	left, right := 0, n-1
	ans := n
	for left <= right {
		mid := (right-left)>>1 + left
		if target <= nums[mid] {
			ans = mid
			right = mid - 1
		} else {
			left = mid + 1
		}
	}
	return ans
}
