package sort0510

// 盲写排序

// 插入排序
func insert_sort(array []int) {
	n := len(array)
	for i := 1; i < n; i++ {
		for j := i; j > 0; j-- {
			if array[j] < array[j-1] {
				array[j], array[j-1] = array[j-1], array[j]
			} else {
				break
			}
		}
	}
}

// 选择排序
func select_sort(array []int) {
	n := len(array)
	for i := 0; i < n; i++ {
		min := i
		for j := i + 1; j < n; j++ {
			if array[min] > array[j] {
				min = j
			}
		}
		array[i], array[min] = array[min], array[i]
	}
}

// 归并排序
// 1、边界值问题：sort递归结束条件为是否只有一个元素，即i>=j，其实只会有等于j的情况
// 2、辅助空间的使用，每次把待merge的内容一次性拷贝到辅助空间的对应段内，再向目标地址merge
func merge_sort(array []int) {
	n := len(array)
	ans := make([]int, n)
	copy(ans, array)

	var merge func(int, int, int)
	merge = func(left, mid, right int) {

		i, j := left, mid+1
		for k := left; k <= right; k++ {
			ans[k] = array[k]
		}

		for k := left; k <= right; k++ {
			if i > mid {
				array[k] = ans[j]
				j++
			} else if j > right {
				array[k] = ans[i]
				i++
			} else if ans[i] < ans[j] {
				array[k] = ans[i]
				i++
			} else {
				array[k] = ans[j]
				j++
			}
		}
		/*
			idx := left
			i, j := left, mid+1
			for i <= mid || j <= right {
				if i == mid {
					array[idx] = ans[j]
					j++
				} else if j == right {
					array[idx] = ans[i]
					i++
				} else if ans[i] < ans[j] {
					array[idx] = ans[i]
					i++
				} else {
					array[idx] = ans[j]
					j++
				}

				idx++
			}
			copy(ans[left:right+1], array[left:right+1])
		*/
	}

	var sort func(i, j int)
	sort = func(i, j int) {
		if i >= j {
			return
		}

		mid := i + (j-i)>>1
		sort(i, mid)
		sort(mid+1, j)
		merge(i, mid, j)
		/*
			if j-i <= 1 {
				return
			}

			mid := i + (j-i)>>1
			sort(i, mid)
			sort(mid+1, j)
			merge(i, mid, j)
		*/
	}
	sort(0, n-1)
}

// 快速排序
// 第一次优化：修复k值未递增问题
// 第二次优化：修复边界问题
// 第三次优化：重写，修正边界问题，去掉递归过程的slicing操作、减少循环中的交换次数
// 三次优化后效率与《Algorithm》版本实现依然存在差距：
//	BenchmarkSort/quick_sort-12          	    2196	    545831 ns/op	       0 B/op	       0 allocs/op
//	BenchmarkSort/quick_sort_old-12      	    2290	    521419 ns/op	       0 B/op	       0 allocs/op
// 总结了一下，差距在于交换次数
// 本次实现的交换逻辑
//	1，2，5，3，4（输入）
//	1，2，5，3，4（1与1交换）
//	1，2，5，3，4（2与2交换）
//	1，2，5，3，4（未交换）
//	1，2，3，5，4（3与5交换）
//	1，2，3，4，5（循环外，4与5交换）
//	交换次数为4次
//
// 《Algorithm》实现
//	1，2，5，3，4（输入）
//	1，2，5，3，4（第一轮未交换）
//	1，2，4，3，5（第二轮4与5交换）
//	1，2，3，4，5（3与4交换）
//	交换次数为2次
// 由于交换需要有临时变量，而比较不需要，所以交换的代价要比比较高
// 第四次，几乎与《Algorithm》版本一致了，仅对赋值部分采用了golang语法糖
// 性能几乎无异
//	BenchmarkSort/quick_sort-12          	    2268	    528591 ns/op	       0 B/op	       0 allocs/op
//	BenchmarkSort/quick_sort_old-12      	    2236	    527855 ns/op	       0 B/op	       0 allocs/op

func quick_sort(array []int) {
	n := len(array)
	if n <= 1 {
		return
	}

	var sort func(int, int)
	sort = func(lo, hi int) {
		if lo >= hi {
			return
		}

		left, right := lo+1, hi
		mid := array[lo]
		for {
			for array[left] <= mid {
				if hi == left {
					break
				}
				left++
			}
			for array[right] > mid {
				right--
				if lo == right {
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

		array[lo], array[right] = array[right], array[lo]

		sort(lo, right-1)
		sort(right+1, hi)
	}
	sort(0, n-1)

	// 第三次优化
	/*
		n := len(array)
		if n <= 1 {
			return
		}

		var sort func(int, int)
		sort = func(lo, hi int) {
			if lo >= hi {
				return
			}

			w := lo
			for i := lo; i < hi; i++ {
				if array[i] < array[hi] {
					array[w], array[i] = array[i], array[w]
					w++
				}
			}
			array[w], array[hi] = array[hi], array[w]
			sort(lo, w-1)
			sort(w+1, hi)
		}
		sort(0, n-1)
	*/

	// 第一、二次优化
	/*
		n := len(array)
		if n <= 1 {
			return
		}

		k := 0
		for i := 1; i < n; i++ {
			if array[i] < array[k] {
				array[i], array[k] = array[k], array[i]
				array[k+1], array[i] = array[i], array[k+1]
				k += 1
			}
		}

		quick_sort(array[0:k])
		quick_sort(array[k+1:])
	*/
}

func quick_sort_old(array []int) {
	n := len(array)

	var sort func(lo, hi int)
	sort = func(lo, hi int) {
		// 边界条件检查，当待排序数组中元素数量小于等于1时，则认为输入数组已经有序
		if lo >= hi {
			// 只有一个元素的情况，无需排序
			return
		}

		begin := lo
		end := hi
		mid := array[lo]
		lo++

		for {
			// 使用<=是因为，等于mid的左边的值，无需移动
			for array[lo] <= mid {
				if lo == end {
					break
				}
				lo++
			}
			for array[hi] > mid {
				// 首先递减索引，如果索引减少到了最小值，即begin处，begin是选取的基准值，所以i一�?>=j
				// 退出外循环后，其实begin和j是相等的
				hi--
				if hi == begin {
					// begin是mid值所在位置的索引，是不参与交换的，j移动到begin，说明当前轮已经结束
					break
				}
			}
			if lo >= hi {
				break
			}

			tmp := array[lo]
			array[lo] = array[hi]
			array[hi] = tmp
			lo++
			hi--
		}

		array[begin] = array[hi]
		array[hi] = mid

		sort(begin, hi-1)
		sort(hi+1, end)
	}
	sort(0, n-1)
}
