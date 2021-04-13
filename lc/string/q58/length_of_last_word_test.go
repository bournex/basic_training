package q58

import (
	"fmt"
	"testing"
)

func TestLengthOfLastWord(t *testing.T) {
	examples := []struct {
		Input  string
		Expect int
	}{
		{
			"hello world",
			5,
		},
		{
			"hello ",
			5,
		},
		{
			" ",
			0,
		},
		{
			"A",
			1,
		},
	}

	for i, v := range examples {
		t.Run(fmt.Sprintf("case%d", i), func(t *testing.T) {
			res := lengthOfLastWord(v.Input)
			if res != v.Expect {
				t.Errorf("input %s, expect %d, got %d", v.Input, v.Expect, res)
			}
		})
	}
}
