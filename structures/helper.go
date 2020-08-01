package structures

import (
	"fmt"
	"math/rand"
	"strconv"
	"time"

	"github.com/bournex/basic_training/structures/graph"
	"github.com/bournex/basic_training/structures/heap"
	"github.com/bournex/basic_training/structures/skiplist"
	"github.com/bournex/basic_training/structures/tree/bst"
	"github.com/bournex/basic_training/structures/trie/sst"
	"github.com/bournex/basic_training/structures/trie/tst"
	/*
		"github.com/bournex/basic_training/structures/bst"
		"github.com/bournex/basic_training/structures/heap"
		"github.com/bournex/basic_training/structures/skiplist"
		"github.com/bournex/basic_training/structures/str/sst"
		"github.com/bournex/basic_training/structures/str/tst"
	*/)

// student 用于测试的数据类型定义
type student struct {
	// 学号
	no int
}

// score 学生分数
type score struct {
	math     int
	physical int
}

// Compare 比较学生学号
func (s student) Compare(m interface{}) int {
	mp := m.(student)
	if s.no > mp.no {
		return 1
	} else if s.no < mp.no {
		return -1
	}
	return 0
}

func (s student) ToString() string {
	return strconv.Itoa(s.no)
}

// StructureEntry StructureEntry
func StructureEntry(bprint bool) {
	// testBst()
	// testSkipList()
	// testHeap()
	// testSst()
	// testTst()
	testGraph()
}

func testGraph() {
	g := graph.MakeGraph()

	dat := []string{
		"A0B",
		"A0C",
		"A0D",
		"B0E",
		"B0F",
		"C0G",
		"C0D",
		"D0H",
		"D0I",
		"D0J",
	}

	//       C - G
	//      / \
	// J - D - A - B - E
	//    / \       \
	//   I   H       F

	s := []*graph.Relation{}
	for _, v := range dat {
		e, _ := strconv.Atoi(v[1 : len(v)-2])
		r := &graph.Relation{V1: string(v[0]), V2: string(v[len(v)-1]), E: e}
		s = append(s, r)
	}

	t := func(v1, v2 string, b bool) {
		if b {
			fmt.Println(v1, "and", v2, "is connected")
		} else {
			fmt.Println(v1, "and", v2, "is not connected")
		}
	}

	g.Init(s)
	t("E", "H", g.Connected("E", "H"))
	t("E", "W", g.Connected("E", "W"))
}

func testTst() {
	t := tst.MakeTst()
	t.Insert("by", score{math: 99})
	t.Insert("bye", score{math: 70})
	t.Insert("byte", score{math: 94})
	t.Insert("be", score{math: 94})

	name := "bye"

	findAndPrint := func(name string) {
		val := t.Find(name)
		if val == nil {
			fmt.Printf("%s's math score not exist\n", name)
		} else {
			score := val.(score)
			fmt.Printf("%s's math score is %d\n", name, score.math)
		}
	}

	findAndPrint(name)
	t.Delete(name)
	findAndPrint(name)
	t.Insert(name, score{math: 89})
	findAndPrint(name)
}

func testSst() {
	s := sst.MakeSst()
	s.Insert("by", score{math: 99})
	s.Insert("bye", score{math: 70})
	s.Insert("byte", score{math: 94})
	s.Insert("be", score{math: 100})

	name := "bye"

	findAndPrint := func(name string) {
		val := s.Find(name)
		if val == nil {
			fmt.Println(name, "'s math score not exist")
		} else {
			score := val.(score)
			fmt.Println(name, "'s math score is", score.math)
		}
	}

	findAndPrint(name)
	s.Delete(name)
	findAndPrint(name)
	s.Insert(name, score{math: 89})
	findAndPrint(name)
}

func testBst() {
	bst := bst.MakeBst()
	bst.Insert(student{no: 4}, score{math: 99})
	bst.Insert(student{no: 1}, score{math: 98})
	bst.Insert(student{no: 5}, score{math: 97})
	bst.Insert(student{no: 2}, score{math: 69})
	bst.Insert(student{no: 3}, score{math: 44})
	fmt.Println(bst.ToString())

	s := student{no: 2}
	v := bst.Find(student{no: 2})
	if v != nil {
		score := v.(score)
		fmt.Println(s.no, " score is ", score.math)
	}
}

func testSkipList() {
	rand.Seed(time.Now().Unix())
	sl := skiplist.MakeSkipList(4)
	for i := 0; i < 30; i++ {
		st := student{no: i + 1}
		sc := score{math: rand.Intn(100)}
		sl.Insert(st, sc)
	}
	fmt.Printf(sl.ToString())

	sl.Delete(student{no: 1})
	fmt.Printf(sl.ToString())

	sl.Delete(student{no: 20})
	fmt.Printf(sl.ToString())

	sl.Delete(student{no: 10})
	fmt.Printf(sl.ToString())
}

func testHeap() {
	h := heap.MakeHeap(10)
	h.Insert(student{no: 4})
	h.Insert(student{no: 1})
	h.Insert(student{no: 5})
	h.Insert(student{no: 2})
	h.Insert(student{no: 3})

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
