package m59

// 1, 2, 3
// 8, 9, 4
// 7, 6, 5

func generateMatrix(n int) [][]int {
	result := make([][]int, n)
	for i := 0; i < n; i++ {
		result[i] = make([]int, n)
	}

	k := 1
	left, right := 0, n-1
	top, bottom := 0, n-1
	for left <= right {
		for i := left; i <= right; i++ {
			result[top][i] = k
			k++
		}
		for i := top + 1; i <= bottom; i++ {
			result[i][right] = k
			k++
		}
		for i := right - 1; i >= left; i-- {
			result[bottom][i] = k
			k++
		}
		for i := bottom - 1; i > top; i-- {
			result[i][left] = k
			k++
		}
		left++
		right--
		top++
		bottom--
	}
	return result
}
