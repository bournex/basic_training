package m17

import (
	"fmt"
	"sort"
	"testing"
)

func Test(t *testing.T) {
	testCases := []struct {
		Input  string
		Expect []string
	}{
		{
			"",
			[]string{},
		},
		{
			"23",
			[]string{"ad", "ae", "af", "bd", "be", "bf", "cd", "ce", "cf"},
		},
		{
			"2",
			[]string{"a", "b", "c"},
		},
	}
	for i, v := range testCases {
		t.Run(fmt.Sprintf("case %d", i+1), func(t *testing.T) {
			res := letterCombinations1(v.Input)
			sort.Strings(v.Expect)

			if len(res) != len(v.Expect) {
				t.Errorf("input %s, expect %+v, got %+v", v.Input, v.Expect, res)
			}

			sort.Strings(res)
			for j, v1 := range res {
				if v1 != v.Expect[j] {
					t.Errorf("input %s, expect %+v, got %+v", v.Input, v.Expect, res)
				}
			}
		})
	}
}
