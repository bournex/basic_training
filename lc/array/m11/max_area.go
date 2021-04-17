package m11

func maxArea(height []int) int {
	i, j := 0, len(height)-1

	cap := 0

	for i < j {
		if height[i] > height[j] {
			if height[j]*(j-i) > cap {
				cap = height[j] * (j - i)
			}
			j--
		} else {
			if height[i]*(j-i) > cap {
				cap = height[i] * (j - i)
			}
			i++
		}
	}
	return cap
}
