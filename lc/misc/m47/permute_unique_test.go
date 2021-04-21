package m47

import (
	"fmt"
	"testing"
)

func TestPermuteUnique(t *testing.T) {
	testCases := []struct {
		Input  []int
		Expect [][]int
	}{
		{
			[]int{0, 1, 0, 0, 9},
			[][]int{{0, 0, 0, 1, 9}, {0, 0, 0, 9, 1}, {0, 0, 1, 0, 9}, {0, 0, 1, 9, 0}, {0, 0, 9, 0, 1}, {0, 0, 9, 1, 0}, {0, 1, 0, 0, 9}, {0, 1, 0, 9, 0}, {0, 1, 9, 0, 0}, {0, 9, 0, 0, 1}, {0, 9, 0, 1, 0}, {0, 9, 1, 0, 0}, {1, 0, 0, 0, 9}, {1, 0, 0, 9, 0}, {1, 0, 9, 0, 0}, {1, 9, 0, 0, 0}, {9, 0, 0, 0, 1}, {9, 0, 0, 1, 0}, {9, 0, 1, 0, 0}, {9, 1, 0, 0, 0}},
		},
		{
			[]int{0, 0, 0, 1},
			[][]int{{0, 0, 0, 1}, {0, 0, 1, 0}, {0, 1, 0, 0}, {1, 0, 0, 0}},
		},

		{
			[]int{2, 2, 1, 1},
			[][]int{{1, 1, 2, 2}, {1, 2, 1, 2}, {1, 2, 2, 1}, {2, 1, 1, 2}, {2, 1, 2, 1}, {2, 2, 1, 1}},
		},
		{
			[]int{1, 1, 5},
			[][]int{{1, 1, 5}, {1, 5, 1}, {5, 1, 1}},
		},
		{
			[]int{1, 2},
			[][]int{{1, 2}, {2, 1}},
		},
		{
			[]int{1},
			[][]int{{1}},
		},
		{
			[]int{1, 2, 3},
			[][]int{{1, 2, 3}, {1, 3, 2}, {2, 1, 3}, {2, 3, 1}, {3, 1, 2}, {3, 2, 1}},
		},
	}
	for i, v := range testCases {
		t.Run(fmt.Sprintf("case %d", i+1), func(t *testing.T) {
			res := permuteUnique(v.Input)
			if len(res) != len(v.Expect) {
				t.Errorf("input %+v, expect %+v, got %+v", v.Input, v.Expect, res)
			} else {
				t.Logf("input %+v, expect %+v, got %+v", v.Input, v.Expect, res)
			}
		})
	}
}
