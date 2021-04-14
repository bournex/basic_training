package q136

func singleNumber(nums []int) int {
	var n int
	for _, v := range nums {
		n ^= v
	}
	return n
}
