package chapter6

import (
	"fmt"
	"testing"
)

func TestDecorator(t *testing.T) {
	testCases := []struct {
		Name   string
		Expect string
		Desc   string
	}{
		{
			"tom",
			"tom wearing shorts t-shirt",
			"",
		},
	}
	for i, v := range testCases {
		t.Run(fmt.Sprintf("%d", i+1), func(t *testing.T) {
			var p AbstractPerson = &Person{v.Name}

			// 先穿短裤，再穿t-shirt
			shorts := Shorts{Decorator: Decorator{P: p}}
			tshirt := TShirt{Decorator: Decorator{P: shorts}}

			res := tshirt.Show()
			if res != v.Expect {
				t.Errorf("input %s, expect %s, got %s", v.Name, v.Expect, res)
			}
		})
	}
}
