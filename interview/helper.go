package interview

import (
	"fmt"

	"github.com/bournex/basic_training/interview/tencent"
)

func InterviewEntry(p bool) {
	// tencentInterview()
}

func tencentInterview() {
	chess := [][]int{
		{0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 1, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 1, 0},
		{0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 1, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 1, 0, 0, 1, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0},
	}

	if tencent.AttackCart(chess) {
		fmt.Println("there are attacking chessman")
	} else {
		fmt.Println("there are no attacking chessman")
	}

	if tencent.AttackQueen(chess) {
		fmt.Println("there are attacking chessman")
	} else {
		fmt.Println("there are no attacking chessman")
	}
}
