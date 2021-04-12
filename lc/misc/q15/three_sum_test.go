package q15

import "testing"

func TestThreeSum(t *testing.T) {
	// input := [][]int{{-1, 0, 1, 2, -1, -4}, {}, {0}, {0, 0, 0}, {0, 0, 0, 0}, {-1, 0, 1, 2, -1, -4}}
	// expect := [][][]int{{{-1, -1, 2}, {-1, 0, 1}}, {}, {}, {{0, 0, 0}}, {{0, 0, 0}}, {{-1, -1, 2}, {-1, 0, 1}}}
	input := [][]int{{-1, 0, 1, 2, -1, -4}}
	expect := [][][]int{{{-1, -1, 2}, {-1, 0, 1}}}

	for i, v := range input {
		res := threeSum(v)
		e := expect[i]
		if len(res) != len(e) {
			t.Errorf("input %+v, expect %+v, got %+v", v, e, res)
		}
		t.Logf("expect %+v, got %+v", e, res)
	}
}
