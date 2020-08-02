package main

import (
	"flag"

	"github.com/bournex/basic_training/algorithms"
	"github.com/bournex/basic_training/interview"
	"github.com/bournex/basic_training/structures"
)

func main() {
	print := flag.Bool("p", false, "print detail")
	flag.Parse()

	algorithms.AlgorithmEntry(*print)
	structures.StructureEntry(*print)
	interview.InterviewEntry(*print)
}
