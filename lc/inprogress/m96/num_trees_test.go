package m96

import (
	"fmt"
	"testing"
)

func TestNumTrees(t *testing.T) {
	testCases := []struct {
		Input  int
		Expect int
	}{
		{
			19,
			1767263190,
		},
		{
			3,
			5,
		},
		{
			1,
			1,
		},
		{
			2,
			2,
		},
	}
	for i, v := range testCases {
		t.Run(fmt.Sprintf("%d", i+1), func(t *testing.T) {
			res := numTrees(v.Input)
			if res != v.Expect {
				t.Errorf("input %d, expect %d, got %d", v.Input, v.Expect, res)
			}
		})
	}
}
