package m48

func rotate(matrix [][]int) {
	if len(matrix) == 0 {
		return
	}
	lt, rt := 0, len(matrix[0])-1
	for lt < rt {
		rotate_ring(matrix, lt, rt)
		lt++
		rt--
	}
}

func rotate_ring(matrix [][]int, lt, rt int) {
	top := lt
	left := lt
	right := rt
	bottom := rt

	for i := 0; i < rt-lt; i++ {
		// 上右调换
		t := matrix[top][left+i]
		matrix[top][left+i] = matrix[top+i][right]
		matrix[top+i][right] = t
		// 上下调换
		t = matrix[top][left+i]
		matrix[top][left+i] = matrix[bottom][right-i]
		matrix[bottom][right-i] = t
		// 上左调换
		t = matrix[top][left+i]
		matrix[top][left+i] = matrix[bottom-i][left]
		matrix[bottom-i][left] = t
	}
}
