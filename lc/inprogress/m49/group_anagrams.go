package m49

func groupAnagrams(strs []string) [][]string {
	table := make(map[[26]int][]string)

	for _, v := range strs {
		var chs [26]int
		for _, v1 := range v {
			chs[v1-'a']++
		}
		if _, ok := table[chs]; ok {
			table[chs] = append(table[chs], v)
		} else {
			table[chs] = []string{v}
		}
	}

	result := make([][]string, 0, len(table))
	for _, v := range table {
		result = append(result, v)
	}
	return result
}
