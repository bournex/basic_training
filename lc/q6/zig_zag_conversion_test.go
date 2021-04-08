package q6

import "testing"

func TestZigZagConversion(t *testing.T) {
	sample := "PAYPALISHIRING"
	expect := "PINALSIGYAHRPI"

	res := convert(sample, 4)
	if res != expect {
		t.Errorf("expect %s, got %s", expect, res)
	}
}

func TestZigZagConversion1(t *testing.T) {
	sample := "AABAAB"
	expect := "AABAAB"

	res := convert(sample, 1)
	if res != expect {
		t.Errorf("expect %s, got %s", expect, res)
	}
}
