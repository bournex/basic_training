package of38

func permutation(s string) []string {
	return permutation1(0, []byte(s))
}

func permutation1(idx int, s []byte) []string {
	if len(s[idx:]) == 1 {
		return []string{string(s)}
	}

	res := make([]string, 0)
	set := make(map[byte]bool)
	for i := idx; i < len(s); i++ {
		if _, ok := set[s[i]]; ok {
			continue
		}

		set[s[i]] = true
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
