package m75

// 思路
//	意识到是双指针了，但是对于细节思考不到位，本不是难题，还是参阅了答案
func sortColors(nums []int) {
	left, right := 0, len(nums)-1

	for i := 0; i <= right; i++ {
		for i <= right && nums[i] == 2 {
			nums[i], nums[right] = nums[right], nums[i]
			right--
		}
		if nums[i] == 0 {
			nums[i], nums[left] = nums[left], nums[i]
			left++
		}
	}
}
