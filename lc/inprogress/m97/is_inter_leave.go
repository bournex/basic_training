package m97

// 思路
//	贪婪匹配，读取s3，轮流匹配s1、s2，如果出现与s1不匹配的字符，则开始匹配s2，如果两个都无法匹配，则回退
//	贪婪不行，回退太多
//	官方实现动态规划
//

// 滚动一维数组版本
func isInterleave(s1 string, s2 string, s3 string) bool {
	n1, n2, n3 := len(s1), len(s2), len(s3)
	if n1+n2 != n3 {
		return false
	}

	ans := make([]bool, n2+1)

	ans[0] = true
	for i := 0; i <= n1; i++ {
		for j := 0; j <= n2; j++ {
			if i > 0 {
				ans[j] = ans[j] && s1[i-1] == s3[i+j-1]
			}
			if j > 0 {
				ans[j] = ans[j] || ans[j-1] && s2[j-1] == s3[i+j-1]
			}
		}
	}
	return ans[n2]
}

// 二维数组版本
func isInterleave1(s1 string, s2 string, s3 string) bool {
	n1, n2, n3 := len(s1), len(s2), len(s3)
	if n1+n2 != n3 {
		return false
	}

	ans := make([][]bool, n1+1)
	for i := 0; i <= n1; i++ {
		ans[i] = make([]bool, n2+1)
	}

	ans[0][0] = true
	for i := 0; i <= n1; i++ {
		for j := 0; j <= n2; j++ {
			if i != 0 && j != 0 {
				ans[i][j] = (ans[i-1][j] && s1[i-1] == s3[i+j-1]) || (ans[i][j-1] && s2[j-1] == s3[i+j-1])
			} else {
				if i == 0 && j != 0 {
					ans[i][j] = ans[i][j-1] && s2[j-1] == s3[i+j-1]
				} else if i != 0 && j == 0 {
					ans[i][j] = ans[i-1][j] && s1[i-1] == s3[i+j-1]
				}
			}
		}
	}
	return ans[n1][n2]
}
