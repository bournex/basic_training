package q169

import (
	"fmt"
	"testing"
)

func TestMajorityElement(t *testing.T) {
	examples := []struct {
		Input  []int
		Expect int
	}{
		{
			[]int{2, 2, 1, 1, 1, 2, 2},
			2,
		},
		{
			[]int{3, 5, 3},
			3,
		},
		{
			[]int{3, 2, 3},
			3,
		},

		{
			[]int{2, 1, 1, 1, 1, 2, 2},
			1,
		},
	}

	for i, v := range examples {
		t.Run(fmt.Sprintf("case %d", i+1), func(t *testing.T) {
			res := majorityElement(v.Input)
			if res != v.Expect {
				t.Errorf("input %+v, expect %d, got %d", v.Input, v.Expect, res)
				return
			}
		})
	}
}
