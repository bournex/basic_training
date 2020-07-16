package str

// 类似快排的方法
// 属于高位优先的排序的一种，与msd区别在于，递归过程不对字符表进行统计，而采用二/三分法
// 适合：有重复键或重复前缀的字符串数组排序
func Quick3(ms []string) {
	l := len(ms)
	quick3(ms, 0, l-1, 0)
}

// ms 	- 完整待排序数组
// lo	- 待排序位置的低位，闭区间
// hi	- 待排序位置的高位，闭区间
// idx	- 排序依据列，假定ms中所有字符串在idx左侧的部分已经有序
func quick3(ms []string, lo, hi, idx int) {
	// 边界条件
	if hi <= lo {
		return
	}

	var t int
	if len(ms[lo]) == idx {
		t = -1
	} else {
		t = int(ms[lo][idx])
	}

	l, h := lo, hi

	// 以下循环将ms在[lo,hi]部分的数据，依据第idx列的ascii值(c)和分界值t进行划分
	// c < t		c == t		c > t
	// [lo, l-1]	[l,h]		[h+1, hi]

	for i := lo + 1; i < h; {
		var c int
		if len(ms[i]) == idx {
			c = -1
		} else {
			c = int(ms[i][idx])
		}

		if c < t {
			tmp := ms[l]
			ms[l] = ms[i]
			ms[i] = tmp
			i++
			l++
		} else if c > t {
			tmp := ms[h]
			ms[h] = ms[i]
			ms[i] = tmp
			h--
			continue
		} else {
			i++
		}
	}

	quick3(ms, lo, l-1, idx)
	// 如果切分用的子串已经到达末尾，则不进行递归，说明l-h之间的内容已经完成排序
	if t > 0 {
		quick3(ms, l, h, idx+1)
	}
	quick3(ms, h+1, hi, idx)
}
