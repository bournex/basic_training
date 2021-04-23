package m5

import (
	"fmt"
	"testing"
)

func TestLongestPalindrome(t *testing.T) {
	testCases := []struct {
		Input  string
		Expect string
	}{
		{
			"bb",
			"bb",
		},
		{
			"ac",
			"a",
		},
		{
			"abccb",
			"bccb",
		},
		{
			"abcba",
			"abcba",
		},
		{
			"abcbd",
			"bcb",
		},
		{
			"abad",
			"aba",
		},
		{
			"bcdef",
			"b",
		},
	}
	for i, v := range testCases {
		t.Run(fmt.Sprintf("case %d", i+1), func(t *testing.T) {
			res := longestPalindrome(v.Input)
			if res != v.Expect {
				t.Errorf("input %s, expect %s, got %s", v.Input, v.Expect, res)
			}
		})
	}
}
