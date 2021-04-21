package m47

import "sort"

// 思路
//	原本以为在m46基础上稍微改动即可，但是这个剪枝动作不是太容易
//	几点心得
//	1.全排列回溯时，不要修改原数组，会使问题变得复杂

func permuteUnique(nums []int) [][]int {
	n := len(nums)
	if n == 0 {
		return nil
	}
	sort.Ints(nums)

	perm := make([]int, 0, n)
	flag := make([]bool, n)
	result := make([][]int, 0)
	var permute_recursion func(int)
	permute_recursion = func(pos int) {
		if pos == n {
			result = append(result, append([]int{}, perm...))
			return
		}

		for i, v := range nums {
			// 这个操作相当于在当前pos代表的位置上，对于同一个数字，只会出现一次
			// 比如nums中有三个0的情况下，相当于另外两个0被剪枝减掉了
			if flag[i] || i > 0 && nums[i-1] == v && flag[i-1] {
				continue
			}

			// 以下逻辑也能通过，没思考明白，TODO
			// if flag[i] || i > 0 && nums[i-1] == v && !flag[i-1] {
			// 	continue
			// }

			perm = append(perm, nums[i])
			flag[i] = true
			permute_recursion(pos + 1)
			flag[i] = false
			perm = perm[:len(perm)-1]
		}
	}
	permute_recursion(0)

	return result
}
