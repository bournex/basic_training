package m29

import (
	"fmt"
	"testing"

	"github.com/bournex/basic_training/lc/tree"
)

func TestDivide(t *testing.T) {
	testCases := []struct {
		InputDevidend int
		InputDevisor  int
		Expect        int
	}{
		{
			-2147483648,
			-1,
			2147483647,
		},
		{
			-2147483648,
			-3,
			715827882,
		},
		{
			tree.INT_MAX,
			1000,
			9223372036854775,
		},
		{
			10,
			3,
			3,
		},
		{
			7,
			-3,
			-2,
		},
	}
	for i, v := range testCases {
		t.Run(fmt.Sprintf("case %d", i+1), func(t *testing.T) {
			res := divide(v.InputDevidend, v.InputDevisor)
			if res != v.Expect {
				t.Errorf("input devidend %d, devisor %d, expect %d, got %d", v.InputDevidend, v.InputDevisor, v.Expect, res)
			}
		})
	}
}
