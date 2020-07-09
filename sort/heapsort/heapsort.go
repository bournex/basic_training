package heapsort

// 堆排序
// 方法：
// 特征：
// 分析：

type HeapSorter struct{}

func (hs *HeapSorter) Name() string {
	return "heap sort"
}

// Sort sort
func (hs *HeapSorter) Sort(arr []int) {

}

func (hs *HeapSorter) SortIndex(arr []int, lo, hi int) {
	l := len(arr)

	for i := l / 2; i >= 1; i-- {
		hs.sink(arr, i, l)
	}

	for l > 1 {
		hs.exch(arr, 1, l)
		l--
		hs.sink(arr, 1, l)
	}
}

func (hs *HeapSorter) sink(arr []int, k, l int) {

	for (k << 1) <= l {

		idx := k << 1
		if idx < l && hs.less(arr, idx, idx+1) {
			idx++
		}

		if hs.less(arr, idx, k) {
			break
		}

		hs.exch(arr, k, idx)

		k = idx
	}
}

func (hs *HeapSorter) exch(arr []int, i, j int) {
	tmp := arr[i-1]
	arr[i-1] = arr[j-1]
	arr[j-1] = tmp
}

func (hs *HeapSorter) less(arr []int, i, j int) bool {
	if arr[i-1] < arr[j-1] {
		return true
	}
	return false
}
