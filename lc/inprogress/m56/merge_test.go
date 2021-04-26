package m56

import (
	"fmt"
	"testing"
)

func TestMerge(t *testing.T) {
	testCases := []struct {
		Input  [][]int
		Expect [][]int
	}{
		{
			[][]int{{1, 10}, {2, 9}, {3, 8}, {4, 7}},
			[][]int{{1, 10}},
		},
		{
			[][]int{{1, 3}, {3, 5}, {5, 7}, {7, 8}},
			[][]int{{1, 8}},
		},
		{
			[][]int{{1, 2}, {3, 4}, {5, 6}, {7, 8}},
			[][]int{{1, 2}, {3, 4}, {5, 6}, {7, 8}},
		},
		{
			[][]int{{1, 3}, {2, 6}, {8, 10}, {15, 18}},
			[][]int{{1, 6}, {8, 10}, {15, 18}},
		},
		{
			[][]int{{1, 4}, {4, 5}},
			[][]int{{1, 5}},
		},
	}
	for i, v := range testCases {
		t.Run(fmt.Sprintf("case %d", i+1), func(t *testing.T) {
			res := merge(v.Input)
			if len(res) != len(v.Expect) {
				t.Errorf("input %+v, expect %+v, got %+v", v.Input, v.Expect, res)
				return
			}

			for j, v1 := range res {
				if v1[0] != v.Expect[j][0] || v1[1] != v.Expect[j][1] {
					t.Errorf("input %+v, expect %+v, got %+v", v.Input, v.Expect, res)
					return
				}
			}
		})
	}
}
