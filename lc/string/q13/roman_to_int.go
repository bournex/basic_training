package q13

func romanToInt(s string) int {
	// I 1
	// V 5
	// X 10
	// L 50
	// C 100
	// D 500
	// M 1000
	var (
		res  int
		prev byte
	)
	for _, v := range []byte(s) {
		switch v {
		case 'I':
			res += 1
		case 'V':
			res += 5
			if prev == 'I' {
				res -= 2
			}
		case 'X':
			res += 10
			if prev == 'I' {
				res -= 2
			}
		case 'L':
			res += 50
			if prev == 'X' {
				res -= 20
			}
		case 'C':
			res += 100
			if prev == 'X' {
				res -= 20
			}
		case 'D':
			res += 500
			if prev == 'C' {
				res -= 200
			}
		case 'M':
			res += 1000
			if prev == 'C' {
				res -= 200
			}
		}
		prev = v
	}

	return res
}
