package of11

import "testing"

func TestMinArray(t *testing.T) {
	examples := []struct {
		Input  []int
		Expect int
	}{
		{
			[]int{1, 3},
			1,
		},
		{
			[]int{3, 4, 5, 1, 2},
			1,
		},
		{
			[]int{1, 3, 5},
			1,
		},
		{
			[]int{2, 2, 2, 3, 1, 2, 2},
			1,
		},
		{
			[]int{2, 2, 3, 1, 2, 2, 2},
			1,
		},
		{
			[]int{2, 2, 2, 3, 1, 2, 2},
			1,
		},
		{
			[]int{3, 1, 2, 2, 2, 2, 2},
			1,
		},
		{
			[]int{2, 2, 2, 2, 3, 1, 2},
			1,
		},
		{
			[]int{2, 3, 1, 2, 2, 2, 2},
			1,
		},
		{
			[]int{1, 1, 1, 1, 1, 1},
			1,
		},
	}

	for _, v := range examples {
		res := minArray(v.Input)
		if res != v.Expect {
			t.Errorf("input %+v, expect %d, got %d", v.Input, v.Expect, res)
		}
	}
}
