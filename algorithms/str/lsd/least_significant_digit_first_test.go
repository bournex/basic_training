package lsd

import "testing"

func TestLeastSignificantDigitFirst(t *testing.T) {
	corpus := []struct {
		Input  []string
		Len    int
		Expect []string
	}{
		{
			[]string{"space", "hello", "world"},
			5,
			[]string{"hello", "space", "world"},
		},
		{
			[]string{},
			0,
			[]string{},
		},
		{
			[]string{"hello"},
			5,
			[]string{"hello"},
		},
		{
			[]string{
				"abcdc",
				"wtomm",
				"uncle",
				"hello",
				"world",
				"abcde",
				"abcce",
				"abcdf",
				"which",
				"heavy",
				"sunny",
				"happy",
			},
			5,
			[]string{
				"abcce",
				"abcdc",
				"abcde",
				"abcdf",
				"happy",
				"heavy",
				"hello",
				"sunny",
				"uncle",
				"which",
				"world",
				"wtomm",
			},
		},
	}

	for _, v := range corpus {
		cp := make([]string, len(v.Input))
		copy(cp, v.Input)
		Lsd(v.Input, v.Len)
		for i, v1 := range v.Input {
			if v.Expect[i] != v1 {
				t.Errorf("input %+v, len %d, expect %+v, got %+v", cp, v.Len, v.Expect, v.Input)
			}
		}
	}
}
