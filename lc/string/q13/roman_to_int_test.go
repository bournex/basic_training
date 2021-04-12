package q13

import "testing"

func TestRomanToInt(t *testing.T) {
	examples := []struct {
		Input  string
		Expect int
	}{
		{
			"III",
			3,
		},
		{
			"VIII",
			8,
		},
		{
			"LVIII",
			58,
		},
		{
			"LIV",
			54,
		},
		{
			"MCMXCIV",
			1994,
		},
	}

	for _, v := range examples {
		res := romanToInt(v.Input)
		if res != v.Expect {
			t.Errorf("input %s, expect %d, got %d", v.Input, v.Expect, res)
		}
	}
}
