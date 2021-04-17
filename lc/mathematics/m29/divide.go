package m29

// 思考
//	主要考察移位操作，和有符号数的边界值处理
func divide(dividend int, divisor int) int {
	sign := 1
	if dividend < 0 {
		sign = -sign
		dividend = -dividend
		if uint32(dividend) == (^uint32(0)>>1)+1 && divisor == -1 {
			return int(^uint32(0) >> 1)
		}
	}
	if divisor < 0 {
		sign = -sign
		divisor = -divisor
	}

	if dividend < divisor {
		return 0
	} else if dividend == divisor {
		return sign
	}

	i := 1
	quotient := 0
	d := divisor

	for dividend >= divisor {
		if d > dividend>>1 {
			quotient += i
			dividend -= d
			i = 1
			d = divisor
		} else if d < dividend {
			i <<= 1
			d <<= 1
		} else {
			break
		}
	}
	if sign < 0 {
		return -quotient
	}
	return quotient
}
