package q1

import "testing"

func TestTwoSum(t *testing.T) {
	nums := []int{2, 7, 11, 15}
	target := 9
	res := twoSum(nums, target)

	expect := []int{0, 1}

	if len(expect) != len(res) {
		t.Errorf("length not expect")
	}

	for i := 0; i < len(expect); i++ {
		if expect[i] != res[i] {
			t.Errorf("result not expect at index[%d], expect %d, get %d", i, expect[i], res[i])
		}
	}
}
