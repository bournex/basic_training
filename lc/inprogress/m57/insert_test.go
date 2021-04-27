package m57

import (
	"fmt"
	"testing"
)

func TestInsert(t *testing.T) {
	testCases := []struct {
		Input       [][]int
		NewInterval []int
		Expect      [][]int
	}{
		{
			[][]int{{1, 3}, {6, 9}},
			[]int{2, 5},
			[][]int{{1, 5}, {6, 9}},
		},
	}
	for i, v := range testCases {
		t.Run(fmt.Sprintf("case %d", i+1), func(t *testing.T) {
			res := insert(v.Input, v.NewInterval)
			if len(res) != len(v.Expect) {
				t.Errorf("input %+v, new interval %+v, expect %+v, got %+v", v.Input, v.NewInterval, v.Expect, res)
				return
			}

			for j, v1 := range res {
				if v1[0] != v.Expect[j][0] || v1[1] != v.Expect[j][1] {
					t.Errorf("input %+v, new interval %+v, expect %+v, got %+v", v.Input, v.NewInterval, v.Expect, res)
					return
				}
			}
		})
	}
}
