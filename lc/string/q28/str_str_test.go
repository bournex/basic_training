package q28

import (
	"fmt"
	"testing"
)

func TestStrStr(t *testing.T) {
	examples := []struct {
		Input  string
		Needle string
		Expect int
	}{
		{
			"hello",
			"ll",
			2,
		},
		{
			"",
			"ll",
			-1,
		},
		{
			"world",
			"ac",
			-1,
		},
		{
			"query",
			"",
			0,
		},
	}

	for i, v := range examples {
		t.Run(fmt.Sprintf("case%d", i), func(t *testing.T) {
			res := strStr(v.Input, v.Needle)
			if res != v.Expect {
				t.Errorf("input %s, needle %s, expect %d, got %d", v.Input, v.Needle, v.Expect, res)
			}
		})
	}
}
