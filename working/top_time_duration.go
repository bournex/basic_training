package working

import (
	"time"

	"github.com/bournex/basic_training/structures/base/stack"
)

// 输入已排序的开始时间与结束时间列表
// 输出最大重叠时间段和重叠次数
func MaxTimeDuration(start, end []time.Time) (int, time.Time, time.Time) {
	var (
		max int
		ts  time.Time
		te  time.Time
		// tmps  time.Time
		stack = stack.MakeStack(100)
	)

	closeTime := func(e time.Time) {
		l := stack.Length()
		v, _ := stack.Pop()
		s := v.(time.Time)

		if l > max {
			// 如果当前close的时间段重叠层级高于之前保存的，则更新ts,te
			max = l
			ts, te = s, e
		} else if l == max {
			// 如果当前close的时间段重叠层级等于当前保存的
			// 判断时间长短，选择时间较长的那个
			if !ts.IsZero() && !te.IsZero() {
				if te.Sub(ts) < e.Sub(s) {
					ts, te = s, e
				}
			} else {
				ts, te = s, e
			}
		}
	}

	for len(start) > 0 {
		s := start[0]
		e := end[0] // start不为空则end不可能为空

		if s.Before(e) {
			start = start[1:]
			stack.Push(s) // 压栈
		} else if s.After(e) {
			closeTime(e) // 弹栈
			end = end[1:]
		}
	}

	closeTime(end[0])

	return max, ts, te
}
