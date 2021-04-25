package m45

// 思考
// 学习
//	正向贪心
//	从nums[0]开始，计算nums[0]能跳的区间范围内，能达到的最远的地方，选做第二跳
//	再在第二跳的范围内，迭代上述过程，直到抵达最后的节点
//
//
func jump(nums []int) int {
	n := len(nums)
	if n <= 1 {
		return 0
	}

	// 遍历nums，记录每个位置能跳到的最远处
	//
	max := func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}

	end := 0
	step := 0
	maxPos := 0
	for i := 0; i < n-1; i++ {
		maxPos = max(maxPos, nums[i]+i)
		if i == end {
			end = maxPos
			step++
		}
	}

	return step
}
