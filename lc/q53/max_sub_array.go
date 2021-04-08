package q53

func maxSubArray(nums []int) int {
	ans := 0
	stat := make([]int, len(nums))
	for i, v := range nums {
		if i > 0 {
			if stat[i-1] < 0 {
				stat[i] = v
			} else {
				stat[i] = stat[i-1] + v
			}
			if stat[i] > ans {
				ans = stat[i]
			}
		} else {
			ans, stat[0] = v, v
		}
	}
	return ans
}

// 官方解答的优势
//  利用并修改了原数组（注意题干是否有此方面限制），优化了空间复杂度
//	i从1开始，对只有一个元素的数组在循环外处理，减少了n次循环中的if判断

func maxSubArray1(nums []int) int {
	max := nums[0]
	for i := 1; i < len(nums); i++ {
		if nums[i]+nums[i-1] > nums[i] {
			nums[i] += nums[i-1]
		}
		if nums[i] > max {
			max = nums[i]
		}
	}
	return max
}

// 作者：LeetCode-Solution
// 链接：https://leetcode-cn.com/problems/maximum-subarray/solution/zui-da-zi-xu-he-by-leetcode-solution/
// 来源：力扣（LeetCode）
// 著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。
