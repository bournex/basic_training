package m17

import "unsafe"

func letterCombinations(digits string) []string {
	if len(digits) == 0 {
		return []string{}
	}
	table := [][]byte{
		{'a', 'b', 'c'}, {'d', 'e', 'f'}, {'g', 'h', 'i'},
		{'j', 'k', 'l'}, {'m', 'n', 'o'}, {'p', 'q', 'r', 's'},
		{'t', 'u', 'v'}, {'w', 'x', 'y', 'z'},
	}
	return letter_combinations(digits, table)
}

func letter_combinations(digits string, table [][]byte) []string {
	n := digits[0]
	line := table[n-'2']

	words := make([]string, 0)
	for _, v := range line {
		if len(digits) == 1 {
			words = append(words, string(v))
		} else {
			for _, v1 := range letter_combinations(digits[1:], table) {
				words = append(words, string(v)+v1)
			}
		}

	}
	return words
}

// 空间改进版
func letterCombinations1(digits string) []string {
	if len(digits) == 0 {
		return []string{}
	}
	table := [][]byte{
		{'a', 'b', 'c'}, {'d', 'e', 'f'}, {'g', 'h', 'i'},
		{'j', 'k', 'l'}, {'m', 'n', 'o'}, {'p', 'q', 'r', 's'},
		{'t', 'u', 'v'}, {'w', 'x', 'y', 'z'},
	}
	nums := []int{3, 3, 3, 3, 3, 4, 3, 4}
	l := len(digits) // 返回字符串平均长度
	n := 1           // 返回的字符串数量
	for _, v := range digits {
		n *= nums[v-'2']
	}

	words := make([][]byte, n)
	for i := range words {
		words[i] = make([]byte, l)
	}

	letter_combinations1(digits, table, nums, words, 0, 0, n)

	result := make([]string, n)
	for i, v := range words {
		result[i] = *(*string)(unsafe.Pointer(&v))
	}
	return result
}

func letter_combinations1(digits string, table [][]byte, nums []int, words [][]byte, col, start, count int) {
	n := digits[0]
	line := table[n-'2']
	stride := count / nums[n-'2']

	for i, v := range line {
		for j := 0; j < stride; j++ {
			words[start+i*stride+j][col] = v
		}
		if len(digits) > 1 {
			letter_combinations1(digits[1:], table, nums, words, col+1, start+i*stride, stride)
		}
	}
}
