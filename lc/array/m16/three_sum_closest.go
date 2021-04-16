package m16

import (
	"sort"

	"github.com/bournex/basic_training/lc/tree"
)

func threeSumClosest(nums []int, target int) int {
	sort.Ints(nums)

	abs := func(n int) int {
		if n < 0 {
			return -1 * n
		}
		return n
	}

	min := tree.INT_MAX
	sum := tree.INT_MAX

	for i := 0; i < len(nums)-2; i++ {
		j, k := i+1, len(nums)-1
		for j < k {
			t := nums[i] + nums[j] + nums[k]
			if abs(t-target) < min {
				min = abs(t - target)
				sum = t
			}

			if t > target {
				k--
			} else if t < target {
				j++
			} else {
				return sum
			}
		}
	}
	return sum
}
