package of38

func permutation(s string) []string {
	return permutation1(0, []byte(s))
}

func permutation1(idx int, s []byte) []string {
	if len(s[idx:]) == 1 {
		return []string{string(s)}
	}

	res := make([]string, 0)
	for i := idx; i < len(s); i++ {
		if i > idx && s[i] == s[idx] {
			continue
		}
		tmp := s[i]
		[]byte(s)[i] = s[idx]
		[]byte(s)[idx] = tmp

		res = append(res, permutation1(idx+1, s)...)

		tmp = s[i]
		[]byte(s)[i] = s[idx]
		[]byte(s)[idx] = tmp
	}
	return res
}
