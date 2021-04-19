package m33

// 思路
//	部分排序的数组，除了二分查找的中间点，多了一个特殊判断点，就是被旋转的点
//	递归中，空间被中间点和旋转点划分为3个部分，针对三个部分分别递归
func search(nums []int, target int) int {

	return searching(nums, 0, len(nums)-1, target)
}

func searching(nums []int, left, right, target int) int {
	if left > right {
		return -1
	}
	mid := left + (right-left)>>1

	if target == nums[mid] {
		return mid
	} else if target == nums[right] {
		return right
	} else {
		if nums[mid] > nums[right] {
			if nums[right] < target && target < nums[mid] {
				return searching(nums, left, mid, target)
			} else {
				return searching(nums, mid+1, right, target)
			}
		} else if nums[mid] < nums[right] {
			if nums[mid] < target && target < nums[right] {
				return searching(nums, mid+1, right, target)
			} else {
				return searching(nums, left, mid-1, target)
			}
		}
	}
	return -1
}
