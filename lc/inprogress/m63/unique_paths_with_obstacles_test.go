package m63

import (
	"fmt"
	"testing"
)

func TestUniquePathsWithObstacles(t *testing.T) {
	testCases := []struct {
		Input  [][]int
		Expect int
	}{
		{
			[][]int{{0, 0}, {1, 1}, {0, 0}},
			0,
		},
		{
			[][]int{{0, 1}, {1, 0}},
			0,
		},
		{
			[][]int{{1, 0}, {0, 0}},
			0,
		},
		{
			[][]int{{0, 0}, {1, 0}},
			1,
		},
		{
			[][]int{{0, 0}, {0, 1}},
			0,
		},
		{
			[][]int{{0, 0, 0}, {0, 1, 0}, {0, 0, 0}},
			2,
		},
	}
	for i, v := range testCases {
		t.Run(fmt.Sprintf("case %d", i+1), func(t *testing.T) {
			res := uniquePathsWithObstacles(v.Input)
			if res != v.Expect {
				t.Errorf("input %+v, expect %d, got %d", v.Input, v.Expect, res)
			}

			input := make([][]int, len(v.Input))
			for i := 0; i < len(input); i++ {
				input[i] = make([]int, len(v.Input[i]))
				copy(input[i], v.Input[i])
			}
			res = uniquePathsWithObstaclesSingleLine(input)
			if res != v.Expect {
				t.Errorf("input %+v, expect %d, got %d", v.Input, v.Expect, res)
			}
		})
	}
}
