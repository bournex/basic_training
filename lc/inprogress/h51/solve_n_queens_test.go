package h51

import (
	"fmt"
	"testing"
)

func TestSolveNQueens(t *testing.T) {
	testCases := []struct {
		Input  int
		Expect [][]string
	}{
		{
			1,
			[][]string{{"Q"}},
		},
		{
			4,
			[][]string{{""}},
		},
	}
	for i, v := range testCases {
		t.Run(fmt.Sprintf("%d", i+1), func(t *testing.T) {
			res := solveNQueens(v.Input)
			// if len(res) != len(v.Expect) {
			// 	t.Errorf("input %d, expect %+v, got %+v", v.Input, v.Expect, res)
			// 	return
			// }
			t.Logf("res %+v", res)
		})
	}
}
