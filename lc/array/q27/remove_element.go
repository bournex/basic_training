package q27

func removeElement(nums []int, val int) int {
	k := 0
	for i, v := range nums {
		if v == val {
			continue
		}
		if k != i {
			nums[k] = v
		}
		k++
	}

	return k
}
