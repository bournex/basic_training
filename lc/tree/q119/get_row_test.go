package q119

import (
	"fmt"
	"testing"
)

func TestGetRow(t *testing.T) {
	examples := []struct {
		Input  int
		Expect []int
	}{
		{
			1,
			[]int{1},
		},
		{
			2,
			[]int{1, 1},
		},
		{
			3,
			[]int{1, 3, 3, 1},
		},
		{
			4,
			[]int{1, 4, 6, 4, 1},
		},
		{
			5,
			[]int{1, 5, 10, 10, 5, 1},
		},
	}

	for i, v := range examples {
		t.Run(fmt.Sprintf("case%d", i+1), func(t *testing.T) {
			res := getRow(v.Input)
			if len(res) != len(v.Expect) {
				t.Errorf("input %d, expect %+v, got %+v", v.Input, v.Expect, res)
				return
			}

			for j, v1 := range res {
				if v1 != v.Expect[j] {
					t.Errorf("input %d, expect %+v, got %+v", v.Input, v.Expect, res)
					return
				}
			}
		})
	}
}
