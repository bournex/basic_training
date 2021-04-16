package m16

import (
	"fmt"
	"testing"
)

func Test(t *testing.T) {
	testCases := []struct {
		Input  []int
		Target int
		Expect int
	}{
		{
			[]int{4, 0, 5, -5, 3, 3, 0, -4, -5},
			-2,
			-2,
		},
		{
			[]int{-1, 2, 1, -4},
			1,
			2,
		},
		{
			[]int{1, 1, 2, 3},
			7,
			6,
		},
		{
			[]int{1, 1, 1},
			100000000,
			3,
		},
	}
	for i, v := range testCases {
		t.Run(fmt.Sprintf("case %d", i+1), func(t *testing.T) {
			res := threeSumClosest(v.Input, v.Target)
			if res != v.Expect {
				t.Errorf("input %+v, target %d, expect %d, got %d", v.Input, v.Target, v.Expect, res)
			}
		})
	}
}
