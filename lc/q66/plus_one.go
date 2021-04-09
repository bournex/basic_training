package q66

func plusOne(digits []int) []int {
	n := len(digits)
	carrier := 1
	for i := n - 1; i >= 0; i-- {
		if digits[i]+carrier == 10 {
			digits[i] = 0
		} else {
			digits[i] += carrier
			return digits
		}
	}
	if carrier == 1 {
		temp := make([]int, n+1)
		temp[0] = 1
		digits = temp
	}
	return digits
}
