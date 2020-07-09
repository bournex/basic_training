package main

import (
	"flag"

	"github.com/bournex/basic_training/structures"
)

func main() {
	print := flag.Bool("p", false, "print detail")
	flag.Parse()

	// sort.AlgSort(*print)

	structures.StructureEntry(*print)
}
