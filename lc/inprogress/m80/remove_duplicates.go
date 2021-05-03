package m80

func removeDuplicates(nums []int) int {

	idx := 1
	num := nums[0]
	cnt := 1

	for i := 1; i < len(nums); i++ {
		if nums[i] == num && cnt < 2 {
			nums[idx] = nums[i]
			idx++
			cnt++
		} else if nums[i] != num {
			num = nums[i]
			nums[idx] = num
			cnt = 1
			idx++
		}
	}
	return idx
}

// 官方双指针解法
func removeDuplicates1(nums []int) int {
	n := len(nums)
	if n <= 2 {
		return n
	}

	slow, fast := 2, 2
	for fast < n {
		if nums[slow-2] != nums[fast] {
			nums[slow] = nums[fast]
			slow++
		}
		fast++
	}
	return slow
}
