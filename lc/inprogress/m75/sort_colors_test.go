package m75

import (
	"fmt"
	"testing"
)

func TestSortColors(t *testing.T) {
	testCases := []struct {
		Input  []int
		Expect []int
	}{
		{
			[]int{2, 0, 2, 1, 1, 0},
			[]int{0, 0, 1, 1, 2, 2},
		},
		{
			[]int{2, 0, 1, 2, 1, 0, 2, 0, 1, 2, 1, 0},
			[]int{0, 0, 0, 0, 1, 1, 1, 1, 2, 2, 2, 2},
		},
	}
	for i, v := range testCases {
		t.Run(fmt.Sprintf("case %d", i+1), func(t *testing.T) {
			input := make([]int, len(v.Input))
			copy(input, v.Input)
			sortColors(input)
			for j, v1 := range input {
				if v1 != v.Expect[j] {
					t.Errorf("input %+v, expect %+v, got %+v", v.Input, v.Expect, input)
					return
				}
			}
		})
	}
}
