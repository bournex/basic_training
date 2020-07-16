package str

// 高位优先字符串排序
// 线性时间内完成字符串数组排序
// 适合：重复键较少、没有超长键、随机字符串
// 重复键会使比较接近或达到线性时间，而超长键会导致cnt占用空间暴涨
func Msd(ms []string) {
	l := len(ms)
	p := new(msd)
	p.arr = make([]string, l)
	p.sort(ms, 0, l-1, 0)
}

type msd struct {
	arr []string
}

// ms		- 待排序数组
// start	- 要排序序列的起始位置索引，从0开始
// end		- 要排序序列的结束位置索引
func (s *msd) sort(ms []string, start, end, idx int) {
	l := end - start + 1
	// TODO：对于较小规模的数据，此处应调整为插排，用于解决大量小型数组排序问题，可显著提高ms规模较大时的性能表现
	if l < 2 {
		return
	}

	// 第0位留空
	// 第1位留做轮空的字符串计数
	// 第2-257位留做ascii码计数
	cnt := make([]int, 256+2)

	for i := 0; i < l; i++ {
		if len(ms[i+start]) == idx {
			// 当前idx在ms[i]处越界，cnt第1位增加计数
			cnt[1]++
		} else {
			// ascii+2表达在cnt中的索引，cnt[..]++增加该字符计数
			cnt[ms[i+start][idx]+2]++
		}
	}
	for i := 1; i < 256+1; i++ {
		// 累加cnt
		cnt[i+1] += cnt[i]
	}
	for i := 0; i < l; i++ {
		// 此时错位使用cnt
		// 第0位为在idx处越界的短字符串起始索引
		// 第1-256位为ascii字符的起始位置
		if len(ms[i+start]) == idx {
			s.arr[cnt[0]+start] = ms[i+start]
			cnt[0]++
		} else {
			s.arr[cnt[ms[i+start][idx]+1]+start] = ms[i+start]
			cnt[ms[i+start][idx]+1]++
		}
	}
	for i := 0; i < l; i++ {
		// ms[i+start] = ms[i+start][:idx] + arr[i][idx:]
		ms[i+start] = s.arr[i+start]
	}

	// 经过循环步骤3后
	// 第0-255位为从start开始，有效ascii字符的起始索引位置
	for i := 0; i < 256; i++ {
		s.sort(ms, start+cnt[i], start+cnt[i+1]-1, idx+1)
	}
}
