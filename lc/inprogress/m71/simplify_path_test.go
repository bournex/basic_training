package m71

import (
	"fmt"
	"testing"
)

func TestSimplifyPath(t *testing.T) {
	testCases := []struct {
		Input  string
		Expect string
	}{
		{
			"/a/b/c/../../..",
			"/",
		},
		{
			"/a//b////c/d//././/..",
			"/a/b/c",
		},
		{
			"/abc/./def",
			"/abc/def",
		},
		{
			"/abc/...",
			"/abc/...",
		},
		{
			"/../",
			"/",
		},
		{
			"/......",
			"/......",
		},
		{
			"/",
			"/",
		},
		{
			"/a/./b/../../c/",
			"/c",
		},
		{
			"/home/",
			"/home",
		},
		{
			"/home//foo/",
			"/home/foo",
		},
	}
	for i, v := range testCases {
		t.Run(fmt.Sprintf("%d", i+1), func(t *testing.T) {
			res := simplifyPath(v.Input)
			if res != v.Expect {
				t.Errorf("input %s, expect %s, got %s", v.Input, v.Expect, res)
			}
		})
	}
}
