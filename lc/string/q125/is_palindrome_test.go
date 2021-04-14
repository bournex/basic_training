package q125

import (
	"fmt"
	"testing"
)

func TestIsPalindrome(t *testing.T) {
	examples := []struct {
		Input  string
		Expect bool
	}{
		{
			"0P",
			false,
		},
		{
			"A man, a plan, a canal: Panama",
			true,
		},
		{
			"race a car",
			false,
		},
		{
			"a",
			true,
		},
		{
			"",
			true,
		},
	}

	for i, v := range examples {
		t.Run(fmt.Sprintf("case %d", i+1), func(t *testing.T) {
			res := isPalindrome(v.Input)
			if res != v.Expect {
				t.Errorf("input %s, expect %t, got %t", v.Input, v.Expect, res)
			}
		})
	}
}
