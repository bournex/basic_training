package h41

// 思路
//	最小的缺失的正整数，一定小于等于len(nums)+1，因为如果要使len(nums)个数组成的数组中缺失数最大，就必须使len(nums)全都大于等于1且小于等于len(nums)，只要不连续，就会出现空缺，则最小缺失数就会取得更小的值。
//	基于这条原则，只需要遍历两遍数组即可，第一遍从0到len(nums)-1，负数保持为负数，大于len(nums)的数置为-1，如果遇到大于等于1，小于等于len(nums)，则将其放置到其值-1对应的slot上
//	如果该位置也有大于等于1，小于等于len(nums)的数字，则循环做同样的事情，直到所有值遍历完
//	第二遍只要从头找第一个为-1的位子，其索引+1就是目标值

func firstMissingPositive(nums []int) int {
	n := len(nums)
	for i := 0; i < n; i++ {
		idx := nums[i]
		nums[i] = -1
		for idx > 0 && idx <= n && idx != nums[idx-1] {
			idx, nums[idx-1] = nums[idx-1], idx
		}
	}

	for i := 0; i < n; i++ {
		if nums[i] == -1 {
			return i + 1
		}
	}

	return n + 1
}
