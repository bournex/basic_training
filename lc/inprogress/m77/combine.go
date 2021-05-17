package m77

// 1,2,3
// n = 3,k = 2
// 1,2
// 1,3
// 2,3
func combine(n int, k int) [][]int {
	if n < k {
		return nil
	}

	result := make([][]int, 0)
	ans := make([]int, k)

	var combine func(idx, deep int)
	combine = func(idx, deep int) {
		if deep == k {
			result = append(result, append([]int{}, ans[0:k]...))
			return
		}

		for i := idx; i <= n; i++ {
			ans[deep] = i
			combine(i+1, deep+1)
		}
	}
	combine(1, 0)
	return result
}
