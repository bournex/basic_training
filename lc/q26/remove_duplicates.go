package q26

// 思路
//	用两个索引完成
//	第一个索引用于遍历整个数组
//	第二个索引用于标志下一个待写入位置
//	当第一个索引遍历到一个没有写入过的数字时，写入第二个索引的位置，并递增第二个索引
func removeDuplicates(nums []int) int {
	l := len(nums)
	if l <= 1 {
		return l
	}

	k := 0
	for _, v := range nums {
		if nums[k] != v {
			nums[k+1] = v
			k++
		}
	}

	return k + 1
}
