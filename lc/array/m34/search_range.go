package m34

// 思路
//	找左右边界，先找左边界，找不到就直接返回-1，-1，然后找右边界
//	注意边界值处理
func searchRange(nums []int, target int) []int {

	// find left
	n := len(nums)
	if n == 0 {
		return []int{-1, -1}
	}

	left, right := 0, n-1
	for left < right {
		mid := left + (right-left)>>1
		if nums[mid] > target {
			right = mid - 1
		} else if nums[mid] < target {
			left = mid + 1
		} else {
			right = mid
		}
	}

	if nums[left] != target {
		return []int{-1, -1}
	}

	l := left
	right = n - 1
	for left < right {
		mid := right - (right-left)>>1
		if nums[mid] > target {
			right = mid - 1
		} else if nums[mid] < target {
			left = mid + 1
		} else {
			left = mid
		}
	}
	return []int{l, right}
}
