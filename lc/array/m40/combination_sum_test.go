package m40

import (
	"fmt"
	"testing"
)

func TestCombinationSum(t *testing.T) {
	testCases := []struct {
		Input  []int
		Target int
		Expect [][]int
	}{
		{
			[]int{1, 2, 5},
			8,
			[][]int{{1, 2, 5}},
		},
		{
			[]int{10, 1, 2, 7, 6, 1, 5},
			8,
			[][]int{{1, 7}, {1, 2, 5}, {2, 6}, {1, 1, 6}},
		},
		{
			[]int{2, 5, 2, 1, 2},
			5,
			[][]int{{1, 2, 2}, {5}},
		},
	}
	for i, v := range testCases {
		t.Run(fmt.Sprintf("case %d", i+1), func(t *testing.T) {
			res := combinationSum2(v.Input, v.Target)
			if len(res) != len(v.Expect) {
				t.Errorf("input %+v, target %d, expect %+v, got %+v", v.Input, v.Target, v.Expect, res)
			} else {
				t.Logf("expect %+v, got %+v", v.Expect, res)
			}
		})
	}
}
