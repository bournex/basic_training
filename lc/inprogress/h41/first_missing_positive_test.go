package h41

import (
	"fmt"
	"testing"
)

func Test(t *testing.T) {
	testCases := []struct {
		Input  []int
		Expect int
	}{
		{},
	}
	for i, v := range testCases {
		t.Run(fmt.Sprintf("case %d", i+1), func(t *testing.T) {
			res := firstMissingPositive(v.Input)
			if res != v.Expect {
				t.Errorf("input %+v, expect %d, got %d", v.Input, v.Expect, res)
				return
			}
		})
	}
}
