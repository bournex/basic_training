package q163

import (
	"fmt"
	"testing"
)

func TestFindMissingRanges(t *testing.T) {
	examples := []struct {
		Input  []int
		Lower  int
		Upper  int
		Expect []string
	}{

		{
			[]int{0, 9},
			0,
			9,
			[]string{"1->8"},
		},
		{
			[]int{-1},
			-2,
			-1,
			[]string{"-2"},
		},
		{
			[]int{0, 1, 3, 50, 75},
			0,
			99,
			[]string{"2", "4->49", "51->74", "76->99"},
		},
		{
			[]int{3, 4, 8, 50, 75},
			0,
			99,
			[]string{"0->2", "5->7", "9->49", "51->74", "76->99"},
		},
	}

	for i, v := range examples {
		t.Run(fmt.Sprintf("case %d", i+1), func(t *testing.T) {
			res := findMissingRanges(v.Input, v.Lower, v.Upper)
			if len(res) != len(v.Expect) {
				t.Errorf("input %+v, lower %d, upper %d, expect %+v, got %+v", v.Input, v.Lower, v.Upper, v.Expect, res)
				return
			}
			for j, v1 := range res {
				if v1 != v.Expect[j] {
					t.Errorf("input %+v, lower %d, upper %d, expect %+v, got %+v", v.Input, v.Lower, v.Upper, v.Expect, res)
					return
				}
			}
		})
	}
}
