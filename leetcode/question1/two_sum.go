package question1

func twoSum(nums []int, target int) []int {
	m := make(map[int]int)
	for i, v := range nums {
		if i1, ok := m[target-v]; ok {
			return []int{i1, i}
		}
		m[v] = i
	}
	return nil
}
