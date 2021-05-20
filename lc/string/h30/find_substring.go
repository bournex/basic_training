package h30

// 思考
//  方案1：采用滑动窗口 和 字符映射表，方案超时
//	方案2：构建字符串查找树
//	2021-05-20，吃了审题不细致的亏，words中的字符串长度相同，导致了方案1超时

func findSubstring(s string, words []string) []int {
	// 思路整理
	// 将words中的字符串放入hash表，hash表的key为words中的字符串
	// 由于words中可能有重复的字符串，所以再创建一个数组，用于保存字符串出现的次数，map的value为数组索引
	// 为什么要再创建一个数组？因为滑动窗口过程需要 修改计数，确保窗口中重复串的数量与words中的数量一致
	// 滑动窗口固定，子串长度固定

	n := len(words)
	if n <= 0 {
		return nil
	}

	singleLength := len(words[0])
	totalLength := singleLength * n

	hash := make(map[string]int, n)
	cnt := make([]int, 0, n)
	for i := 0; i < n; i++ {
		// 构建hash表和cnt数组
		if _, ok := hash[words[i]]; !ok {
			cnt = append(cnt, 1)
			hash[words[i]] = len(cnt) - 1
		} else {
			cnt[hash[words[i]]]++
		}
	}

	var total int
	temp := make([]int, len(cnt))
	head, tail := 0, totalLength

	result := make([]int, 0)
	for tail <= len(s) {
		copy(temp, cnt)

		i := head
		j := i + singleLength // i,j为子串起止索引，左闭右开
		for j <= tail {
			if v, ok := hash[s[i:j]]; ok {
				if temp[v] == 0 {
					// 在建就小于0了，说明s[i:j]这个字符串超过了限制，不匹配
					break
				}
				temp[v]--

				i = j
				j = i + singleLength
				total++
			} else {
				// 子串不在hash中，退出内循环
				break
			}
		}

		if total == n {
			// 如果总数能对的上，说明匹配成功
			result = append(result, head)
		}
		head++
		tail = head + totalLength
		total = 0
	}

	return result
}

func findSubstring1(s string, words []string) []int {
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

	for tail <= len(s) {
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
