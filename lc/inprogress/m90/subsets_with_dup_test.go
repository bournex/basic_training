package m90

import (
	"fmt"
	"testing"
)

func TestSubsetsWithDup(t *testing.T) {
	testCases := []struct {
		Input  []int
		Expect [][]int
	}{
		{
			[]int{0},
			[][]int{{}, {0}},
		},
		{
			[]int{1, 2, 2},
			[][]int{{}, {1}, {1, 2}, {1, 2, 2}, {2}, {2, 2}},
		},
	}
	for i, v := range testCases {
		t.Run(fmt.Sprintf("case %d", i+1), func(t *testing.T) {
			input := make([]int, len(v.Input))
			copy(input, v.Input)
			res := subsetsWithDup(input)
			if len(res) != len(v.Expect) {
				t.Errorf("input %+v, expect %+v, res %+v", v.Input, v.Expect, res)
				return
			}

			for j, v1 := range res {
				for k, v2 := range v1 {
					if v2 != v.Expect[j][k] {
						t.Errorf("input %+v, expect %+v, res %+v", v.Input, v.Expect, res)
						return
					}
				}
			}
		})
	}
}
