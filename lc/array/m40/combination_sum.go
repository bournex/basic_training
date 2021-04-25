package m40

import "sort"

func combinationSum2(candidates []int, target int) [][]int {
	result := make([][]int, 0)
	path := make([]int, 0)

	sort.Ints(candidates)
	combination_sum(candidates, 0, target, path, &result)
	return result
}

func combination_sum(candidates []int, current int, target int, path []int, result *[][]int) {
	for i, v := range candidates {
		if current+v > target {
			continue
		} else if current+v < target {
			path = append(path, v)
			combination_sum(candidates[i:], current+v, target, path, result)
			path = path[0 : len(path)-1]
		} else {
			s := make([]int, len(path)+1)
			copy(s, path)
			s[len(s)-1] = v
			*result = append(*result, s)
		}
	}
}
