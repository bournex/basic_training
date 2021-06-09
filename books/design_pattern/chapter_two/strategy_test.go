package chaptertwo

import (
	"fmt"
	"testing"
)

func TestStrategy(t *testing.T) {
	testCases := []struct {
		Input    float64
		Strategy string
		Expect   float64
		Desc     string
	}{
		{
			300,
			StrategyTypeOne,
			300,
			"origin price",
		},
		{
			300,
			StrategyTypeTwo,
			255,
			"15% off",
		},
		{
			300,
			StrategyTypeThree,
			250,
			"50￥ off for each 200￥",
		},
	}
	for i, v := range testCases {
		t.Run(fmt.Sprintf("%d", i+1), func(t *testing.T) {
			cc := NewCashContext(v.Strategy)
			res := cc.AcceptCash(v.Input)
			if res != v.Expect {
				t.Errorf("strategy %s, input %f, expect %f, got %f", v.Strategy, v.Input, v.Expect, res)
			}
		})
	}
}
