package algorithms

import (
	"fmt"

	"github.com/bournex/basic_training/algorithms/sort"
	"github.com/bournex/basic_training/algorithms/str"
)

// AlgorithmEntry AlgorithmEntry
func AlgorithmEntry(bprint bool) {
	// testSort(bprint)
	// testLsd()
	// testMsd()
	// testQuick3()
}

func testSort(p bool) {
	sort.AlgSort(p)
}

func testMsd() {
	str.Msd(miscLengthStringArray)
	for _, s := range miscLengthStringArray {
		fmt.Println(s)
	}
}

func testLsd() {
	str.Lsd(sameLengthStringSlice, len(sameLengthStringSlice[0]))
	for _, s := range sameLengthStringSlice {
		fmt.Println(s)
	}
}

func testQuick3() {
	str.Quick3(miscLengthStringArray)
	for _, s := range miscLengthStringArray {
		fmt.Println(s)
	}
}

var (
	shortStringSlice = []string{
		"int",
		"infomation",
		"in",
		"info",
	}
	sameLengthStringSlice = []string{
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
	miscLengthStringArray = []string{
		"GitHub",
		"collects",
		"information",
		"directly",
		"from",
		"you",
		"for",
		"your",
		"registration",
		"payment",
		"transactions",
		"and",
		"user",
		"profile",
		"We",
		"also",
		"automatically",
		"collect",
		"from",
		"you",
		"your",
		"usage",
		"information",
		"cookies",
		"and",
		"similar",
		"technologies",
		"and",
		"device",
		"information",
		"subject",
		"where",
		"necessary",
		"to",
		"your",
		"consent",
		"GitHub",
		"may",
		"also",
		"collect",
		"User",
		"Personal",
		"Information",
		"from",
		"third",
		"parties",
		"We",
		"only",
		"collect",
		"the",
		"minimum",
		"amount",
		"of",
		"personal",
		"information",
		"necessary",
		"from",
		"you",
		"unless",
		"you",
		"choose",
		"to",
		"provide",
		"more",
	}
)
