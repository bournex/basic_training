package h30

// 思考
//  方案1：采用滑动窗口 和 字符映射表，方案超时
//	方案2：构建字符串查找树
//
func findSubstring(s string, words []string) []int {
	n := len(words)
	if n <= 0 {
		return nil
	}

	var length int
	var table [26]bool
	hash := make(map[string]int, n)
	cnt := make([]int, 0, n)
	for i := 0; i < n; i++ {
		length += len(words[i])
		if length > len(s) {
			return nil
		}

		for j := 0; j < len(words[i]); j++ {
			table[words[i][j]-'a'] = true
		}

		if _, ok := hash[words[i]]; !ok {
			cnt = append(cnt, 1)
			hash[words[i]] = len(cnt) - 1
		} else {
			cnt[hash[words[i]]]++
		}
	}

	var total int
	temp := make([]int, len(cnt))
	head, tail := 0, length-1

	result := make([]int, 0)

	for tail < len(s) {
		copy(temp, cnt)
		i := head
		for j := i; j <= tail; j++ {
			if !table[s[j]-'a'] {
				head = i
				break
			}
			if v, ok := hash[s[i:j+1]]; ok {
				if temp[v] == 0 {
					// a没有匹配上，不代表aa匹配不上
					continue
				}
				temp[v]--
				i = j + 1
				total++
			}
		}
		if total == n {
			result = append(result, head)
		}
		head++
		tail = head + length - 1
		total = 0
	}

	return result
}
