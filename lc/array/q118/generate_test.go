package q118

import (
	"fmt"
	"testing"
)

func TestGenerate(t *testing.T) {
	examples := []struct {
		Input  int
		Expect [][]int
	}{
		{
			1,
			[][]int{{1}},
		},
		{
			2,
			[][]int{{1}, {1, 1}},
		},
		{
			3,
			[][]int{{1}, {1, 1}, {1, 2, 1}},
		},
		{
			4,
			[][]int{{1}, {1, 1}, {1, 2, 1}, {1, 3, 3, 1}},
		},
		{
			5,
			[][]int{{1}, {1, 1}, {1, 2, 1}, {1, 3, 3, 1}, {1, 4, 6, 4, 1}},
		},
	}
	for i, v := range examples {
		t.Run(fmt.Sprintf("case%d", i+1), func(t *testing.T) {
			res := generate(v.Input)
			if len(res) != len(v.Expect) {
				t.Errorf("input %d, expect %+v, got %+v", v.Input, v.Expect, res)
				return
			}

			for j, v1 := range res {
				if len(v1) != len(v.Expect[j]) {
					t.Errorf("input %d, expect %+v, got %+v", v.Input, v.Expect, res)
					return
				}

				for k, v2 := range v1 {
					if v2 != v.Expect[j][k] {
						t.Errorf("input %d, expect %+v, got %+v", v.Input, v.Expect, res)
						return
					}
				}
			}
		})
	}
}
