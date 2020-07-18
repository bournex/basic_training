package structures

import (
	"fmt"
	"math/rand"
	"strconv"
	"time"

	"github.com/bournex/basic_training/structures/bst"
	"github.com/bournex/basic_training/structures/heap"
	"github.com/bournex/basic_training/structures/skiplist"
	"github.com/bournex/basic_training/structures/str/sst"
	"github.com/bournex/basic_training/structures/str/tst"
)

// 用于测试的数据类型定义
type Student struct {
	// 学号
	no int
}

type Score struct {
	math     int
	physical int
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

// StructureEntry StructureEntry
func StructureEntry(bprint bool) {
	// testBst()
	// testSkipList()
	// testHeap()
	// testSst()
	testTst()
}

func testTst() {
	t := tst.MakeTst()
	t.Insert("by", Score{math: 99})
	t.Insert("bye", Score{math: 70})
	t.Insert("byte", Score{math: 94})
	t.Insert("be", Score{math: 94})

	name := "bye"

	findAndPrint := func(name string) {
		val := t.Find(name)
		if val == nil {
			fmt.Printf("%s's math score not exist\n", name)
		} else {
			score := val.(Score)
			fmt.Printf("%s's math score is %d\n", name, score.math)
		}
	}

	findAndPrint(name)
	t.Delete(name)
	findAndPrint(name)
	t.Insert(name, Score{math: 89})
	findAndPrint(name)
}

func testSst() {
	s := sst.MakeSst()
	s.Insert("by", Score{math: 99})
	s.Insert("bye", Score{math: 70})
	s.Insert("byte", Score{math: 94})
	s.Insert("be", Score{math: 100})

	name := "bye"

	findAndPrint := func(name string) {
		val := s.Find(name)
		if val == nil {
			fmt.Println(name, "'s math score not exist")
		} else {
			score := val.(Score)
			fmt.Println(name, "'s math score is", score.math)
		}
	}

	findAndPrint(name)
	s.Delete(name)
	findAndPrint(name)
	s.Insert(name, Score{math: 89})
	findAndPrint(name)
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
