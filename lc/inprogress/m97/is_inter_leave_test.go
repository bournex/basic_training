package m97

import (
	"fmt"
	"testing"
)

func TestIsInterLeave(t *testing.T) {
	testCases := []struct {
		InputA string
		InputB string
		Target string
		Expect bool
	}{
		{
			"aabcc",
			"dbbca",
			"aadbbcbcac",
			true,
		},
		{
			"aabcc",
			"dbbca",
			"aadbbbaccc",
			false,
		},
		{
			"",
			"",
			"",
			true,
		},
		{
			"",
			"ab",
			"ab",
			true,
		},
		{
			"",
			"abd",
			"abc",
			false,
		},
	}
	for i, v := range testCases {
		t.Run(fmt.Sprintf("%d", i+1), func(t *testing.T) {
			res := isInterleave(v.InputA, v.InputB, v.Target)
			if res != v.Expect {
				t.Errorf("string A %s, string B %s, target %s, expect %t, got %t", v.InputA, v.InputB, v.Target, v.Expect, res)
			}
		})
	}
}
