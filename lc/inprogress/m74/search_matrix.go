package m74

func searchMatrix(matrix [][]int, target int) bool {
	m := len(matrix)
	n := len(matrix[0])
	right := m * n
	left := 0
	for left < right {
		mid := left + (right-left)>>1
		i := mid / n
		j := mid % n
		if matrix[i][j] > target {
			right = mid
		} else if matrix[i][j] < target {
			left = mid + 1
		} else {
			return true
		}
	}
	return false
}
