package of5

// 主要关注点
// 1.空数组或仅有空格的数组
// 2.避免子串的重复拷贝
func replaceSpace(s string) string {
	ret := make([]byte, 0, len(s)<<1+len(s))
	i := 0
	j := 0
	for ; j < len(s); j++ {
		if s[j] == ' ' {
			ret = append(ret, []byte(s[i:j])...)
			ret = append(ret, []byte("%20")...)
			i = j + 1
		}
	}

	if i != j {
		ret = append(ret, []byte(s[i:j])...)
	}

	return string(ret[0:])
}
