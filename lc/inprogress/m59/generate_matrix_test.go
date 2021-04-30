package m59

import (
	"fmt"
	"testing"
)

func TestGenerateMatrix(t *testing.T) {
	testCases := []struct {
		Input  int
		Expect [][]int
	}{
		{
			0,
			[][]int{},
		},
		{
			1,
			[][]int{{1}},
		},
		{
			2,
			[][]int{{1, 2}, {4, 3}},
		},
		{
			3,
			[][]int{{1, 2, 3}, {8, 9, 4}, {7, 6, 5}},
		},
		{
			4,
			[][]int{{1, 2, 3, 4}, {12, 13, 14, 5}, {11, 16, 15, 6}, {10, 9, 8, 7}},
		},
	}
	for i, v := range testCases {
		t.Run(fmt.Sprintf("case %d", i+1), func(t *testing.T) {
			res := generateMatrix(v.Input)
			if len(res) != len(v.Expect) {
				t.Errorf("input %d, expect %+v, got %+v", v.Input, v.Expect, res)
				return
			}

			for j, v1 := range res {
				for k, v2 := range v1 {
					if v2 != v.Expect[j][k] {
						t.Errorf("input %d, expect %+v, got %+v", v.Input, v.Expect, res)
						return
					}
				}
			}
		})
	}
}
