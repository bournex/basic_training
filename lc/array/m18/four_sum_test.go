package m18

import (
	"fmt"
	"testing"
)

func Test(t *testing.T) {
	testCases := []struct {
		Input  []int
		Target int
		Expect [][]int
	}{
		{
			[]int{1, 0, -1, 0, -2, 2},
			0,
			[][]int{{-2, -1, 1, 2}, {-2, 0, 0, 2}, {-1, 0, 0, 1}},
		},
	}
	for i, v := range testCases {
		t.Run(fmt.Sprintf("case %d", i+1), func(t *testing.T) {
			res := fourSum(v.Input, v.Target)

			if len(res) != len(v.Expect) {
				t.Errorf("input %+v, target %d, expect %+v, got %+v", v.Input, v.Target, v.Expect, res)
			}
		})
	}
}
