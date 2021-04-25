package m45

import (
	"fmt"
	"testing"
)

func TestJump(t *testing.T) {
	testCases := []struct {
		Input  []int
		Expect int
	}{
		{
			[]int{2, 3, 1, 1, 4},
			2,
		},
		{
			[]int{5, 6, 4, 4, 6, 9, 4, 4, 7, 4, 4, 8, 2, 6, 8, 1, 5, 9, 6, 5, 2, 7, 9, 7, 9, 6, 9, 4, 1, 6, 8, 8, 4, 4, 2, 0, 3, 8, 5},
			3,
		},
	}
	for i, v := range testCases {
		t.Run(fmt.Sprintf("case %d", i+1), func(t *testing.T) {
			res := jump(v.Input)
			if res != v.Expect {
				t.Errorf("input %+v, expect %d, got %d", v.Input, v.Expect, res)
			}
		})
	}
}
