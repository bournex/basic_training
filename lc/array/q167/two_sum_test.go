package q167

import (
	"fmt"
	"testing"
)

func TestTwoSum(t *testing.T) {
	examples := []struct {
		Input  []int
		Target int
		Expect []int
	}{
		{
			[]int{2, 7, 11, 15},
			9,
			[]int{1, 2},
		},
		{
			[]int{2, 3, 4},
			6,
			[]int{1, 3},
		},
		{
			[]int{-1, 0},
			-1,
			[]int{1, 2},
		},
	}

	for i, v := range examples {
		t.Run(fmt.Sprintf("case %d", i+1), func(t *testing.T) {
			res := twoSum(v.Input, v.Target)
			if len(res) != len(v.Expect) {
				t.Errorf("input %+v, target %d, expect %+v, got %+v", v.Input, v.Target, v.Expect, res)
				return
			}
			for j, v1 := range res {
				if v1 != v.Expect[j] {
					t.Errorf("input %+v, target %d, expect %+v, got %+v", v.Input, v.Target, v.Expect, res)
					return
				}
			}
		})
	}
}
