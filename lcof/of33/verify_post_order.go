package of33

func verifyPostorder(postorder []int) bool {
	if len(postorder) <= 1 {
		return true
	}

	root := postorder[len(postorder)-1]
	i := 0
	for ; i < len(postorder)-1; i++ {
		if root < postorder[i] {
			break
		}
	}
	for j := i; j < len(postorder)-1; j++ {
		if root > postorder[j] {
			return false
		}
	}
	return verifyPostorder(postorder[0:i]) && verifyPostorder(postorder[i:len(postorder)-1])
}
