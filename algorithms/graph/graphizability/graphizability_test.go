package graphizability

import (
	"fmt"
	"testing"
)

func TestHavelHakimi(t *testing.T) {
	testCases := []struct {
		Input  []int
		Expect bool
	}{
		{
			[]int{1, 0, 0},
			false,
		},
		{
			[]int{7, 7, 4, 3, 3, 3, 2, 1},
			false,
		},
		{
			[]int{5, 4, 3, 3, 2, 2, 2, 1, 1, 1},
			true,
		},
	}
	for i, v := range testCases {
		t.Run(fmt.Sprintf("case %d", i+1), func(t *testing.T) {
			input := make([]int, len(v.Input))
			copy(input, v.Input)
			res := havelHakimi(input)
			if res != v.Expect {
				t.Errorf("input %+v, expect %t, got %t", v.Input, v.Expect, res)
			}
		})
	}
}
