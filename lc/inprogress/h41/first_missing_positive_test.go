package h41

import (
	"fmt"
	"testing"
)

func TestFirstMissingPositive(t *testing.T) {
	testCases := []struct {
		Input  []int
		Expect int
	}{
		{
			[]int{1, 2, 0},
			3,
		},
		{
			[]int{3, 7, 6, 4, 1, 5, 0, 9, 2},
			8,
		},
		{
			[]int{1, 2, 3},
			4,
		},
		{
			[]int{100, 101, 102},
			1,
		},
		{
			[]int{3, 4, -1, 1},
			2,
		},
		{
			[]int{7, 8, 9, 11, 12},
			1,
		},
	}
	for i, v := range testCases {
		t.Run(fmt.Sprintf("case %d", i+1), func(t *testing.T) {
			res := firstMissingPositive(v.Input)
			if res != v.Expect {
				t.Errorf("input %+v, expect %d, got %d", v.Input, v.Expect, res)
				return
			}
		})
	}
}
