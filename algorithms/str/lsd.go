package str

// Least-Significant-Digit-First，低位优先字符串排序
// 其基础是键索引计数法，对字符串的每固定index位置进行索引
// 从后向前逐渐完成字符串数组的排序
// ss - 输入字符串数组
// m  - ss中定长字符串的长度
// 排序过程示例，竖线右侧为排序好的byte位
// aba - ab|a - c|aa - |aaa
// aab - ca|a - a|aa - |aab
// abc - aa|a - a|ab - |aba
// caa - aa|b - a|ba - |abc
// aaa - ab|c - a|bc - |caa
func Lsd(ss []string, m int) {

	// 数组长度
	n := len(ss)
	if n == 0 {
		return
	}

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
