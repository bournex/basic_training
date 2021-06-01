package q121

import (
	"fmt"
	"testing"
)

func TestMaxProfit(t *testing.T) {
	testCases := []struct {
		Input  []int
		Expect int
	}{
		{
			[]int{1},
			0,
		},
		{
			[]int{7, 1, 5, 3, 6, 4},
			5,
		},
		{
			[]int{7, 6, 4, 3, 1},
			0,
		},
		{
			[]int{4, 2, 34, 6, 4, 2, 4, 4, 2, 23, 4, 2, 3, 5},
			32,
		},
	}
	for i, v := range testCases {
		t.Run(fmt.Sprintf("case %d", i+1), func(t *testing.T) {
			res := maxProfit(v.Input)
			if res != v.Expect {
				t.Errorf("input %+v, expect %d, got %d", v.Input, v.Expect, res)
			}
		})
	}
}
