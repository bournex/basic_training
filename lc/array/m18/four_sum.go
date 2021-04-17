package m18

import "sort"

func fourSum(nums []int, target int) [][]int {
	length := len(nums)
	if length < 4 {
		return [][]int{}
	}

	result := make([][]int, 0)
	sort.Ints(nums)

	for i := 0; i < length-3; i++ {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		for j := i + 1; j < length-2; j++ {
			if j > i+1 && nums[j] == nums[j-1] {
				continue
			}

			m := j + 1
			n := length - 1
			for m < n {
				if m > j+1 && nums[m] == nums[m-1] {
					m++
					continue
				}
				if n < length-1 && nums[n] == nums[n+1] {
					n--
					continue
				}
				if nums[i]+nums[j]+nums[m]+nums[n] > target {
					n--
				} else if nums[i]+nums[j]+nums[m]+nums[n] < target {
					m++
				} else {
					result = append(result, []int{nums[i], nums[j], nums[m], nums[n]})
					m++
					n--
				}
			}
		}
	}
	return result
}
