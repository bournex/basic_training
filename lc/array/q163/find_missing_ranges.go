package q163

//  TODO
func findMissingRanges(nums []int, lower int, upper int) []string {

	// 处理结尾
	return nil
}

// 思考
//	审题不细致
//	理解错了，还在尝试找左侧和右侧最接近节点索引
func findLeftClose(nums []int, left, right, n int) int {
	for left < right {
		mid := left + (right-left)>>1
		if nums[mid] < n {
			left = mid + 1
		} else if nums[mid] > n {
			right = mid
		} else {
			return mid
		}
	}

	return right
}

func findRightClose(nums []int, left, right, n int) int {
	for left < right {
		mid := right - (right-left)>>1
		if nums[mid] < n {
			left = mid
		} else if nums[mid] > n {
			right = mid - 1
		} else {
			return mid
		}
	}
	return left
}
