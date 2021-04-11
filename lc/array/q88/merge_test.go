package q88

import (
	"testing"
)

func TestMerge(t *testing.T) {
	examples := []struct {
		Input1 []int
		M      int
		Input2 []int
		N      int
		Expect []int
	}{
		{
			[]int{4, 5, 6, 0, 0, 0},
			3,
			[]int{1, 2, 3},
			3,
			[]int{1, 2, 3, 4, 5, 6},
		},
		{
			[]int{1, 2, 3, 0, 0, 0},
			3,
			[]int{2, 4, 5},
			3,
			[]int{1, 2, 2, 3, 4, 5},
		},
		{
			[]int{0, 0, 0},
			0,
			[]int{2, 4, 5},
			3,
			[]int{2, 4, 5},
		},
		{
			[]int{1, 2, 4, 0},
			3,
			[]int{3},
			1,
			[]int{1, 2, 3, 4},
		},
	}

	for _, v := range examples {
		input1 := make([]int, len(v.Input1))
		copy(input1, v.Input1)
		merge(input1, v.M, v.Input2, v.N)
		if len(input1) != len(v.Expect) {
			t.Errorf("input1 %+v, m %d, input2 %+v, n %d, expect %+v, got %+v", v.Input1, v.M, v.Input2, v.N, v.Expect, input1)
			continue
		}

		for i, v1 := range input1 {
			if v1 != v.Expect[i] {
				t.Errorf("input1 %+v, m %d, input2 %+v, n %d, expect %+v, got %+v", v.Input1, v.M, v.Input2, v.N, v.Expect, input1)
				break
			}
		}
	}
}
