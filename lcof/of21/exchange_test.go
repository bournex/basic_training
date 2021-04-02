package of18

import "testing"

func TestExchange(t *testing.T) {
	examples := []struct {
		Input  []int
		Expect []int
	}{
		{
			[]int{1, 2, 3, 4},
			[]int{1, 3, 2, 4},
		},
		{
			[]int{},
			[]int{},
		},
		{
			[]int{1},
			[]int{1},
		},
	}
	for _, v := range examples {
		input := make([]int, len(v.Input))
		copy(input, v.Input)
		res := exchange(v.Input)
		if len(res) != len(v.Expect) {
			t.Errorf("input %+v, expect %+v, got %+v", input, v.Expect, res)
		}

		for i := 0; i < len(res); i++ {
			if res[i] != v.Expect[i] {
				t.Errorf("input %+v, expect %+v, got %+v", input, v.Expect, res)
			}
		}
	}
}
