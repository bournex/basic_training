package q169

// TODO
func majorityElement(nums []int) int {
	i, j := 0, 1
	for j < len(nums) {
		if nums[i] >= nums[j] {
			temp := nums[i]
			nums[i] = nums[j]
			nums[j] = nums[i+1]
			nums[i+1] = temp
			i++
		}
		j++
	}

	return nums[len(nums)>>1]
}
