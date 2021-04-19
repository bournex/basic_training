package m34

import (
	"fmt"
	"testing"
)

func TestSearchRange(t *testing.T) {
	testCases := []struct {
		Input  []int
		Target int
		Expect []int
	}{
		{
			[]int{5, 7, 7, 8, 8, 10},
			8,
			[]int{3, 4},
		},
		{
			[]int{5, 7, 7, 8, 8, 10},
			5,
			[]int{0, 0},
		},
		{
			[]int{3},
			3,
			[]int{0, 0},
		},
		{
			[]int{4, 4, 4, 4, 4, 4},
			3,
			[]int{-1, -1},
		},
		{
			[]int{4, 4, 4, 4, 4, 4},
			4,
			[]int{0, 5},
		},
		{
			[]int{},
			0,
			[]int{-1, -1},
		},
	}
	for i, v := range testCases {
		t.Run(fmt.Sprintf("case %d", i+1), func(t *testing.T) {
			res := searchRange(v.Input, v.Target)
			if len(res) != len(v.Expect) {
				t.Errorf("input %+v, target %d, expect %+v, got %+v", v.Input, v.Target, v.Expect, res)
				return
			}

			for i, v1 := range res {
				if v1 != v.Expect[i] {
					t.Errorf("input %+v, target %d, expect %+v, got %+v", v.Input, v.Target, v.Expect, res)
					return
				}
			}
		})
	}
}
