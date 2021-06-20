package h52

import (
	"fmt"
	"testing"
)

func TestTotalNQueens(t *testing.T) {
	testCases := []struct {
		Input  int
		Expect int
	}{
		{
			1,
			1,
		},
		{
			2,
			0,
		},
		{
			3,
			0,
		},
		{
			4,
			2,
		},
		{
			5,
			10,
		},
	}
	for i, v := range testCases {
		t.Run(fmt.Sprintf("%d", i+1), func(t *testing.T) {
			res := totalNQueens(v.Input)
			if res != v.Expect {
				t.Errorf("input %d, expect %d, got %d", v.Input, v.Expect, res)
			}
		})
	}
}
