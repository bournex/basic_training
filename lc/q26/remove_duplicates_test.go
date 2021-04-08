package q26

import "testing"

func TestRemoveDuplicates(t *testing.T) {
	examples := []struct {
		Input     []int
		Expect    []int
		ExpectNum int
	}{
		{
			[]int{},
			[]int{},
			0,
		},
		{
			[]int{1},
			[]int{1},
			1,
		},
		{
			[]int{1, 1, 2},
			[]int{1, 2},
			2,
		},
	}

	for _, v := range examples {
		input := make([]int, len(v.Input))
		copy(input, v.Input)

		res := removeDuplicates(input)

		if res != v.ExpectNum {
			t.Errorf("input %+v, expect %d, got %d", v.Input, v.ExpectNum, res)
		}

		for i := 0; i < res; i++ {
			if input[i] != v.Expect[i] {
				t.Errorf("input %+v, expect %+v, got %+v", v.Input, v.Expect, input)
			}
		}
	}
}
