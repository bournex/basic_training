package h30

import (
	"fmt"
	"testing"
)

func TestFindSubstring(t *testing.T) {
	testCases := []struct {
		Input  string
		Words  []string
		Expect []int
	}{
		{
			"foobarfootool",
			[]string{"foo", "bar"},
			[]int{0, 3},
		},
		{
			"fooabarbfootool",
			[]string{"foo", "bar"},
			[]int{},
		},
		{
			"fooaaaabfootool",
			[]string{"aa", "b"},
			[]int{5},
		},
		{
			"fooaaaabfootool",
			[]string{"aa", "a"},
			[]int{3, 4},
		},
	}
	for i, v := range testCases {
		t.Run(fmt.Sprintf("case %d", i+1), func(t *testing.T) {
			res := findSubstring(v.Input, v.Words)
			if len(res) != len(v.Expect) {
				t.Errorf("input %s, words %+v, expect %+v, got %+v", v.Input, v.Words, v.Expect, res)
				return
			}
			for j, v1 := range res {
				if v1 != v.Expect[j] {
					t.Errorf("input %s, words %+v, expect %+v, got %+v", v.Input, v.Words, v.Expect, res)
					return
				}
			}
		})
	}
}
