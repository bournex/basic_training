package m73

// 思路
//	思考到陷入怀疑，想到了用矩阵本身作为标记位，但是没有想到用额外的两个变量保存首行、首列的内容
//
func setZeroes(matrix [][]int) {
	var left, top bool
	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[0]); j++ {
			if matrix[i][j] == 0 {
				if i > 0 && j > 0 {
					matrix[i][0] = 0
					matrix[0][j] = 0
				} else {
					if i == 0 {
						top = true
					}
					if j == 0 {
						left = true
					}
				}
			}
		}
	}

	for i := 1; i < len(matrix[0]); i++ {
		if matrix[0][i] == 0 {
			for j := 1; j < len(matrix); j++ {
				matrix[j][i] = 0
			}
		} else if top {
			matrix[0][i] = 0
		}
	}

	for i := 1; i < len(matrix); i++ {
		if matrix[i][0] == 0 {
			for j := 1; j < len(matrix[0]); j++ {
				matrix[i][j] = 0
			}
		} else if left {
			matrix[i][0] = 0
		}
	}

	if top || left {
		matrix[0][0] = 0
	}
}
