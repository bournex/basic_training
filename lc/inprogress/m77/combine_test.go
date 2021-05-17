package m77

import (
	"fmt"
	"testing"
)

func TestCombine(t *testing.T) {
	testCases := []struct {
		InputN int
		InputK int
		Expect [][]int
	}{
		{
			3,
			2,
			[][]int{{1, 2}, {1, 3}, {2, 3}},
		},
		{
			4,
			2,
			[][]int{{2, 4}, {3, 4}, {2, 3}, {1, 2}, {1, 3}, {1, 4}},
		},
	}
	for i, v := range testCases {
		t.Run(fmt.Sprintf("case %d", i+1), func(t *testing.T) {
			res := combine(v.InputN, v.InputK)
			if len(res) != len(v.Expect) {
				t.Errorf("input n %d k %d, expect %+v, got %+v", v.InputN, v.InputK, v.Expect, res)
				return
			}
		})
	}
}
