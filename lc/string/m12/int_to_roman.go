package m12

// 字符          数值
// I             1
// V             5
// X             10
// L             50
// C             100
// D             500
// M             1000

// 原则：同一个阿拉伯数字有多种罗马数字表达，尽量选择大的罗马数字
// 贪婪查表
func intToRoman(num int) string {
	arabic := []int{1000, 900, 500, 400, 100, 90, 50, 40, 10, 9, 5, 4, 1}
	roman := []string{"M", "CM", "D", "CD", "C", "XC", "L", "XL", "X", "IX", "V", "IV", "I"}

	var result string
	n := len(arabic)
	for num > 0 {
		for i := 0; i < n; i++ {
			if num >= arabic[i] {
				num -= arabic[i]
				result += roman[i]
				break
			}
		}
	}

	return result
}
