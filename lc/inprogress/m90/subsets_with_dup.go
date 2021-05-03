package m90

import "sort"

// 思路
//	难得一次过的回溯
//	在每一层做去重，每一层从i递进到n，递进过程中如果遇到重复值则跳过
func subsetsWithDup(nums []int) [][]int {
	n := len(nums)
	sort.Ints(nums)

	// 预分配，避免过程中append导致重复分配，size指定为1，自动添加{}集合
	result := make([][]int, 1, 1<<n)
	tmp := make([]int, n)
	pos := 0

	var backtrace func(int)
	backtrace = func(idx int) {
		for i := idx; i < n; i++ {
			if i > idx && nums[i] == nums[i-1] {
				// 跳过当前层中的重复元素
				continue
			}
			tmp[pos] = nums[i]
			pos++
			result = append(result, append([]int{}, tmp[:pos]...))
			backtrace(i + 1)
			pos--
		}
	}
	backtrace(0)

	return result
}
