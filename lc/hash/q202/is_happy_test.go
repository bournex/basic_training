package q202

import (
	"fmt"
	"testing"
)

func TestIsHappy(t *testing.T) {
	examples := []struct {
		Input  int
		Expect bool
	}{
		{
			19,
			true,
		},
		{
			2,
			false,
		},
	}

	for i, v := range examples {
		t.Run(fmt.Sprintf("case %d", i+1), func(t *testing.T) {
			res := isHappy(v.Input)
			if res != v.Expect {
				t.Errorf("input %d, expect %t, got %t", v.Input, v.Expect, res)
			}
		})
	}
}
