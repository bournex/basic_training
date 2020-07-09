package main

import (
	"flag"

	"github.com/bournex/basic_training/sort"
)

func main() {
	print := flag.Bool("p", false, "print detail")
	flag.Parse()

	sort.AlgSort(*print)
}
