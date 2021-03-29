package question15

import "sort"

func threeSum(nums []int) [][]int {
	if len(nums) < 3 {
		return nil
	}
	sort.Ints(nums)
	res := make([][]int, 0)

	for i := 0; i < len(nums); i++ {
		j := i + 1         // 正向索引
		k := len(nums) - 1 // 反向索引

		if i > 0 && nums[i] == nums[i-1] {
			continue
		}

		for j < len(nums) && j != k {
			if j > i+1 && nums[j] == nums[j-1] {
				j++
				continue
			}

			if nums[i]+nums[j]+nums[k] < 0 {
				j++
				continue
			} else if nums[i]+nums[j]+nums[k] > 0 {
				k--
			} else if nums[i]+nums[j]+nums[k] == 0 {
				res = append(res, []int{nums[i], nums[j], nums[k]})
				j++
			}
		}
	}

	return res
}
