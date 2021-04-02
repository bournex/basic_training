package of18

func exchange(nums []int) []int {
	if len(nums) == 0 {
		return nums
	}
	i, j := 0, len(nums)-1
	for i != j {
		if nums[i]&int(1) == 1 {
			// 奇数
			i++
		} else if nums[j]&int(1) == 0 {
			j--
		} else {
			tmp := nums[i]
			nums[i] = nums[j]
			nums[j] = tmp
		}
	}
	return nums
}
