package e69

import (
	"fmt"
	"testing"
)

func TestMySqrt(t *testing.T) {
	testCases := []struct {
		Input  int
		Expect int
	}{
		{
			4,
			2,
		},
		{
			8,
			2,
		},
	}
	for i, v := range testCases {
		t.Run(fmt.Sprintf("case %d", i+1), func(t *testing.T) {
			res := mySqrt(v.Input)
			if res != v.Expect {
				t.Errorf("input %d, expect %d, got %d", v.Input, v.Expect, res)
			}
		})
	}
}
