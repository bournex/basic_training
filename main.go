package main

import (
	"flag"

	//"github.com/bournex/basic_training/interview"
	"github.com/bournex/basic_training/structures"
	"github.com/bournex/basic_training/working"
)

func main() {
	print := flag.Bool("p", false, "print detail")
	flag.Parse()

	// algorithms.AlgorithmEntry(*print)
	structures.StructureEntry(*print)
	//interview.InterviewEntry(*print)
	working.WorkingEntry(*print)
}
