package m78

import (
	"fmt"
	"testing"
)

func TestSubSets(t *testing.T) {
	testCases := []struct {
		Input  []int
		Expect [][]int
	}{
		{
			[]int{1, 2, 3},
			[][]int{{}, {1}, {1, 2}, {1, 2, 3}, {1, 3}, {2}, {2, 3}, {3}},
		},
		{
			[]int{1, 2},
			[][]int{{}, {1}, {1, 2}, {2}},
		},
		{
			[]int{1, 2, 3, 4},
			[][]int{{}, {1}, {1, 2}, {1, 2, 3}, {1, 2, 3, 4}, {1, 2, 4}, {1, 3}, {1, 3, 4}, {1, 4}, {2}, {2, 3}, {2, 3, 4}, {2, 4}, {3}, {3, 4}, {4}},
		},
	}
	for i, v := range testCases {
		t.Run(fmt.Sprintf("case %d", i+1), func(t *testing.T) {
			res := subsets(v.Input)

			if len(res) != len(v.Expect) {
				t.Errorf("input %+v, expect %+v, got %+v", v.Input, v.Expect, res)
			}

			for j, v1 := range res {
				for k, v2 := range v1 {
					if v2 != v.Expect[j][k] {
						t.Errorf("input %+v, expect %+v, got %+v", v.Input, v.Expect, res)
					}
				}
			}
		})
	}
}
