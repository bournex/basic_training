package sort0514

func merge_sort(array []int) {
	n := len(array)
	if n <= 1 {
		return
	}

	assist := make([]int, n)
	var sort func(int, int)
	var merge func(int, int, int)

	sort = func(lo, hi int) {
		if lo >= hi {
			return
		}

		mid := lo + (hi-lo)>>1

		sort(lo, mid)
		sort(mid+1, hi)
		merge(lo, mid, hi)
	}

	merge = func(lo, mid, hi int) {
		// merge [lo,mid],[mid+1,hi]
		for i := lo; i <= hi; i++ {
			assist[i] = array[i]
		}

		idx := lo
		i1, i2 := lo, mid+1
		for i1 <= mid || i2 <= hi {
			if i1 > mid {
				array[idx] = assist[i2]
				i2++
			} else if i2 > hi {
				array[idx] = assist[i1]
				i1++
			} else if assist[i1] < assist[i2] {
				array[idx] = assist[i1]
				i1++
			} else {
				array[idx] = assist[i2]
				i2++
			}
			idx++
		}
	}

	sort(0, n-1)
}

func quick_sort(array []int) {
	n := len(array)
	if n <= 1 {
		return
	}

	var sort func(lo, hi int)
	sort = func(lo, hi int) {
		if lo >= hi {
			return
		}

		mid := lo
		left, right := lo+1, hi
		for {
			for array[left] <= array[mid] {
				if left == hi {
					break
				}
				left++
			}
			for array[right] > array[mid] {
				right--
				if right == lo {
					break
				}
			}

			if left >= right {
				break
			} else {
				array[left], array[right] = array[right], array[left]
				left++
				right--
			}
		}

		array[mid], array[right] = array[right], array[mid]
		sort(lo, right-1)
		sort(right+1, hi)
	}
	sort(0, n-1)
}

func heap_sort(array []int) {
	n := len(array)
	if n <= 1 {
		return
	}

	// 传入为待sink索引和堆数组长度，注意i是从1开始计算的，所以把索引应用到array之前，需要减1
	var sink func(int, int)
	sink = func(i, n int) {
		for i<<1 <= n {
			max := i
			l, r := i<<1, i<<1+1

			// 左子节点
			if l-1 < n && array[l-1] > array[max-1] {
				max = l
			}
			// 右子节点
			if r-1 < n && array[r-1] > array[max-1] {
				max = r
			}

			if max == i {
				// 当前sink到位了
				return
			}

			array[max-1], array[i-1] = array[i-1], array[max-1]
			i = max
		}
	}

	for i := n / 2; i >= 1; i-- {
		sink(i, n)
	}

	// 令i=n-1，将极大值(array[0])与j位置值交换，将array[0]下沉，i--，再递归对长度为i的序列做同样的事情
	for n > 1 {
		array[n-1], array[0] = array[0], array[n-1]
		n--
		sink(1, n)
	}
}
