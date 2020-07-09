package structures
import (
	"strconv"
	"fmt"

	"github.com/bournex/basic_training/structures/bst"
)

// StructureEntry StructureEntry
func StructureEntry(bprint bool){
	bst := bst.MakeBst()
	bst.Insert(Student{no:4}, Score{math:99})
	bst.Insert(Student{no:1}, Score{math:98})
	bst.Insert(Student{no:5}, Score{math:97})
	bst.Insert(Student{no:2}, Score{math:69})
	bst.Insert(Student{no:3}, Score{math:44})
	fmt.Println(bst.ToString())

	student := Student{no:2}
	v,err := bst.Find(Student{no:2})
	if err == nil {
		score := v.(Score)
		fmt.Println(student.no," score is ",score.math)
	}

	//bst.Delete(student)
	//fmt.Println(bst.ToString())

	student1 := Student{no:4}
	bst.Delete(student1)
	fmt.Println(bst.ToString())
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

func (s Student) ToString()string{
	return strconv.Itoa(s.no)
}