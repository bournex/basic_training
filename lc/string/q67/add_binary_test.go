package q67

import (
	"fmt"
	"testing"
)

func TestAddBinary(t *testing.T) {
	examples := []struct {
		InputA string
		InputB string
		Expect string
	}{
		{
			"100",
			"110010",
			"110110",
		},
		{
			"101",
			"110",
			"1011",
		},
		{
			"11111111",
			"1",
			"100000000",
		},
		{
			"1010101",
			"0",
			"1010101",
		},
		{
			"1111",
			"111",
			"10110",
		},
	}
	for i, v := range examples {
		t.Run(fmt.Sprintf("case%d", i+1), func(t *testing.T) {
			res := addBinary(v.InputA, v.InputB)
			if res != v.Expect {
				t.Errorf("input A %s, input B %s, expect %s, got %s", v.InputA, v.InputB, v.Expect, res)
			}
		})
	}
}
