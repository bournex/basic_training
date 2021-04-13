package q38

func countAndSay(n int) string {

	get_num_chars := func(m int) string {
		var s string
		char := []string{"0", "1", "2", "3", "4", "5", "6", "7", "8", "9"}
		for m != 0 {
			s = char[m%10] + s
			m /= 10
		}
		return s
	}

	s := "1"
	for i := 1; i < n; i++ {
		var s1 string
		j, k := 0, 0
		for j < len(s) && k < len(s) {
			if s[j] == s[k] {
				k++
			} else {
				s1 += get_num_chars(k - j)
				s1 += string(s[j])
				j = k
			}
		}
		s1 += get_num_chars(k - j)
		s1 += string(s[j])

		s = s1
	}

	return s
}
