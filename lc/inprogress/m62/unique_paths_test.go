package m62

import (
	"fmt"
	"testing"
)

func TestUniquePaths(t *testing.T) {
	testCases := []struct {
		InputM int
		InputN int
		Expect int
	}{
		{
			23,
			12,
			6,
		},
		{
			3,
			3,
			6,
		},
		{
			3,
			2,
			3,
		},
		{
			3,
			7,
			28,
		},
	}
	for i, v := range testCases {
		t.Run(fmt.Sprintf("case %d", i+1), func(t *testing.T) {
			res := uniquePaths(v.InputM, v.InputN)
			if res != v.Expect {
				t.Errorf("input m %d, input n %d, expect %d, got %d", v.InputM, v.InputN, v.Expect, res)
			}
		})
	}
}
