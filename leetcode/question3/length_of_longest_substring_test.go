package question3

import "testing"

func TestLengthOfLongestSubstring(t *testing.T) {
	sample := "abcabcbb"
	expect := 3

	res := lengthOfLongestSubstring(sample)
	if res != expect {
		t.Errorf("expect %d, got %d", expect, res)
	}
}
