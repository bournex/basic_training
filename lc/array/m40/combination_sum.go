package m40

import "sort"

func combinationSum2(candidates []int, target int) [][]int {
	sort.Ints(candidates)

	n := len(candidates)
	// perm := 0               // 累加值
	array := make([]int, n) // 保存中间结果
	flag := make([]bool, n)

	result := make([][]int, 0)
	var proc func(int, int, int)
	proc = func(pos int, perm int, index int) {
		// pos是candidates的索引
		if perm == target {
			tmp := make([]int, index)
			copy(tmp, array[:index])
			result = append(result, tmp)
			return
		} else if perm > target {
			return
		}

		for i := pos; i < n; i++ {
			if flag[i] || i > 0 && candidates[i] == candidates[i-1] && !flag[i] {
				continue
			}

			t := candidates[i]
			candidates[i] = candidates[pos]
			candidates[pos] = t

			flag[i] = true
			array[index] = t
			proc(pos+1, perm+t, index+1)
			flag[i] = false

			// 还原candidates
			t = candidates[i]
			candidates[i] = candidates[pos]
			candidates[pos] = t
		}
	}
	proc(0, 0, 0)

	return result
}
