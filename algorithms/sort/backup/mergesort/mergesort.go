package mergesort

// 归并排序
// 方法：分治法，将待排序数组逐渐二分，每次二分的部分分别排序，最后合并
// 特征：
// 分析：

type MergeSorter struct {
	aux []int
}

func (ms *MergeSorter) Name() string {
	return "merge sort"
}

// Sort sort
func (ms *MergeSorter) Sort(arr []int) {

}

func (ms *MergeSorter) SortIndex(arr []int, lo, hi int) {
	if ms.aux == nil {
		ms.aux = make([]int, len(arr))
	}

	// 递归停止条件，当待排序数组仅有一个元素时
	if hi <= lo {
		return
	}
	// 计算待排序部分的中间位置
	mid := lo + (hi-lo)>>1

	// 分别排序[lo,mid]和[mid+1,hi]的两部分
	ms.SortIndex(arr, lo, mid)
	ms.SortIndex(arr, mid+1, hi)

	// 两部分已经有序，合并两部分内容
	ms.merge(arr, lo, mid, hi)
}

func (ms *MergeSorter) merge(arr []int, lo, mid, hi int) {
	// [lo, mid] 和 [mid+1, hi] 两部分的起始索引
	i := lo
	j := mid + 1

	// 将待merge数据拷贝到辅助数组
	for k := lo; k < hi; k++ {
		ms.aux[k] = arr[k]
	}

	// 分别遍历两部分的所有元素
	for k := lo; k <= hi; k++ {
		if i > mid {
			arr[k] = ms.aux[j]
			j++
		} else if j > hi {
			arr[k] = ms.aux[i]
			i++
		} else if ms.aux[i] < ms.aux[j] {
			arr[k] = ms.aux[i]
			i++
		} else {
			arr[k] = ms.aux[j]
			j++
		}
	}
}
