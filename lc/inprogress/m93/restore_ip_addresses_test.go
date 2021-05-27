package m93

import (
	"fmt"
	"sort"
	"testing"
)

func TestRestoreIpAddresses(t *testing.T) {
	testCases := []struct {
		Input  string
		Expect []string
	}{
		{
			"25525511135",
			[]string{"255.255.11.135", "255.255.111.35"},
		},
		{
			"0000",
			[]string{"0.0.0.0"},
		},
		{
			"255255255255",
			[]string{"255.255.255.255"},
		},
		{
			"010010",
			[]string{"0.10.0.10", "0.100.1.0"},
		},
		{
			"12121212",
			[]string{},
		},
	}
	for i, v := range testCases {
		t.Run(fmt.Sprintf("case %d", i+1), func(t *testing.T) {
			res := restoreIpAddresses(v.Input)
			if len(res) != len(v.Expect) {
				t.Errorf("input %s, expect %+v, got %+v", v.Input, v.Expect, res)
				return
			}

			sort.Strings(v.Expect)
			sort.Strings(res)

			for j, v1 := range res {
				if v1 != v.Expect[j] {
					t.Errorf("input %s, expect %+v, got %+v", v.Input, v.Expect, res)
					return
				}
			}
		})
	}
}
