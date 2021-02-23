package working

import (
	"fmt"
	"time"
)

func WorkingEntry(print bool) {
	testMaxTimeDuration()
}

func testMaxTimeDuration() {
	now := time.Now()

	s1 := now
	e1 := now.Add(5 * time.Minute)

	s2 := now.Add(3 * time.Minute)
	e2 := now.Add(10 * time.Minute)

	s3 := now.Add(4 * time.Minute)
	e3 := now.Add(8 * time.Minute)

	s4 := now.Add(9 * time.Minute)
	e4 := now.Add(15 * time.Minute)

	start := []time.Time{s1, s2, s3, s4}
	end := []time.Time{e1, e3, e2, e4}

	for _, v := range start {
		fmt.Println(v.Format("2006-01-02 15:04:05"))
	}

	for _, v := range end {
		fmt.Println(v.Format("2006-01-02 15:04:05"))
	}

	m, s, e := MaxTimeDuration(start, end)
	fmt.Println(m, s.Format("2006-01-02 15:04:05"), e.Format("2006-01-02 15:04:05"))
}
