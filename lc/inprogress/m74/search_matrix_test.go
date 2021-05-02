package m74

import (
	"fmt"
	"testing"
)

func TestSearchMatrix(t *testing.T) {
	testCases := []struct {
		Input  [][]int
		Target int
		Expect bool
	}{
		{
			[][]int{{1}},
			1,
			true,
		},
		{
			[][]int{{0}},
			1,
			false,
		},
		{
			[][]int{{1, 3, 5, 70}},
			6,
			false,
		},
		{
			[][]int{{1, 3, 5, 70}},
			70,
			true,
		},
		{
			[][]int{{1}, {3}, {5}, {7}},
			5,
			true,
		},
		{
			[][]int{{1}, {3}, {5}, {7}},
			4,
			false,
		},
		{
			[][]int{{1, 3, 5, 7}, {10, 11, 16, 20}, {23, 30, 34, 60}},
			16,
			true,
		},
		{
			[][]int{{1, 3, 5, 7}, {10, 11, 16, 20}, {23, 30, 34, 60}},
			13,
			false,
		},
	}
	for i, v := range testCases {
		t.Run(fmt.Sprintf("case %d", i+1), func(t *testing.T) {
			res := searchMatrix(v.Input, v.Target)
			if res != v.Expect {
				t.Errorf("input %+v, target %d, expect %t, got %t", v.Input, v.Target, v.Expect, res)
				return
			}
		})
	}
}
