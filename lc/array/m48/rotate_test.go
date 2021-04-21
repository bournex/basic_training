package m48

import (
	"fmt"
	"testing"
)

func TestRotate(t *testing.T) {
	testCases := []struct {
		Input  [][]int
		Expect [][]int
	}{
		{
			[][]int{},
			[][]int{},
		},
		{
			[][]int{{1}},
			[][]int{{1}},
		},
		{
			[][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}},
			[][]int{{7, 4, 1}, {8, 5, 2}, {9, 6, 3}},
		},
		{
			[][]int{{5, 1, 9, 11}, {2, 4, 8, 10}, {13, 3, 6, 7}, {15, 14, 12, 16}},
			[][]int{{15, 13, 2, 5}, {14, 3, 4, 1}, {12, 6, 8, 9}, {16, 7, 10, 11}},
		},
	}
	for i, v := range testCases {
		t.Run(fmt.Sprintf("case %d", i+1), func(t *testing.T) {
			input := make([][]int, len(v.Input))
			for i, v := range v.Input {
				input[i] = make([]int, len(v))
				copy(input[i], v)
			}
			rotate(input)

			for i := 0; i < len(v.Input); i++ {
				for j := 0; j < len(v.Input[0]); j++ {
					if input[i][j] != v.Expect[i][j] {
						t.Errorf("input %+v, expect %+v, got %+v", v.Input, v.Expect, input)
						return
					}
				}
			}
		})
	}
}
