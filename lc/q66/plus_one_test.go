package q66

import "testing"

func TestPlusOne(t *testing.T) {
	examples := []struct {
		Input  []int
		Expect []int
	}{
		{
			[]int{0},
			[]int{1},
		},
		{
			[]int{9},
			[]int{1, 0},
		},
		{
			[]int{9, 9, 9},
			[]int{1, 0, 0, 0},
		},
	}

	for _, v := range examples {
		input := make([]int, len(v.Input))
		copy(input, v.Input)
		res := plusOne(input)
		if len(res) != len(v.Expect) {
			t.Errorf("input %+v, expect %+v, got %+v", v.Input, v.Expect, res)
		}

		for i, v1 := range res {
			if v1 != v.Expect[i] {
				t.Errorf("input %+v, expect %+v, got %+v", v.Input, v.Expect, res)
			}
		}
	}
}
