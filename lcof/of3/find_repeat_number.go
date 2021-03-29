package of3

func findRepeatNumber(nums []int) int {
	// 原地置换虽然可以更快，但是注意修改了原数组
	for i, v := range nums {
		if v == -1 {
			continue
		}
		if nums[v] != -1 {
			if i == v {
				nums[i] = -1
			} else if nums[v] == v {
				return v
			} else {
				tmp := nums[v]
				nums[v] = -1
				nums[i] = tmp
			}
		} else {
			return v
		}
	}
	return -1
}

func findRepeatNumber1(nums []int) int {
	// 原地置换虽然可以更快，但是注意修改了原数组
	arr := make([]int, len(nums))
	for _, v := range nums {
		if arr[v] == 0 {
			arr[v] = 1
		} else {
			return v
		}
	}
	return -1
}
