package algorithms

import (
	"fmt"

	"github.com/bournex/basic_training/algorithms/str"
)

// AlgorithmEntry AlgorithmEntry
func AlgorithmEntry(bprint bool) {
	testLsd()
}

func testLsd() {
	ss := []string{
		"abcdc",
		"wtomm",
		"uncle",
		"hello",
		"world",
		"abcde",
		"abcce",
		"abcdf",
		"which",
		"heavy",
		"sunny",
		"happy",
	}

	str.Lsd(ss)
	for _, s := range ss {
		fmt.Println(s)
	}
}
