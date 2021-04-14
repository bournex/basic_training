package q202

func isHappy(n int) bool {
	history := make(map[int]bool)
	result := 0
	for result != 1 {
		result = 0
		for n > 0 {
			t := n % 10
			result += t * t
			n /= 10
		}

		if _, ok := history[result]; !ok {
			history[result] = true
		} else {
			return false
		}
		n = result
	}
	return true
}
