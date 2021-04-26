package m54

import (
	"fmt"
	"testing"
)

func TestSpiralOrder(t *testing.T) {
	testCases := []struct {
		Input  [][]int
		Expect []int
	}{
		{
			[][]int{},
			[]int{},
		},
		{
			[][]int{{1}},
			[]int{1},
		},
		{
			[][]int{{1, 2, 3}},
			[]int{1, 2, 3},
		},
		{
			[][]int{{1}, {2}, {3}},
			[]int{1, 2, 3},
		},
		{
			[][]int{{1, 2}, {3, 4}},
			[]int{1, 2, 4, 3},
		},
		{
			[][]int{{1, 2, 3}, {4, 5, 6}},
			[]int{1, 2, 3, 6, 5, 4},
		},
		{
			[][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}},
			[]int{1, 2, 3, 6, 9, 8, 7, 4, 5},
		},
		{
			[][]int{{1, 2, 3, 4}, {5, 6, 7, 8}, {9, 10, 11, 12}},
			[]int{1, 2, 3, 4, 8, 12, 11, 10, 9, 5, 6, 7},
		},
		{
			[][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}, {10, 11, 12}},
			[]int{1, 2, 3, 6, 9, 12, 11, 10, 7, 4, 5, 8},
		},
	}
	for i, v := range testCases {
		t.Run(fmt.Sprintf("case %d", i+1), func(t *testing.T) {
			res := spiralOrder(v.Input)
			if len(res) != len(v.Expect) {
				t.Errorf("input %+v, expect %+v, got %+v", v.Input, v.Expect, res)
				return
			}

			for j, v1 := range res {
				if v1 != v.Expect[j] {
					t.Errorf("input %+v, expect %+v, got %+v", v.Input, v.Expect, res)
					return
				}
			}
		})
	}
}
