package m43

// 思考
//	基本乘法运算，细节较多，需要细心，主要是处理好进位、ascii与整数的互转

func multiply(num1 string, num2 string) string {
	if num1 == "0" || num2 == "0" {
		return "0"
	}
	if num1 == "1" {
		return num2
	} else if num2 == "1" {
		return num1
	}

	n1, n2 := len(num1), len(num2)
	result := make([]byte, n1+n2)
	pos := n1 + n2 - 1

	var c1, c2, carrier byte
	for i := n1 - 1; i >= 0; i-- {
		c1 = num1[i] - '0'
		position := pos

		for j := n2 - 1; j >= 0; j-- {
			c2 = num2[j] - '0'

			result[position] = c1*c2 + carrier + result[position]
			carrier = result[position] / 10
			result[position] = result[position] % 10
			position--
		}
		result[position] += carrier
		carrier = 0
		pos--
	}

	for i := range result {
		result[i] += '0'
	}

	if result[0] == '0' {
		return string(result[1:])
	}
	return string(result)
}
