package shellsort

// 希尔排序
// 方法：跳跃的插入排序，跳跃间隔选择自定，在外循环中逐渐降为1，当间隔为1时，等同于插排
// 特征：希尔排序的特征不太容易量化，一般认为在N到N²之间
// 分析：通过快速将末尾的小元素前移，大量减少比较次数。只需要基于插排稍微修改就可以实现

type ShellSorter struct{}

func (ss *ShellSorter) Name() string {
	return "shell sort"
}

// Sort sort
func (ss *ShellSorter) Sort(arr []int) {

}

func (ss *ShellSorter) SortIndex(arr []int, lo, hi int) {
	n := len(arr)
	h := 1

	// 一种希尔元素间隔的选择方式
	for h < n/3 {
		h = 3*h + 1
	}

	for h >= 1 {
		// 当h>1时，是对不相邻的局部排序
		// 当h=1时，是对局部有序的数组进行插入排序

		for i := h; i < n; i++ {
			// 外循环遍历多个元素不相邻的数组

			for j := i; j >= h; j -= h {
				// 内循环处理单个元素不相邻数组中的元素插入

				if arr[j] < arr[j-h] {
					tmp := arr[j]
					arr[j] = arr[j-h]
					arr[j-h] = tmp
				}
			}
		}

		h = h / 3
		// fmt.Println("shell h is ", h)
	}
}
