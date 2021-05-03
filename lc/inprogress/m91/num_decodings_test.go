package m91

import (
	"fmt"
	"testing"
)

func TestNumDecodings(t *testing.T) {
	testCases := []struct {
		Input  string
		Expect int
	}{
		{
			"2101",
			1,
		},
		{
			"1000",
			0,
		},
		{
			"70",
			0,
		},
		{
			"12",
			2,
		},
		{
			"11111111",
			34,
		},
	}
	for i, v := range testCases {
		t.Run(fmt.Sprintf("case %d", i+1), func(t *testing.T) {
			res := numDecodings(v.Input)
			if res != v.Expect {
				t.Errorf("input %s, expect %d, got %d", v.Input, v.Expect, res)
			}
		})
	}
}
