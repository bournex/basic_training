package m50

import (
	"fmt"
	"testing"
)

func TestMyPow(t *testing.T) {
	testCases := []struct {
		Input  float64
		N      int
		Expect float64
	}{
		{
			2.0000,
			-2,
			0.25,
		},
		{
			2.0000,
			10,
			1024.0000,
		},
		{
			2.0000,
			2,
			4.0000,
		},

		{
			2.1000,
			3,
			9.2610,
		},
		{
			-2.1000,
			3,
			-9.2610,
		},
	}
	const MIN = 0.000001
	// MIN 为用户自定义的比较精度
	isEqual := func(f1, f2 float64) bool {
		if f1 > f2 {
			return f1-f2 < MIN
		} else {
			return f2-f1 < MIN
		}
	}
	for i, v := range testCases {
		t.Run(fmt.Sprintf("case %d", i+1), func(t *testing.T) {
			res := myPow(v.Input, v.N)
			if !isEqual(res, v.Expect) {
				t.Errorf("input %f, n %d, expect %f, got %f", v.Input, v.N, v.Expect, res)
			}
		})
	}
}
