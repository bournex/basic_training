package m80

import (
	"fmt"
	"testing"
)

func TestRemoveDuplicates(t *testing.T) {
	testCases := []struct {
		Input   []int
		Expect  []int
		ExpectN int
	}{
		{
			[]int{1, 1, 1, 2, 2, 3},
			[]int{1, 1, 2, 2, 3},
			5,
		},
		{
			[]int{1, 1, 1, 1, 1, 1, 1, 1, 1, 2},
			[]int{1, 1, 2},
			3,
		},
		{
			[]int{1, 1, 1, 1, 1, 1, 1, 1, 1, 1},
			[]int{1, 1},
			2,
		},
		{
			[]int{1},
			[]int{1},
			1,
		},
	}
	for i, v := range testCases {
		t.Run(fmt.Sprintf("case %d", i+1), func(t *testing.T) {
			input := make([]int, len(v.Input))
			copy(input, v.Input)
			res := removeDuplicates1(input)
			if res != v.ExpectN {
				t.Errorf("input %+v, expect N %d, got N %d", v.Input, v.ExpectN, res)
			}
			for j := 0; j < res; j++ {
				if input[j] != v.Expect[j] {
					t.Errorf("input %+v, expect %+v, got %+v", v.Input, v.Expect, input)
				}
			}
		})
	}
}
