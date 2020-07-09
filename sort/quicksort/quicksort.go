package quicksort

// 快速排序
// 方法：
// 特征：
// 分析：

// QuickSorter quicksorter
type QuickSorter struct{}

func (qs *QuickSorter) Name() string {
	return "quick sort"
}

// Sort sort
func (qs *QuickSorter) Sort(arr []int) {

}

// SortIndex sortindex
func (qs *QuickSorter) SortIndex(arr []int, lo, hi int) {
	// 边界条件检查，当待排序数组中元素数量小于等于1时，则认为输入数组已经有序
	if lo >= hi {
		// 只有一个元素的情况，无需排序
		return
	}

	begin := lo
	end := hi
	mid := arr[lo]
	lo++

	for {
		// 使用<=是因为，等于mid的左边的值，无需移动
		for arr[lo] <= mid {
			if lo == end {
				break
			}
			lo++
		}
		for arr[hi] > mid {
			// 首先递减索引，如果索引减少到了最小值，即begin处，begin是选取的基准值，所以i一�?>=j
			// 退出外循环后，其实begin和j是相等的�?
			hi--
			if hi == begin {
				// begin是mid值所在位置的索引，是不参与交换的，j移动到begin，说明当前轮已经结束
				break
			}
		}
		if lo >= hi {
			break
		}

		tmp := arr[lo]
		arr[lo] = arr[hi]
		arr[hi] = tmp
		lo++
		hi--
	}

	arr[begin] = arr[hi]
	arr[hi] = mid

	qs.SortIndex(arr, begin, hi-1)
	qs.SortIndex(arr, hi+1, end)
}
