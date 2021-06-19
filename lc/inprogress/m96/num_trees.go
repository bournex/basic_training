package m96

// 思路
//	看起来是BST的问题，实际上跟BST貌似没什么关系
//	遍历1-n，每个值i都可以作为根节点，把小于i的和大于i的值分别递归处理即可

// 非递归
func numTrees(n int) int {
	if n <= 2 {
		return n
	}
	ans := make([]int, n)
	ans[0] = 1
	ans[1] = 2
	for i := 2; i < n; i++ {
		var sum int
		for j := 0; j <= i; j++ {
			if j == 0 {
				sum += ans[i-1]
			} else if j == i {
				sum += ans[i-1]
			} else {
				sum += ans[j-1] * ans[i-j-1]
			}
		}
		ans[i] = sum
	}
	return ans[n-1]
}

// 递归
func numTrees1(n int) int {
	if n <= 2 {
		return n
	}
	var sum int
	for i := 1; i <= n; i++ {
		left := numTrees1(i - 1)
		right := numTrees1(n - i)
		if left == 0 {
			sum += right
		} else if right == 0 {
			sum += left
		} else {
			sum += left * right
		}
	}
	return sum
}
