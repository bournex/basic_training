package of15

import "testing"

func TestHammingWeight(t *testing.T) {
	examples := []struct {
		Input  uint32
		Expect int
	}{
		{
			uint32(0b_00001111000011110000111100001111),
			16,
		},
		{
			uint32(0b_00000000000000000000000000000000),
			0,
		},
		{
			uint32(0b_10000000000000000000000000000000),
			1,
		},
		{
			uint32(0b_11111111111111111111111111111111),
			32,
		},
	}

	for _, v := range examples {
		res := hammingWeight(v.Input)
		if res != v.Expect {
			t.Errorf("input %d, expect %d, got %d", v.Input, v.Expect, res)
		}
	}
}
