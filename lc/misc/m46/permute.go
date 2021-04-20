package m46

// 思路
//	leetcode前排的中级题目回溯很多
func permute(nums []int) [][]int {
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
	}

	for i, v := range nums {
		nums[i] = nums[0]
		(*mark)[depth] = v
		permute_recursion(nums[1:], depth+1, mark, result)
		nums[i] = v
	}
}
