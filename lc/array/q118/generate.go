package q118

func generate(numRows int) [][]int {
	triangle := [][]int{[]int{1}, []int{1, 1}}
	if numRows < 3 {
		return triangle[0:numRows]
	}

	for i := 2; i < numRows; i++ {
		up := triangle[i-1]
		down := make([]int, len(up)+1)
		for j := 1; j < len(down)-1; j++ {
			down[j] = up[j-1] + up[j]
		}
		down[0], down[len(down)-1] = 1, 1
		triangle = append(triangle, down)
	}

	return triangle[0:numRows]
}
