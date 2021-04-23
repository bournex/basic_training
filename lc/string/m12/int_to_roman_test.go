package m12

import (
	"fmt"
	"testing"
)

func TestIntToRoman(t *testing.T) {
	testCases := []struct {
		Input  int
		Expect string
	}{
		{
			3,
			"III",
		},
		{
			4,
			"IV",
		},
		{
			9,
			"IX",
		},
		{
			1994,
			"MCMXCIV",
		},
		{
			58,
			"LVIII",
		},
		{
			0,
			"",
		},
	}
	for i, v := range testCases {
		t.Run(fmt.Sprintf("case %d", i+1), func(t *testing.T) {
			res := intToRoman(v.Input)
			if res != v.Expect {
				t.Errorf("input %d, expect %s, got %s", v.Input, v.Expect, res)
			}
		})
	}
}
