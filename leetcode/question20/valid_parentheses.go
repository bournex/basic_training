package question20

// 栈应用

func isValid(s string) bool {
	if len(s) == 0 {
		return true
	}
	if len(s)%2 != 0 {
		return false
	}

	stack := make([]rune, 0)
	for _, c := range s {
		if c == '{' || c == '[' || c == '(' {
			stack = append(stack, c)
		} else {

			var c1 rune
			switch c {
			case '}':
				c1 = '{'
			case ']':
				c1 = '['
			case ')':
				c1 = '('
			default:
				return false
			}
			if len(stack) == 0 || stack[len(stack)-1] != c1 {
				return false
			}
			stack = stack[:len(stack)-1]
		}
	}
	return len(stack) == 0
}
