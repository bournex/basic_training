package str

// Least-Significant-Digit First，低位优先字符串排序
func Lsd(ss []string) {
	// 思路
	// 假定前提，输入的字符串长度是相同的
	// 从后向前按ascii进行排序

	// 数组长度
	n := len(ss)
	if n == 0 {
		return
	}
	// 单个字符串长度
	m := len(ss[0])

	// 使用一个长度为256的整形数组，用于统计单列字符的分布情况，也保存分量排序后的起始位置
	// 使用一个长度为n的string数组，用于保存临时排序结果
	arr := make([]string, n)

	for i := m - 1; i >= 0; i-- {
		cnt := make([]int, 256+1)

		// cnt内部
		// ......[a][b][c][d]....
		// ...[0][1][2][1][3]....
		for j := 0; j < n; j++ {
			// 统计所有字符串在第i列字符在256个ascii上的分布情况
			cnt[ss[j][i]+1]++
		}
		// cnt内部
		// ......[a][b][c][d]....
		// ...[0][1][3][4][7]....
		for j := 1; j < len(cnt); j++ {
			// 计算出256个分量在新数组中的起始位置
			cnt[j] = cnt[j] + cnt[j-1]
		}
		// 使用cnt
		// ...[a][b][c][d]....
		// ...[0][1][3][4][7]....
		for j := 0; j < n; j++ {
			// 按照256分量的起始位置索引，将字符串从ss写入到arr，同时更新分量索引
			// 错位使用cnt
			arr[cnt[ss[j][i]]] = ss[j]
			cnt[ss[j][i]]++
		}
		for j := 0; j < n; j++ {
			// 回写至原结果
			ss[j] = arr[j]
		}
	}
}
