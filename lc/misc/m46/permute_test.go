package m46

import (
	"fmt"
	"testing"
)

func TestPermute(t *testing.T) {
	testCases := []struct {
		Input  []int
		Expect [][]int
	}{
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
			res := permute(v.Input)
			if len(res) != len(v.Expect) {
				t.Errorf("input %+v, expect %+v, got %+v", v.Input, v.Expect, res)
			} else {
				t.Logf("input %+v, expect %+v, got %+v", v.Input, v.Expect, res)
			}
		})
	}
}
