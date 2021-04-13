package q67

func addBinary(a string, b string) string {
	la := len(a)
	lb := len(b)
	if la == 0 {
		return b
	}
	if lb == 0 {
		return a
	}

	max := func(i, j int) int {
		if i > j {
			return i
		}
		return j
	}

	for i := max(la, lb) - 1; i >= 0; i-- {

	}
	return ""
}
