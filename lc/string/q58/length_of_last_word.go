package q58

func lengthOfLastWord(s string) int {
	if len(s) == 0 {
		return 0
	}
	e := -1
	i := len(s) - 1
	for ; i >= 0; i-- {
		if s[i] >= 'a' && s[i] <= 'z' || s[i] >= 'A' && s[i] <= 'Z' {
			if e == -1 {
				e = i
			}
		} else if e != -1 {
			return e - i
		}
	}

	if e != -1 {
		return e - i
	}
	return 0
}
