package m47

import "sort"

// 思路
//	在m46基础上稍微改动即可

func permuteUnique(nums []int) [][]int {
	sort.Ints(nums)
	if len(nums) == 0 {
		return nil
	}

	result := make([][]int, 0)
	mark := make([]int, len(nums))
	permute_recursion(nums, 0, &mark, &result)
	return result
}

func permute_recursion(nums []int, depth int, mark *[]int, result *[][]int) {
	if len(nums) == 0 {
		array := make([]int, len(*mark))
		copy(array, *mark)
		*result = append(*result, array)
		return
	}

	last := 0
	for i, v := range nums {
		if i == 0 || v != last {
			last = v
		} else {
			continue
		}

		nums[i] = nums[0]
		(*mark)[depth] = v
		permute_recursion(nums[1:], depth+1, mark, result)
		nums[i] = v
	}
}
