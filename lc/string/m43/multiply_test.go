package m43

import (
	"fmt"
	"testing"
)

func TestMultiply(t *testing.T) {
	testCases := []struct {
		InputA string
		InputB string
		Expect string
	}{
		{
			"123",
			"456",
			"56088",
		},
		{
			"3",
			"4",
			"12",
		},
		{
			"0",
			"123329841932754983745932795473921234",
			"0",
		},
		{
			"74627487",
			"74627487",
			"5569261815935169",
		},
	}
	for i, v := range testCases {
		t.Run(fmt.Sprintf("case %d", i+1), func(t *testing.T) {
			res := multiply(v.InputA, v.InputB)
			if res != v.Expect {
				t.Errorf("input A %s, input B %s, expect %s, got %s", v.InputA, v.InputB, v.Expect, res)
				return
			}
		})
	}
}
