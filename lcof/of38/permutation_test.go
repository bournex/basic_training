package of38

import "testing"

func TestPermutation(t *testing.T) {
	examples := []struct {
		Input  string
		Expect []string
	}{
		{
			"aaaab",
			[]string{"aaaab", "aaaba", "aabaa", "abaaa", "baaaa"},
		},
		{
			"aaaaaa",
			[]string{"aaaaaa"},
		},
		{
			"aab",
			[]string{"aab", "aba", "baa"},
		},
		{
			"ab",
			[]string{"ab", "ba"},
		},
		{
			"abc",
			[]string{"abc", "acb", "bca", "bac", "cba", "cab"},
		},
	}

	for _, v := range examples {
		expect := make(map[string]bool)
		for _, v1 := range v.Expect {
			expect[v1] = true
		}

		res := permutation(v.Input)
		if len(res) != len(expect) {
			t.Errorf("input %s, expect %+v, got %+v", v.Input, v.Expect, res)
		}

		resmap := make(map[string]bool)

		for _, v1 := range res {
			resmap[v1] = true
		}

		for k := range resmap {
			if _, ok := expect[k]; !ok {
				t.Errorf("input %s, expect %+v, got %+v", v.Input, v.Expect, res)
			}
		}
		for k := range expect {
			if _, ok := resmap[k]; !ok {
				t.Errorf("input %s, expect %+v, got %+v", v.Input, v.Expect, res)
			}
		}
	}
}
