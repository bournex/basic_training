package m57

// 思路
//	二分查找左边界和右边界，找到后合并区间
func insert(intervals [][]int, newInterval []int) [][]int {
	n := len(intervals)
	if n == 0 {
		return [][]int{newInterval}
	}

	lbound, rbound := 0, 0
	left, right := 0, n-1
	for left < right {
		mid := left + (right-left)>>1
		if intervals[mid][0] < newInterval[0] {
			left = mid
			lbound = mid
		} else if intervals[mid][0] > newInterval[0] {
			right = mid - 1
		} else {
			lbound = mid
			break
		}
	}

	left, right = 0, n-1
	for left < right {
		mid := left + (right-left)>>1
		if intervals[mid][0] > newInterval[1] {
			right = mid
			rbound = mid
		} else if intervals[mid][0] < newInterval[1] {
			left = mid + 1
		} else {
			rbound = mid
			break
		}
	}

	result := make([][]int, 0, n)
	if lbound > 1 {
		result = append(result, intervals[0:lbound-1]...)
	}

	if intervals[lbound][1] > newInterval[0] && intervals[rbound][0] < newInterval[1] {
		result = append(result, []int{intervals[lbound][0], intervals[rbound][1]})
	} else {
		lbound = intervals[lbound][0]
	}

	if intervals[rbound][0] <= newInterval[1] {
		rbound = intervals[rbound][1]
	}

	return result
}
