package sort0511

// 快速排序
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
		mid := lo
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
				array[right], array[left] = array[left], array[right]
				left++
				right--
			}
		}

		array[right], array[lo] = array[lo], array[right]
		sort(lo, right-1)
		sort(right+1, hi)
	}
	sort(0, n-1)
}

// 归并排序
func merge_sort(array []int) {
	n := len(array)
	ans := make([]int, n)

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

		for i := lo; i <= hi; i++ {
			ans[i] = array[i]
		}

		i, j := lo, mid+1
		idx := lo

		for i <= mid || j <= hi {
			if i > mid {
				array[idx] = ans[j]
				j++
			} else if j > hi {
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
	}
	sort(0, n-1)
}

// 堆排序
// 对比旧版本实现性能
//	BenchmarkSort/heap_sort-12          	    1248	    921101 ns/op	   81920 B/op	       1 allocs/op
//	BenchmarkSort/heap_sort_closure-12  	     710	   1709181 ns/op	       0 B/op	       0 allocs/op
// 旧版本为原址排序，空间复杂度为O(1)，对前1/2元素进行排序，复杂度为O((n*logn)/2)，完成建堆，然后pop大顶堆并下沉顶部元素，复杂度为O((n*logn)/2)
// 新版本借助了辅助空间，空间复杂度为O(n)，对所有元素在辅助空间上建堆，复杂度为O((n*logn)/2)，完成建堆，然后pop大顶堆并下沉顶部元素，复杂度为O((n*logn)/2)
// 刨除常量后时间复杂度都是O(n*logn)，按说不应该出现悬殊的性能差异，而新实现的辅助空间，也仅仅是为了编码方便，并没有带来性能提升
//
// 分析代码后，怀疑是旧版本实现的闭包方法调用导致了执行成本高，所以使用直接代码替换掉了less和exch方法，替换性能对比
//	BenchmarkSort/heap_sort-12          	    1248	    921101 ns/op	   81920 B/op	       1 allocs/op
//	BenchmarkSort/heap_sort_inline-12   	    1736	    690103 ns/op	       0 B/op	       0 allocs/op
//
// 但是闭包函数很短，为什么不会被内联掉呢？怀疑闭包不会被内联，遂改造heap_sort_inline方法，将heap_sort_closure中的闭包函数提出外部，修改后性能比对
//	BenchmarkSort/heap_sort-12          	    1248	    921101 ns/op	   81920 B/op	       1 allocs/op
//	BenchmarkSort/heap_sort_closure-12  	     710	   1709181 ns/op	       0 B/op	       0 allocs/op
//	BenchmarkSort/heap_sort_inline-12   	    1736	    690103 ns/op	       0 B/op	       0 allocs/op
// 所以闭包不会被内联？

func heap_sort(array []int) {
	n := len(array)
	if n <= 1 {
		return
	}

	heap := make([]int, n+1)
	idx := 1

	var rise func(int)
	var sink func(int)

	rise = func(i int) {
		for {
			parent := i >> 1
			if parent > 0 && heap[parent] < heap[i] {
				heap[parent], heap[i] = heap[i], heap[parent]
				i = parent
			} else {
				break
			}
		}
	}

	sink = func(i int) {

		for {
			left, right := i<<1, i<<1+1
			max := i

			if left < idx && heap[max] < heap[left] {
				max = left
			}
			if right < idx && heap[max] < heap[right] {
				max = right
			}
			if max == i {
				break
			} else {
				heap[max], heap[i] = heap[i], heap[max]
			}

			i = max
		}
	}

	for i := 0; i < n; i++ {
		heap[idx] = array[i]
		rise(idx)
		idx++
	}

	for i := n - 1; i >= 0; i-- {
		idx--
		array[i] = heap[1]
		heap[1] = heap[idx]
		sink(1)
	}
}

func heap_sort_closure(array []int) {
	l := len(array)

	exch := func(arr []int, i, j int) {
		tmp := arr[i-1]
		arr[i-1] = arr[j-1]
		arr[j-1] = tmp
	}

	less := func(arr []int, i, j int) bool {
		return arr[i-1] < arr[j-1]
	}

	sink := func(arr []int, k, l int) {

		for (k << 1) <= l {

			idx := k << 1
			if idx < l && less(arr, idx, idx+1) {
				idx++
			}

			if less(arr, idx, k) {
				break
			}

			exch(arr, k, idx)
			// array[k-1], array[idx-1] = array[idx-1], array[k-1]

			k = idx
		}
	}

	for i := l / 2; i >= 1; i-- {
		sink(array, i, l)
	}

	for l > 1 {
		//array[0], array[l-1] = array[l-1], array[0]
		exch(array, 1, l)
		l--
		sink(array, 1, l)
	}
}

func exch(arr []int, i, j int) {
	tmp := arr[i-1]
	arr[i-1] = arr[j-1]
	arr[j-1] = tmp
}

func less(arr []int, i, j int) bool {
	return arr[i-1] < arr[j-1]
}

func heap_sort_inline(array []int) {
	l := len(array)

	sink := func(arr []int, k, l int) {

		for (k << 1) <= l {

			idx := k << 1
			if idx < l && less(arr, idx, idx+1) {
				idx++
			}

			if less(arr, idx, k) {
				break
			}

			exch(arr, k, idx)
			// array[k-1], array[idx-1] = array[idx-1], array[k-1]

			k = idx
		}
	}

	for i := l / 2; i >= 1; i-- {
		sink(array, i, l)
	}

	for l > 1 {
		// array[0], array[l-1] = array[l-1], array[0]
		exch(array, 1, l)
		l--
		sink(array, 1, l)
	}
}
