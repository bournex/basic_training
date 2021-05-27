package m93

// 12121212
// 1.2.1.21212	X
// 1.2.12.1212	X
// 1.2.121.212
// 1.21.2.1212	X
// 1.21.21.212
// 1.21.212.12
// 1.212.1.212
// 1.212.12.12
// 1.212.121.2
// 12.1.2.1212	X
// 12.1.21.212
// 12.1.212.12
// 12.12.1.212
// 12.12.12.12
// 12.12.121.2
// 12.121.2.12
// 12.121.21.2
// 121.2.1.212
// 121.2.12.12
// 121.2.121.2
// 121.21.2.12
// 121.21.21.2
// 121.212.1.2

// 思考
//	过程中还尝试使用三指针方法，不过发现不用递归的话，三指针完成内循环后恢复状态很复杂，所以还是考虑用回溯

// func restoreIpAddresses(s string) []string {
// 	n := len(s)
// 	if n < 4 || n > 12 {
// 		return nil
// 	}

// 	result := make([]string, 0)
// 	left, right, move := 1, n+1, 3

// 	m := n + 3
// 	temp := make([]byte, m)
// 	temp[left], temp[right], temp[move] = '.', '.', '.'
// 	temp[0], temp[2], temp[n+2] = s[0], s[1], s[n-1]

// 	for i := 2; i < n-1; i++ {
// 		temp[i+2] = s[i]
// 	}

// 	result = append(result, string(temp))

// 	for left < m-4 {
// 		for right > left+4 {
// 			for i := move; i < right-1; i++ {
// 				temp[i], temp[i+1] = temp[i+1], temp[i]
// 				if validIp(temp) {
// 					result = append(result, string(append([]byte{}, temp...)))
// 				}
// 			}
// 		}
// 	}

// 	return result
// }

func restoreIpAddresses(s string) []string {
	n := len(s)

	dotcnt := 0
	temp := make([]byte, n+3)
	result := make([]string, 0)
	var backtrace func(idx int) // 参数表示当前带读取的s索引位置
	backtrace = func(idx int) {
		if dotcnt == 3 {
			if idx+3 < n || idx == n { // 最后一段超过了255
				return
			}
			for i := idx; i < n; i++ {
				if i > idx && s[idx] == '0' {
					return
				}
				temp[i+dotcnt] = s[i]
			}
			if validIp(temp) {
				result = append(result, string(append([]byte{}, temp...)))
			}
			return
		}

		for i := idx; i < idx+3 && i < n; i++ {
			if i > idx && s[idx] == '0' {
				// 单个段0开头的，不会有1位以上的情况，比如.0.可以，.01.非法
				return
			}

			temp[i+dotcnt] = s[i]

			temp[i+dotcnt+1] = '.'
			dotcnt++
			backtrace(i + 1)
			dotcnt--
		}
	}
	backtrace(0)

	return result
}

func validIp(s []byte) bool {
	var temp int
	var dotcnt int
	for i := 0; i < len(s); i++ {
		if s[i] == '.' {
			dotcnt++
			if temp > 255 {
				return false
			}
			temp = 0
		} else {
			temp = temp*10 + int(s[i]-'0')
		}
	}
	return dotcnt == 3 && temp <= 255
}
