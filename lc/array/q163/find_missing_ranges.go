package q163

import "fmt"

//  TODO
func findMissingRanges(nums []int, lower int, upper int) []string {
	if len(nums) == 0 {
		if lower == upper {
			return []string{fmt.Sprintf("%d", lower)}
		} else {
			return []string{fmt.Sprintf("%d->%d", lower, upper)}
		}
	}
	result := make([]string, 0)
	for i := 1; i < len(nums); i++ {
		n := nums[i] - nums[i-1]

		if n == 2 {
			result = append(result, fmt.Sprintf("%d", nums[i-1]+1))
		} else if n > 2 {
			result = append(result, fmt.Sprintf("%d->%d", nums[i-1]+1, nums[i]-1))
		}
	}

	start, end := "", ""
	if nums[0]-lower == 1 {
		start = fmt.Sprintf("%d", lower)
	} else if nums[0]-lower > 1 {
		start = fmt.Sprintf("%d->%d", lower, nums[0]-1)
	}

	if upper-nums[len(nums)-1] == 1 {
		end = fmt.Sprintf("%d", upper)
	} else if upper-nums[len(nums)-1] > 1 {
		end = fmt.Sprintf("%d->%d", nums[len(nums)-1]+1, upper)
	}

	if start != "" {
		result = append([]string{start}, result[:]...)
	}
	if end != "" {
		result = append(result, end)
	}

	// 处理结尾
	return result
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
