package q14

import "testing"

func TestLongestCommonPrefix(t *testing.T) {
	examples := []struct {
		Input  []string
		Expect string
	}{
		{
			[]string{"helo", "hello", "hell"},
			"hel",
		},
		{
			[]string{"helo"},
			"helo",
		},
		{
			[]string{"hello", "world"},
			"",
		},
		{
			[]string{"flower", "flow", "flight"},
			"fl",
		},
	}

	for _, v := range examples {
		res := longestCommonPrefix(v.Input)
		if res != v.Expect {
			t.Errorf("input %+v, expect %s, got %s", v.Input, v.Expect, res)
		}
	}
}
