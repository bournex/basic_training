package m54

func spiralOrder(matrix [][]int) []int {
	if len(matrix) == 0 {
		return []int{}
	}
	if len(matrix) == 1 {
		return matrix[0]
	}

	w, h := len(matrix[0]), len(matrix)

	result := make([]int, 0, w*h)

	left, top := 0, 0
	right, bottom := w-1, h-1
	for left <= right && top <= bottom {
		for i := 0; i < right-left+1; i++ {
			result = append(result, matrix[top][left+i])
		}
		for i := 1; i < bottom-top+1; i++ {
			result = append(result, matrix[top+i][right])
		}
		for i := right - left - 1; i >= 0 && bottom-top > 0; i-- {
			result = append(result, matrix[bottom][left+i])
		}
		for i := bottom - top - 1; i >= 1 && right-left > 0; i-- {
			result = append(result, matrix[top+i][left])
		}

		left += 1
		top += 1
		right -= 1
		bottom -= 1
	}
	return result
}
