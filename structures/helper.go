package structures

import (
	"fmt"
	"math/rand"
	"strconv"
	"time"

	"github.com/bournex/basic_training/structures/bst"
	"github.com/bournex/basic_training/structures/heap"
	"github.com/bournex/basic_training/structures/skiplist"
)

// StructureEntry StructureEntry
func StructureEntry(bprint bool) {
	// testBst()
	// testSkipList()
	// testHeap()
}

func testBst() {
	bst := bst.MakeBst()
	bst.Insert(Student{no: 4}, Score{math: 99})
	bst.Insert(Student{no: 1}, Score{math: 98})
	bst.Insert(Student{no: 5}, Score{math: 97})
	bst.Insert(Student{no: 2}, Score{math: 69})
	bst.Insert(Student{no: 3}, Score{math: 44})
	fmt.Println(bst.ToString())

	student := Student{no: 2}
	v := bst.Find(Student{no: 2})
	if v != nil {
		score := v.(Score)
		fmt.Println(student.no, " score is ", score.math)
	}
}

func testSkipList() {
	rand.Seed(time.Now().Unix())
	sl := skiplist.MakeSkipList(4)
	for i := 0; i < 30; i++ {
		st := Student{no: i + 1}
		sc := Score{math: rand.Intn(100)}
		sl.Insert(st, sc)
	}
	fmt.Printf(sl.ToString())

	sl.Delete(Student{no: 1})
	fmt.Printf(sl.ToString())

	sl.Delete(Student{no: 20})
	fmt.Printf(sl.ToString())

	sl.Delete(Student{no: 10})
	fmt.Printf(sl.ToString())
}

func testHeap() {
	h := heap.MakeHeap(10)
	h.Insert(Student{no: 4})
	h.Insert(Student{no: 1})
	h.Insert(Student{no: 5})
	h.Insert(Student{no: 2})
	h.Insert(Student{no: 3})

	fmt.Println(h.ToString())

	for i := 0; i < 6; i++ {
		fmt.Println("round ", i+1, " extra max-----------------")
		v, err := h.ExtraMax()
		if err != nil {
			fmt.Println("error occur,", err.Error())
		} else {
			fmt.Printf("max in heap [%s] extra\n", v.ToString())
		}
		fmt.Println(h.ToString())
	}
}

type Student struct {
	// 学号
	no int
}

type Score struct {
	math int
}

func (s Student) Compare(m interface{}) int {
	mp := m.(Student)
	if s.no > mp.no {
		return 1
	} else if s.no < mp.no {
		return -1
	}
	return 0
}

func (s Student) ToString() string {
	return strconv.Itoa(s.no)
}
