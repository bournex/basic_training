package m49

import (
	"fmt"
	"sort"
	"testing"
)

func TestGroupAnagrams(t *testing.T) {
	testCases := []struct {
		Input  []string
		Expect [][]string
	}{
		{
			[]string{"eat", "tea", "tan", "ate", "nat", "bat"},
			[][]string{{"eat", "tea", "ate"}, {"tan", "nat"}, {"bat"}},
		},
	}
	for i, v := range testCases {
		t.Run(fmt.Sprintf("case %d", i+1), func(t *testing.T) {
			res := groupAnagrams(v.Input)
			if len(res) != len(v.Expect) {
				t.Errorf("input %+v, expect %+v, got %+v", v.Input, v.Expect, res)
				return
			}

			for j, v1 := range res {
				sort.Strings(v1)
				sort.Strings(v.Expect[j])
				if len(v1) != len(v.Expect[j]) {
					t.Errorf("input %+v, expect %+v, got %+v", v.Input, v.Expect, res)
					return
				}
				for k, v2 := range v1 {
					if v2 != v.Expect[j][k] {
						t.Errorf("input %+v, expect %+v, got %+v", v.Input, v.Expect, res)
						return
					}
				}
			}
		})
	}
}
