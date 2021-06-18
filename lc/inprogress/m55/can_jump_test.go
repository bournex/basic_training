package m55

import (
	"fmt"
	"testing"
)

func TestCanJump(t *testing.T) {
	testCases := []struct {
		Input  []int
		Expect bool
	}{
		{
			[]int{0, 2, 3},
			false,
		},
		{
			[]int{1, 2},
			true,
		},

		{
			[]int{3},
			true,
		},
		{
			[]int{0, 0, 0, 0, 4},
			false,
		},
		{
			[]int{3, 2, 1, 0, 4},
			false,
		},
		{
			[]int{2, 3, 1, 1, 4},
			true,
		},
	}
	for i, v := range testCases {
		t.Run(fmt.Sprintf("%d", i+1), func(t *testing.T) {
			res := canJump(v.Input)
			if res != v.Expect {
				t.Errorf("input %+v, expect %t, got %t", v.Input, v.Expect, res)
			}
		})
	}
}
