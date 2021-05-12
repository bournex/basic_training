package graphizability

import "sort"

// Havel Hakimi定理判定度序列可图性
func havelHakimi(sequence []int) bool {
	n := len(sequence)

	sort := func(lo, hi int) {
		sort.Ints(sequence[lo : hi+1])
		for i := 0; i <= (hi-lo)/2; i++ {
			sequence[lo+i], sequence[hi-i] = sequence[hi-i], sequence[lo+i]
		}
	}

	for i := 0; i < n; i++ {
		sort(i, n-1)
		d1 := sequence[i]
		for j := 0; j < d1 && i+j+1 < n; j++ {
			sequence[i+j+1]--
			if sequence[i+j+1] < 0 {
				return false
			}
		}
	}

	return true
}
