package q88

// 思路
//	从后向前merge，减少nums1中数据的拷贝量
func merge(nums1 []int, m int, nums2 []int, n int) {
	i := m - 1
	j := n - 1
	k := len(nums1) - 1 // writing position

	for j >= 0 && i >= 0 {
		if nums1[i] <= nums2[j] {
			nums1[k] = nums2[j]
			j--
		} else {
			nums1[k] = nums1[i]
			i--
		}
		k--
	}

	if j >= 0 {
		copy(nums1[0:], nums2[0:j+1])
	}
}
