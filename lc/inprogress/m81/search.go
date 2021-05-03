package m81

// 思考
//	官方答案是O(n)，也是二分的思路
//	不过O(n)的复杂度确实不如用遍历了，所以关键问题退化成，能否认识到在有重复元素的旋转非降序排序数组的情况下
//	二分是否会退化成线性查找的问题。
//	我这里用了递归，官方的更优一些，当边界值相等时，缩小边界
func search(nums []int, target int) bool {
	left, right := 0, len(nums)-1

	for left <= right {
		mid := left + (right-left)>>1

		if target == nums[mid] || target == nums[right] {
			return true
		}
		if nums[mid] > nums[right] {
			if target < nums[mid] && target > nums[right] {
				right = mid - 1
			} else {
				left = mid + 1
			}
		} else if nums[mid] < nums[right] {
			if target > nums[mid] && target < nums[right] {
				left = mid + 1
			} else {
				right = mid - 1
			}
		} else {
			if nums[mid] == target {
				return true
			}
			return search(nums[left:mid], target) || search(nums[mid:right], target)
		}
	}
	return false
}
