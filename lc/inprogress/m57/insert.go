package m57

// 思路
//	二分查找左边界和右边界，找到后合并区间
//	上面是最开始的思路，看过了答案，考虑复杂了，二分查找并不能提速多少，主要的耗时还在于拷贝
//

func insert(intervals [][]int, newInterval []int) [][]int {
	n := len(intervals)
	if n == 0 {
		return [][]int{newInterval}
	}

	max := func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}
	min := func(a, b int) int {
		if a > b {
			return b
		}
		return a
	}

	mergeIndex := -1
	result := make([][]int, 0, n+1) // 最多可能有n+1个区间，所以预分配好空间，避免数据拷贝
	for i, v := range intervals {
		if v[0] > newInterval[1] {
			// newInterval在v区间的前面
			if mergeIndex == -1 {
				result = append(result, newInterval)
				mergeIndex = i
			}
			result = append(result, intervals[i:]...)
			break
		} else if v[1] < newInterval[0] {
			// newInterval在v区间的后面
			result = append(result, v)
		} else {
			// newInterval与v有交集，合并两个区间
			if mergeIndex == -1 {
				left := min(newInterval[0], v[0])
				right := max(newInterval[1], v[1])
				result = append(result, []int{left, right})
				mergeIndex = i
			} else {
				left := min(min(newInterval[0], v[0]), intervals[mergeIndex][0])
				right := max(max(newInterval[1], v[1]), intervals[mergeIndex][1])
				result[mergeIndex][0], result[mergeIndex][1] = left, right
			}
		}
	}
	if mergeIndex == -1 {
		result = append(result, newInterval)
	}
	return result
}

func insert1(intervals [][]int, newInterval []int) [][]int {
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
	} else if intervals[lbound][1] < newInterval[0] {
		lbound = intervals[lbound][0]
	}

	if intervals[rbound][0] <= newInterval[1] {
		rbound = intervals[rbound][1]
	}

	return result
}
