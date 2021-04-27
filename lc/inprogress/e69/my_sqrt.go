package e69

// 思路
//	二分
func mySqrt(x int) int {
	l, r := 0, x
	for l < r {
		mid := l + (r-l)>>1 + 1
		if mid*mid > x {
			r = mid - 1
		} else if mid*mid < x {
			l = mid
		} else {
			return mid
		}
	}
	return l
}
