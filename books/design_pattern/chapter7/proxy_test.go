package chapter7

import (
	"fmt"
	"testing"
)

func TestProxy(t *testing.T) {
	testCases := []struct {
		Source string
		Target string
		Expect string
	}{
		{
			"tom",
			"kitty",
			"tom send flower to kitty\ntom send chocolate to kitty\ntom send dolls to kitty\n",
		},
	}
	for i, v := range testCases {
		t.Run(fmt.Sprintf("%d", i+1), func(t *testing.T) {
			var result string
			target := NewTarget(v.Target)
			proxy := NewProxy(v.Source, target)
			result += proxy.GiveFlower()
			result += proxy.GiveChocolate()
			result += proxy.GiveDolls()
			if result != v.Expect {
				t.Errorf("source %s, target %s, expect %s, got %s", v.Source, v.Target, v.Expect, result)
			}
		})
	}
}
