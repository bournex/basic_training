package of42

import "testing"

func TestMaxSubArray(t *testing.T) {
	examples := []struct {
		Input  []int
		Expect int
	}{
		{
			[]int{-2, 1, -3, 4, -1, 2, 1, -5, 4},
			6,
		},
	}
	for _, v := range examples {
		res := maxSubArray(v.Input)
		if res != v.Expect {
			t.Errorf("input %+v, expect %d, got %d", v.Input, v.Expect, res)
		}
	}
}
