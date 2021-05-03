package m78

// 思路
//	朴素的回溯算法
func subsets(nums []int) [][]int {
	n := len(nums)
	if n == 0 {
		return [][]int{{}}
	} else if n == 1 {
		return [][]int{{}, nums}
	}
	var pos int
	arr := make([]int, n)
	result := make([][]int, 1, 1<<n)
	var backtrace func(int)
	backtrace = func(idx int) {
		if idx == n {
			return
		}

		for i := idx; i < n; i++ {
			arr[pos] = nums[i]
			result = append(result, append([]int{}, arr[:pos+1]...))
			pos += 1
			backtrace(i + 1)
			pos -= 1
		}
	}
	backtrace(0)
	return result
}
