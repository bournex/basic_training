package m55

// https://leetcode-cn.com/problems/jump-game/
// 思考
//	看起来感觉是动态规划题目，需要从左向右计算能达到的最远点
func canJump(nums []int) bool {
	n := len(nums)
	if n <= 1 {
		return true
	}

	var max int
	for i := 0; i < n-1; i++ {
		if i+nums[i] > max {
			max = i + nums[i]
			if max >= n-1 {
				return true
			}
		} else if max <= i {
			break
		}
	}

	return false
}

func canJump1(nums []int) bool {
	n := len(nums)
	if n <= 1 {
		return true
	}

	ans := make([]int, n)
	ans[0] = nums[0]
	for i := 1; i < n-1; i++ {
		if i+nums[i] > ans[i-1] {
			ans[i] = i + nums[i]
		} else {
			ans[i] = ans[i-1]
		}
		if ans[i] >= n-1 {
			return true
		}
	}

	return false
}
