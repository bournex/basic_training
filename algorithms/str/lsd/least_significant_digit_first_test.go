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
	}

	for _, v := range corpus {
		Lsd(v.Input, v.Len)
		for i, v1 := range v.Input {
			if v.Expect[i] != v1 {
				t.Errorf("input %+v, len %d, expect %+v, got %+v", v.Input, v.Len, v.Expect, v.Input)
			}
		}
	}
}
