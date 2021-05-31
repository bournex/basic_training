package h32

import (
	"fmt"
	"testing"
)

func TestLongestValidParentheses(t *testing.T) {
	testCases := []struct {
		Input  string
		Expect int
	}{
		{
			"(()(()()))",
			10,
		},
		{
			"()()()",
			6,
		},
		{
			"(()()())",
			8,
		},
		{
			"(()())",
			6,
		},
		{
			"((())())",
			8,
		},
		{
			"(())()",
			6,
		},
		{
			"()))()(()(",
			2,
		},
		{
			")((",
			0,
		},

		{
			"(())",
			4,
		},
		{
			")(()",
			2,
		},
	}
	for i, v := range testCases {
		t.Run(fmt.Sprintf("case %d", i+1), func(t *testing.T) {
			res := longestValidParentheses(v.Input)
			if res != v.Expect {
				t.Errorf("input %s, expect %d, got %d", v.Input, v.Expect, res)
			}
		})
	}
}
