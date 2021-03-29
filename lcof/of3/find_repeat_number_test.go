package of3

import "testing"

func TestFindRepeatNumber(t *testing.T) {
	sample := [][]int{{2, 3, 1, 0, 2, 5, 3}, {1, 2, 3, 4, 5, 6, 7, 8, 7}, {}, {3, 1, 2, 3}}
	expect := []int{2, 7, -1, 3}
	for i, v := range sample {
		res := findRepeatNumber(v)
		if res != expect[i] {
			t.Errorf("expect %d, got %d", expect[i], res)
		}
	}
}

func TestFindRepeatNumber1(t *testing.T) {
	sample := [][]int{{2, 3, 1, 0, 2, 5, 3}, {1, 2, 3, 4, 5, 6, 7, 8, 7}, {}, {3, 1, 2, 3}}
	expect := []int{2, 7, -1, 3}
	for i, v := range sample {
		res := findRepeatNumber1(v)
		if res != expect[i] {
			t.Errorf("expect %d, got %d", expect[i], res)
		}
	}
}
