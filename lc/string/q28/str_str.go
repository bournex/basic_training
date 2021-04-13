package q28

func strStr(haystack string, needle string) int {
	n1 := len(haystack)
	n2 := len(needle)
	if n1 < n2 {
		return -1
	}
	if n2 == 0 {
		return 0
	}

	for i := 0; i < (n1 - n2 + 1); i++ {
		j := 0
		for ; j < n2; j++ {
			if haystack[i+j] != needle[j] {
				break
			}
		}
		if j == n2 {
			return i
		}
	}

	return -1
}
