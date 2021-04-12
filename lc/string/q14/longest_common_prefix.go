package q14

// 思路
//	纵向按位比较，主要是考察细节处理
func longestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}
	if len(strs) == 1 {
		return strs[0]
	}

	min := 201
	for _, v := range strs {
		if len(v) < min {
			min = len(v)
		}
	}

	var arr string
	for i := 0; i < min; i++ {
		b := strs[0][i]
		for j := 1; j < len(strs); j++ {
			if strs[j][i] != b {
				return arr
			}
		}
		arr += string(b)
	}

	return arr
}
