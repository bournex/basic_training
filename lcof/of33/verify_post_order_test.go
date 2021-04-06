package of33

import "testing"

func TestVerfiyPostOrder(t *testing.T) {
	examples := []struct {
		Input  []int
		Expect bool
	}{
		{
			[]int{1, 2, 5, 10, 6, 9, 4, 3},
			false,
		},
		{
			[]int{5, 7, 6, 9, 11, 10, 8},
			true,
		},
		{
			[]int{7, 4, 6, 5},
			false,
		},
		{
			[]int{2, 3},
			true,
		},
		{
			[]int{3, 2},
			true,
		},
		{
			[]int{1},
			true,
		},
	}

	for _, v := range examples {
		res := verifyPostorder(v.Input)
		if res != v.Expect {
			t.Errorf("input %+v, expect %t, got %t", v.Input, v.Expect, res)
		}
	}
}
