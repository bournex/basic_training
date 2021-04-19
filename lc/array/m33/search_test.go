package m33

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
			[]int{3, 5, 1},
			3,
			0,
		},
		{
			[]int{1, 3},
			1,
			0,
		},
		{
			[]int{1, 3},
			3,
			1,
		},
		{
			[]int{4, 5, 6, 7, 0, 1, 2},
			0,
			4,
		},
		{
			[]int{4, 5, 6, 7, 0, 1, 2},
			3,
			-1,
		},
		{
			[]int{1},
			1,
			0,
		},
		{
			[]int{1},
			0,
			-1,
		},
	}
	for i, v := range testCases {
		t.Run(fmt.Sprintf("case %d", i+1), func(t *testing.T) {
			res := search(v.Input, v.Target)
			if res != v.Expect {
				t.Errorf("input %+v, target %d, expect %d, got %d", v.Input, v.Target, v.Expect, res)
			}
		})
	}
}
