package question20

import "testing"

func TestIsValid(t *testing.T) {
	input := []string{"()", "()[]{}", "(]", "([)]", "{[]}", "((", "){"}
	expect := []bool{true, true, false, false, true, false, false}

	for i, v := range input {
		res := isValid(v)
		if expect[i] != res {
			t.Errorf("input %s, expect %t, got %t", v, expect[i], res)
		}
	}
}
