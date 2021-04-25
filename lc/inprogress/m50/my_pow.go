package m50

// 思考
//	主要是细节考察
//	另外浮点数不能移位，需要渐进移位
func myPow(x float64, n int) float64 {

	var (
		isFraction = n < 0
		// isNegative = (x < 0 && n%2 == 1)
	)

	if isFraction {
		n = -n
	}

	m := 1
	y := x
	sum := 1.0
	for n > 0 {
		if m<<1 > n {
			n -= m
			sum *= y
			y = x
			m = 1
		} else {
			y *= y
			m <<= 1
		}
	}

	// if isNegative {
	// 	sum *= -1
	// }
	if isFraction {
		sum = 1 / sum
	}
	return sum
}
