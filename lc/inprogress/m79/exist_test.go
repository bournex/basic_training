package m79

import (
	"fmt"
	"testing"
)

func TestExist(t *testing.T) {
	testCases := []struct {
		Input  [][]byte
		Target string
		Expect bool
	}{
		{
			[][]byte{{'C', 'A', 'A'}, {'A', 'A', 'A'}, {'B', 'C', 'D'}},
			"AAB",
			true,
		},
		{
			[][]byte{{'A'}},
			"A",
			true,
		},
		{
			[][]byte{{'B'}},
			"A",
			false,
		},
		{
			[][]byte{{'A', 'B', 'C', 'E'}, {'S', 'F', 'C', 'S'}, {'A', 'D', 'E', 'E'}},
			"ABCCED",
			true,
		},
		{
			[][]byte{{'A', 'B', 'C', 'E'}, {'S', 'F', 'C', 'S'}, {'A', 'D', 'E', 'E'}},
			"SEE",
			true,
		},
		{
			[][]byte{{'A', 'B', 'C', 'E'}, {'S', 'F', 'C', 'S'}, {'A', 'D', 'E', 'E'}},
			"SEEN",
			false,
		},
	}
	for i, v := range testCases {
		t.Run(fmt.Sprintf("case %d", i+1), func(t *testing.T) {
			res := exist(v.Input, v.Target)
			if res != v.Expect {
				t.Errorf("input %+v, target %s, expect %t, got %t", v.Input, v.Target, v.Expect, res)
			}
		})
	}
}
