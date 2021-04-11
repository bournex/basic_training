package q35

import "testing"

func TestSearchInsert(t *testing.T) {
	examples := []struct {
		Input  []int
		Target int
		Expect int
	}{
		{
			[]int{1, 3, 5, 6},
			7,
			4,
		},
		{
			[]int{1, 3, 5, 6},
			0,
			0,
		},
		{
			[]int{1, 3, 5, 6},
			2,
			1,
		},
		{
			[]int{1, 3, 5, 6},
			5,
			2,
		},

		{
			[]int{},
			1,
			0,
		},
		{
			[]int{1},
			2,
			1,
		},
	}

	for _, v := range examples {
		res := searchInsert(v.Input, v.Target)
		if res != v.Expect {
			t.Errorf("input %+v, target %d, expect %d, got %d", v.Input, v.Target, v.Expect, res)
		}
	}
}
