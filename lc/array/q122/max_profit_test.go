package q122

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
			[]int{7, 1, 5, 3, 6, 4},
			7,
		},
		{
			[]int{1, 2, 3, 4, 5},
			4,
		},
	}
	for i, v := range testCases {
		t.Run(fmt.Sprintf("case %d", i+1), func(t *testing.T) {
			res := maxProfit(v.Input)
			if res != v.Expect {
				t.Errorf("input %+v, expect %d, got %d", v.Input, v.Expect, res)
				return
			}
		})
	}
}
