package h37

import (
	"fmt"
	"testing"
)

func TestSudoku(t *testing.T) {
	testCases := []struct {
		Input [][]byte
		Desc  string
	}{
		{
			[][]byte{{'5', '3', '.', '.', '7', '.', '.', '.', '.'}, {'6', '.', '.', '1', '9', '5', '.', '.', '.'}, {'.', '9', '8', '.', '.', '.', '.', '6', '.'}, {'8', '.', '.', '.', '6', '.', '.', '.', '3'}, {'4', '.', '.', '8', '.', '3', '.', '.', '1'}, {'7', '.', '.', '.', '2', '.', '.', '.', '6'}, {'.', '6', '.', '.', '.', '.', '2', '8', '.'}, {'.', '.', '.', '4', '1', '9', '.', '.', '5'}, {'.', '.', '.', '.', '8', '.', '.', '7', '9'}},
			"",
		},
	}
	for i, v := range testCases {
		t.Run(fmt.Sprintf("%d", i+1), func(t *testing.T) {
			input := make([][]byte, len(v.Input))
			for j := 0; j < len(v.Input); j++ {
				input[j] = make([]byte, len(v.Input[j]))
				copy(input[j], v.Input[j])
			}
			solveSudoku(input)

			for j := 0; j < len(input); j++ {
				hash := make(map[byte]bool, 9)
				for k := 0; k < len(input[j]); k++ {
					t.Logf("%d ", input[j][k]-'0')
					if _, exist := hash[input[j][k]]; exist {
						t.Errorf("input %+v, got %+v", v.Input, input)
						return
					}
				}
				t.Logf("\n")
			}
		})
	}
}
