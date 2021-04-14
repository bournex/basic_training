package q119

// 思路
//	O(n)空间复杂度，需要一开始就申请好返回需要的全部空间，循环使用
//	为了避免每次覆盖前次的结果，每层数据存储在结果slice的尾部
func getRow(rowIndex int) []int {
	if rowIndex == 0 {
		return []int{1}
	}
	if rowIndex == 1 {
		return []int{1, 1}
	}

	level := make([]int, rowIndex+1)
	level[rowIndex], level[rowIndex-1] = 1, 1
	for i := 2; i < rowIndex+1; i++ {
		level[rowIndex-i] = 1
		for j := rowIndex - i + 1; j < rowIndex; j++ {
			level[j] = level[j] + level[j+1]
		}
	}

	return level
}
