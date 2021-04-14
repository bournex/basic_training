package q136

import (
	"fmt"
	"testing"
)

func TestSingleNumber(t *testing.T) {
	examples := []struct {
		Input  []int
		Expect int
	}{
		{
			[]int{0, 0, 1, 2, 1, 3, 2},
			3,
		},
		{
			[]int{4},
			4,
		},
		{
			[]int{2, 2, 1},
			1,
		},
	}

	for i, v := range examples {
		t.Run(fmt.Sprintf("case %d", i+1), func(t *testing.T) {
			res := singleNumber(v.Input)
			if res != v.Expect {
				t.Errorf("input %+v, expect %d, got %d", v.Input, v.Expect, res)
			}
		})
	}
}
