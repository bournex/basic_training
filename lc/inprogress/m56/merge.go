package m56

import "sort"

type Interval [][]int

func (interval Interval) Len() int           { return len(([][]int)(interval)) }
func (interval Interval) Swap(i, j int)      { interval[i], interval[j] = interval[j], interval[i] }
func (interval Interval) Less(i, j int) bool { return interval[i][0] < interval[j][0] }

func merge(intervals [][]int) [][]int {
	n := len(intervals)
	if n <= 1 {
		return intervals
	}

	result := make([][]int, 0, n)
	sort.Sort(Interval(intervals))

	i, j := 0, 1
	left, right := intervals[i][0], intervals[i][1]
	for i < n && j < n {
		if left <= intervals[j][0] && intervals[j][0] <= right {
			// j的左侧落在i的区间内
			if intervals[j][1] > right {
				right = intervals[j][1]
			}
			j++
		} else {
			// 两区间没有交集
			result = append(result, []int{left, right})
			i = j
			left, right = intervals[i][0], intervals[i][1]
			j++
		}
	}
	return append(result, []int{left, right})
}
