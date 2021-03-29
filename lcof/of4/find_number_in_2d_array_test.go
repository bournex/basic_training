package of4

import "testing"

func TestFindNumberIn2DArray(t *testing.T) {

	examples := []struct {
		Input  [][]int
		Target int
		Expect bool
	}{
		{
			[][]int{{1, 4, 7, 11, 15}, {2, 5, 8, 12, 19}, {3, 6, 9, 16, 22}, {10, 13, 14, 17, 24}, {18, 21, 23, 26, 30}},
			5,
			true,
		},
		{
			[][]int{},
			0,
			false,
		},
		{
			[][]int{{-5}},
			-5,
			true,
		},
	}

	for _, v := range examples {
		res := findNumberIn2DArray(v.Input, v.Target)
		if res != v.Expect {
			t.Errorf("input %+v, target %d, expect %+v, got %+v", v.Input, v.Target, v.Expect, res)
		}
	}
}
