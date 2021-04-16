package m22

import "math"

// 思路
//	回溯法
//	与一般回溯法不同的是，当当前左右括号数量相等时，下一次只能取左括号

func generateParenthesis(n int) []string {
	pairs := make([]byte, n<<1)
	result := make([]string, 0, int(math.Pow(2, float64(n))))
	generate_parenthesis(pairs, 0, n, n, &result)
	return result[:]
}

func generate_parenthesis(pairs []byte, index, leftN, rightN int, result *[]string) {
	if leftN == 0 {
		for i := 0; i < rightN; i++ {
			pairs[i+index] = ')'
		}
		*result = append(*result, string(pairs))
		return
	}

	if leftN == rightN {
		pairs[index] = '('
		leftN--
		index++
	}

	if leftN > 0 {
		pairs[index] = '('
		generate_parenthesis(pairs, index+1, leftN-1, rightN, result)
	}

	if rightN > 0 {
		pairs[index] = ')'
		generate_parenthesis(pairs, index+1, leftN, rightN-1, result)
	}
}
