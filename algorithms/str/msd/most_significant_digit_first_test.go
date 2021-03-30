package msd

import "testing"

func TestMostSignificantDigitFirst(t *testing.T) {
	corpus := []struct {
		Input  []string
		Expect []string
	}{
		{
			[]string{
				"int",
				"infomation",
				"in",
				"info",
			},
			[]string{
				"in",
				"info",
				"infomation",
				"int",
			},
		},
		{
			[]string{
				"GitHub", "collects", "information", "directly", "from",
				"you", "for", "your", "registration", "payment",
				"transactions", "and", "user", "profile", "We",
				"also", "automatically", "collect", "from", "you",
				"your", "usage", "information", "cookies", "and",
				"similar", "technologies", "and", "device", "information",
				"subject", "where", "necessary", "to", "your",
				"consent", "GitHub", "may", "also", "collect",
				"User", "Personal", "Information", "from", "third",
				"parties", "We", "only", "collect", "the",
				"minimum", "amount", "of", "personal", "information",
				"necessary", "from", "you", "unless", "you",
				"choose", "to", "provide", "more",
			},
			[]string{
				"GitHub", "GitHub", "Information", "Personal", "User",
				"We", "We", "also", "also", "amount",
				"and", "and", "and", "automatically", "choose",
				"collect", "collect", "collect", "collects", "consent",
				"cookies", "device", "directly", "for", "from",
				"from", "from", "from", "information", "information",
				"information", "information", "may", "minimum", "more",
				"necessary", "necessary", "of", "only", "parties",
				"payment", "personal", "profile", "provide", "registration",
				"similar", "subject", "technologies", "the", "third",
				"to", "to", "transactions", "unless", "usage",
				"user", "where", "you", "you", "you",
				"you", "your", "your", "your",
			},
		},
	}

	for _, v := range corpus {
		cp := make([]string, len(v.Input))
		copy(cp, v.Input)
		Msd(v.Input)
		for i, v1 := range v.Input {
			if v.Expect[i] != v1 {
				t.Errorf("input %+v, expect %+v, got %+v", cp, v.Expect, v.Input)
			}
		}
	}
}
