package of43

import "testing"

func TestCountDigitOne(t *testing.T) {
	examples := []struct {
		Input  int
		Expect int
	}{
		{
			0,
			0,
		},
		{
			1,
			1,
		},
		{
			-1,
			1,
		},
		{
			12,
			5,
		},
	}

	for _, v := range examples {
		res := countDigitOne(v.Input)
		if res != v.Expect {
			t.Errorf("input %d, expect %d, got %d", v.Input, v.Expect, res)
		}
	}
}
