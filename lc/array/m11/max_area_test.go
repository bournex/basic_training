package m11

import (
	"fmt"
	"testing"
)

func Test(t *testing.T) {
	testCases := []struct {
		Input  []int
		Expect int
	}{
		{
			[]int{1, 8, 6, 2, 5, 4, 8, 3, 7},
			49,
		},
		{
			[]int{1, 1},
			1,
		},
		{
			[]int{4, 3, 2, 1, 4},
			16,
		},
		{
			[]int{1, 2, 1},
			2,
		},
	}

	for i, v := range testCases {
		t.Run(fmt.Sprintf("case %d", i+1), func(t *testing.T) {

			res := maxArea(v.Input)
			if res != v.Expect {
				t.Errorf("input %+v, expect %d, got %d", v.Input, v.Expect, res)
			}
		})
	}
}
