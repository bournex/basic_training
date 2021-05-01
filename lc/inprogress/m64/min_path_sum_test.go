package m64

import (
	"fmt"
	"testing"
)

func TestMinPathSum(t *testing.T) {
	testCases := []struct {
		Input  [][]int
		Expect int
	}{
		{
			[][]int{{1, 3, 1}, {1, 5, 1}, {4, 2, 1}},
			7,
		},
		{
			[][]int{{1, 2, 3}, {4, 5, 6}},
			12,
		},
		{
			[][]int{{0, 0, 1}, {1, 6, 5}, {4, 0, 0}},
			5,
		},
	}
	for i, v := range testCases {
		t.Run(fmt.Sprintf("case %d", i+1), func(t *testing.T) {
			res := minPathSum(v.Input)
			if res != v.Expect {
				t.Errorf("input %+v, expect %d, res %d", v.Input, v.Expect, res)
			}
		})
	}
}
