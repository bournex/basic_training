package of5

import "testing"

func TestReplaceSpace(t *testing.T) {
	examples := []struct {
		Input  string
		Expect string
	}{
		{
			"",
			"",
		},
		{
			" ",
			"%20",
		},
		{
			" space ship",
			"%20space%20ship",
		},
		{
			"space  ship",
			"space%20%20ship",
		},
		{
			"space ship",
			"space%20ship",
		},
		{
			"space ship ",
			"space%20ship%20",
		},
	}

	for _, v := range examples {
		res := replaceSpace(v.Input)
		if res != v.Expect {
			t.Errorf("input %s, expect %s, got %s", v.Input, v.Expect, res)
		}
	}
}
