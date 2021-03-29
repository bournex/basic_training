package question6

func convert(s string, numRows int) string {
	if numRows == 1 {
		return s
	}
	s1 := make([]byte, len(s))
	var idx int
	var step int = (numRows - 1) * 2

	for i := 0; i < numRows; i++ {
		l, r := i, step-i
		if l == 0 || l == r {
			for j := 0; j*step < len(s); j++ {
				if j*step+l >= len(s) {
					break
				} else {
					s1[idx] = s[j*step+l]
					idx++
				}
			}
		} else if l != r {
			for j := 0; j*step < len(s); j++ {
				if j*step+l >= len(s) {
					break
				} else {
					s1[idx] = s[j*step+l]
					idx++
				}

				if j*step+r >= len(s) {
					break
				} else {
					s1[idx] = s[j*step+r]
					idx++
				}
			}
		}
	}

	return string(s1)
}
