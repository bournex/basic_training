package m81

import (
	"fmt"
	"testing"
)

func TestSearch(t *testing.T) {
	testCases := []struct {
		Input  []int
		Target int
		Expect bool
	}{
		{
			[]int{1, 1, 1, 1, 1, 1, 1},
			0,
			false,
		},
		{
			[]int{2, 5, 6, 0, 0, 1, 2},
			5,
			true,
		},
		{
			[]int{1, 0, 1, 1, 1},
			0,
			true,
		},
		{
			[]int{1},
			1,
			true,
		},
		{
			[]int{},
			1,
			false,
		},
		{
			[]int{2, 5, 6, 0, 0, 1, 2},
			1,
			true,
		},
		{
			[]int{2, 5, 6, 0, 0, 1, 2},
			4,
			false,
		},

		{
			[]int{0, 0, 1, 2, 2, 6, 7, 8, 9},
			8,
			true,
		},
		{
			[]int{0, 0, 1, 2, 2, 6, 7, 8, 10},
			9,
			false,
		},
	}
	for i, v := range testCases {
		t.Run(fmt.Sprintf("case %d", i+1), func(t *testing.T) {
			res := search(v.Input, v.Target)
			if res != v.Expect {
				t.Errorf("input %+v, target %d, expect %t, got %t", v.Input, v.Target, v.Expect, res)
			}
		})
	}
}
