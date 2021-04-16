package m22

import (
	"fmt"
	"sort"
	"testing"
)

func Test(t *testing.T) {
	testCases := []struct {
		Input  int
		Expect []string
	}{
		{
			3,
			[]string{"((()))", "(()())", "(())()", "()(())", "()()()"},
		},
	}
	for i, v := range testCases {
		t.Run(fmt.Sprintf("case %d", i+1), func(t *testing.T) {
			res := generateParenthesis(v.Input)
			sort.Strings(res)
			sort.Strings(v.Expect)

			if len(res) != len(v.Expect) {
				t.Errorf("input %d, expect %+v, got %+v", v.Input, v.Expect, res)
				return
			}

			for j, v1 := range res {
				if v1 != v.Expect[j] {
					t.Errorf("input %d, expect %+v, got %+v", v.Input, v.Expect, res)
					return
				}
			}
		})
	}
}
