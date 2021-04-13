package q38

import (
	"fmt"
	"testing"
)

func TestCountAndSay(t *testing.T) {
	examples := []struct {
		Input  int
		Expect string
	}{
		{
			2,
			"11",
		},
		{
			5,
			"111221",
		},
	}

	for i, v := range examples {
		t.Run(fmt.Sprintf("case%d", i), func(t *testing.T) {
			res := countAndSay(v.Input)
			if res != v.Expect {
				t.Errorf("input %d, expect %s, got %s", v.Input, v.Expect, res)
			}
		})
	}
}
