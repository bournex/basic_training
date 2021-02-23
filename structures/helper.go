package structures

import (
	"fmt"
	"math/rand"
	"strconv"
	"strings"
	"time"

	"github.com/bournex/basic_training/structures/base/heap"
	"github.com/bournex/basic_training/structures/base/stack"
	"github.com/bournex/basic_training/structures/cache"
	"github.com/bournex/basic_training/structures/graph"
	"github.com/bournex/basic_training/structures/skiplist"
	"github.com/bournex/basic_training/structures/tree/bst"
	"github.com/bournex/basic_training/structures/trie/sst"
	"github.com/bournex/basic_training/structures/trie/tst"
)

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
	testHeap()
	// testSst()
	// testTst()
	// testGraph()
	// testStack()
	// testFifo()
	// testLru()
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

	g.Init(s)

	// DFS判定连通性
	t1 := func(v1, v2 string, b bool) {
		if b {
			fmt.Println(v1, "and", v2, "is connected")
		} else {
			fmt.Println(v1, "and", v2, "is not connected")
		}
	}
	t1("E", "H", g.Connected("E", "H"))
	t1("E", "W", g.Connected("E", "W"))

	// BFS查找最短路径
	t2 := func(v1, v2 string, p []string) {
		if p != nil {
			path := ""
			for _, v := range p {
				path += v
				path += " -> "
			}
			path = strings.TrimSuffix(path, " -> ")
			fmt.Printf("shortest path from %s to %s is %s\n", v1, v2, path)
		} else {
			fmt.Printf("there is no path from %s to %s\n", v1, v2)
		}
	}

	t2("E", "G", g.ShortestPath("E", "G"))
	t2("H", "E", g.ShortestPath("H", "E"))
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
	h := heap.MakeHeap(4)
	h.Insert(student{no: 4})
	h.Insert(student{no: 1})
	h.Insert(student{no: 5})
	h.Insert(student{no: 2})
	h.Insert(student{no: 3})
	h.Insert(student{no: 11})
	h.Insert(student{no: 25})
	h.Insert(student{no: 13})
	h.Insert(student{no: 9})
	h.Insert(student{no: 8})

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

func testStack() {
	stackWithoutLimit := stack.MakeStack(0)
	for i := 0; i < 16; i++ {
		err := stackWithoutLimit.Push(i)
		if err != nil {
			fmt.Println(err.Error())
		} else {
			fmt.Println(i)
		}
	}
	fmt.Printf("stack info %+v\n", stackWithoutLimit)
	for i := 0; i < 16; i++ {
		v, err := stackWithoutLimit.Pop()
		if err != nil {
			fmt.Println(err.Error())
		} else {
			fmt.Println(v.(int))
		}
	}
	fmt.Printf("stack info %+v\n", stackWithoutLimit)

	stackWithLimit := stack.MakeStack(3)

	for i := 0; i < 4; i++ {
		err := stackWithLimit.Push(i)
		if err != nil {
			fmt.Println(err.Error())
		} else {
			fmt.Println(i)
		}
	}
	fmt.Printf("stack info %+v\n", stackWithLimit)
	for i := 0; i < 4; i++ {
		v, err := stackWithLimit.Pop()
		if err != nil {
			fmt.Println(err.Error())
		} else {
			fmt.Println(v.(int))
		}
	}
	fmt.Printf("stack info %+v\n", stackWithLimit)
}

func testFifo() {
	f := cache.MakeFifo(3)
	f.Set("a", 1)
	f.Set("b", 2)
	f.Set("c", 3)
	// 期望c -> 3、b -> 2、a -> 1
	fmt.Println(f.ToString())
	f.Set("b", 4)
	// 期望c -> 3、b -> 4、a -> 1
	fmt.Println(f.ToString())
	f.Set("d", 4)
	// 期望d -> 4、c -> 3、b -> 4
	fmt.Println(f.ToString())
	f.Set("e", 5)
	// 期望e -> 5、d -> 4、c -> 3
	fmt.Println(f.ToString())
}

func testLru() {
	c := cache.MakeLru(3)
	c.Set("a", 1)
	c.Set("b", 2)
	c.Set("c", 3)
	// 期望c -> 3、b -> 2、a - > 1
	fmt.Println(c.ToString())
	c.Set("b", 4)
	// 期望c -> 3、b -> 4、a - > 1
	fmt.Println(c.ToString())
	c.Get("a")
	// 期望a - > 1、c -> 3、b -> 4
	fmt.Println(c.ToString())
	c.Set("d", 4)
	// 期望d -> 4、a - > 1、c -> 3
	fmt.Println(c.ToString())
	c.Del("a")
	// 期望d -> 4、c -> 3
	fmt.Println(c.ToString())
}
