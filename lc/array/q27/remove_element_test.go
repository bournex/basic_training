package q27

import "testing"

func TestRemoveElement(t *testing.T) {
	examples := []struct {
		Input  []int
		Num    int
		Expect []int
		ENum   int
	}{
		{
			[]int{3, 2, 2, 3},
			3,
			[]int{2, 2},
			2,
		},
	}

	for _, v := range examples {
		input := make([]int, len(v.Input))
		copy(input, v.Input)
		res := removeElement(v.Input, v.Num)

		if res != v.ENum {
			t.Errorf("input %+v, num %d, expect %d, got %d", input, v.Num, v.ENum, res)
		}

		for i := 0; i < res; i++ {
			if v.Input[i] != v.Expect[i] {
				t.Errorf("input %+v, expect %+v, got %+v", input, v.Expect, v.Input)
			}
		}
	}
}
