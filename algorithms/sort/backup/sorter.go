package backup

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/bournex/basic_training/algorithms/sort/backup/heapsort"
	"github.com/bournex/basic_training/algorithms/sort/backup/insertsort"
	"github.com/bournex/basic_training/algorithms/sort/backup/mergesort"
	"github.com/bournex/basic_training/algorithms/sort/backup/quicksort"
	"github.com/bournex/basic_training/algorithms/sort/backup/selectsort"
	"github.com/bournex/basic_training/algorithms/sort/backup/shellsort"
)

// helper functions
func init() {
	rand.Seed(time.Now().Unix())
}

func checkSort(arr []int) bool {
	var last int
	for _, v := range arr {
		if v < last {
			return false
		}
		last = v
	}
	return true
}

func randomSource(n int) []int {
	target := make([]int, n)
	for i := 0; i < n; i++ {
		target[i] = rand.Intn(n)
	}
	return target
}

func randomDescSource(n int) []int {
	target := make([]int, n)
	for i := n; i > 0; i-- {
		target[n-i] = i
	}
	return target
}

func randomAscSource(n int) []int {
	target := make([]int, n)
	for i := 0; i < n; i++ {
		target[i] = i
	}
	return target
}

// Sorter sorter interface
type Sorter interface {
	Sort(target []int)
	SortIndex(target []int, lo, hi int)
	Name() string
}

func runSort(s Sorter, source []int, needPrint bool) {
	fmt.Printf("| running %s |------------------------------------\n", s.Name())
	if needPrint {
		fmt.Println("[origin array] ", source)
	}

	target := make([]int, len(source))
	copy(target, source)

	t1 := time.Now()
	s.SortIndex(target, 0, len(target)-1)
	d := time.Since(t1)

	if needPrint {
		fmt.Println("[after sort]", target)
	}

	if checkSort(target) {
		fmt.Printf("%s cost %f seconds\n", s.Name(), d.Seconds())
	} else {
		fmt.Printf("%s failed\n", s.Name())
	}
}

// AlgSort algsort
func AlgSort(bPrint bool) {
	doPrint := bPrint
	source := randomAscSource(10000)

	s1 := new(selectsort.SelectSorter)
	runSort(s1, source, doPrint)

	s2 := new(insertsort.InsertSorter)
	runSort(s2, source, doPrint)

	s3 := new(shellsort.ShellSorter)
	runSort(s3, source, doPrint)

	s4 := new(quicksort.QuickSorter)
	runSort(s4, source, doPrint)

	s5 := new(mergesort.MergeSorter)
	runSort(s5, source, doPrint)

	s6 := new(heapsort.HeapSorter)
	runSort(s6, source, doPrint)
}
