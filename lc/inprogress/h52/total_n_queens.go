package h52

// 思路
//	与h51一个思路
func totalNQueens(n int) int {
	var mask [3]int

	var sc int // solution count
	var qc int // queen count
	backtrace(n, 0, mask, &sc, &qc)
	return sc
}

func backtrace(n, line int, mask [3]int, cnt, qc *int) {
	for i := 0; i < n; i++ {
		var (
			ic  = 1 << i
			ix1 = 1 << (i - line + n - 1)
			ix2 = 1 << (i + line)
		)
		if !(mask[0]&ic > 0 || mask[1]&ix1 > 0 || mask[2]&ix2 > 0) {
			mask[0] |= 1 << i
			mask[1] |= 1 << (i - line + n - 1)
			mask[2] |= 1 << (i + line)
			*qc++
			if *qc == n {
				*cnt++
			} else {
				backtrace(n, line+1, mask, cnt, qc)
			}
			*qc--
			mask[0] &= ^ic
			mask[1] &= ^ix1
			mask[2] &= ^ix2
		}
	}
}

func totalNQueens2(n int) int {
	var column int
	var x1 int
	var x2 int

	var sc int // solution count
	var qc int // queen count
	backtrace2(n, 0, column, x1, x2, &sc, &qc)
	return sc
}

func backtrace2(n, line, column, x1, x2 int, cnt, qc *int) {
	for i := 0; i < n; i++ {
		var (
			ic  = 1 << i
			ix1 = 1 << (i - line + n - 1)
			ix2 = 1 << (i + line)
		)
		if !(column&ic > 0 || x1&ix1 > 0 || x2&ix2 > 0) {
			column |= 1 << i
			x1 |= 1 << (i - line + n - 1)
			x2 |= 1 << (i + line)
			*qc++
			if *qc == n {
				*cnt++
			} else {
				backtrace2(n, line+1, column, x1, x2, cnt, qc)
			}
			*qc--
			column &= ^ic
			x1 &= ^ix1
			x2 &= ^ix2
		}
	}
}

func totalNQueens1(n int) int {
	var column int
	var x1 int
	var x2 int

	var sc int // solution count
	var qc int // queen count

	var backtrace func(line int)
	backtrace = func(line int) {
		for i := 0; i < n; i++ {
			var (
				ic  = 1 << i
				ix1 = 1 << (i - line + n - 1)
				ix2 = 1 << (i + line)
			)
			if !(column&ic > 0 || x1&ix1 > 0 || x2&ix2 > 0) {
				column |= 1 << i
				x1 |= 1 << (i - line + n - 1)
				x2 |= 1 << (i + line)
				qc++
				if qc == n {
					sc++
				} else {
					backtrace(line + 1)
				}
				qc--
				column &= ^ic
				x1 &= ^ix1
				x2 &= ^ix2
			}
		}
	}
	backtrace(0)

	return sc
}
