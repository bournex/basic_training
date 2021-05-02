package m73

import (
	"fmt"
	"testing"
)

func TestSetZeros(t *testing.T) {
	testCases := []struct {
		Input  [][]int
		Expect [][]int
	}{
		{
			[][]int{{1, 0}},
			[][]int{{0, 0}},
		},
		{
			[][]int{{1, 1, 1}, {1, 0, 1}, {1, 1, 1}},
			[][]int{{1, 0, 1}, {0, 0, 0}, {1, 0, 1}},
		},
		{
			[][]int{{0, 1, 2, 0}, {3, 4, 5, 2}, {1, 3, 1, 5}},
			[][]int{{0, 0, 0, 0}, {0, 4, 5, 0}, {0, 3, 1, 0}},
		},
		{
			[][]int{{0, 1, 2, 0}, {3, 4, 5, 2}, {1, 3, 1, 0}},
			[][]int{{0, 0, 0, 0}, {0, 4, 5, 0}, {0, 0, 0, 0}},
		},
	}
	for i, v := range testCases {
		t.Run(fmt.Sprintf("case %d", i+1), func(t *testing.T) {
			input := make([][]int, len(v.Input))
			for i := 0; i < len(v.Input); i++ {
				input[i] = make([]int, len(v.Input[i]))
				copy(input[i], v.Input[i])
			}

			setZeroes(input)
			for j, v1 := range input {
				for k, v2 := range v1 {
					if v.Expect[j][k] != v2 {
						t.Errorf("input %+v, expect %+v, got %+v", v.Input, v.Expect, input)
						return
					}
				}
			}
		})
	}
}
