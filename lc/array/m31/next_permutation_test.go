package m31

import (
	"fmt"
	"testing"
)

func Test(t *testing.T) {
	testCases := []struct {
		Input  []int
		Expect []int
	}{
		{
			[]int{1, 2, 3},
			[]int{1, 3, 2},
		},
		{
			[]int{3, 2, 1},
			[]int{1, 2, 3},
		},
	}
	for i, v := range testCases {
		t.Run(fmt.Sprintf("case %d", i+1), func(t *testing.T) {
			input := make([]int, len(v.Input))
			copy(input, v.Input)
			nextPermutation(input)

			for i, v1 := range v.Input {
				if v1 != v.Expect[i] {
					t.Errorf("input %+v, expect %+v, got %+v", v.Input, v.Expect, input)
					return
				}
			}
		})
	}
}
