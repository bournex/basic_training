package selectsort

// 选择排序
// 方法：左侧有序，遍历右侧未排序元素，选择最小的交换到左侧有序序列尾部，直到达到序列尾部
// 特征：移动N次，比较N²/2次
// 分析：由于每一轮外循环之间不存在关系，所以上一次扫描全部元素对下一次扫描没有帮助，稳定

type SelectSorter struct{}

func (ss *SelectSorter) Name() string {
	return "select sort"
}

// Sort sort
func (ss *SelectSorter) Sort(arr []int) {

}

func (ss *SelectSorter) SortIndex(arr []int, lo, hi int) {
	// 左侧有序，右侧无序，左侧序列起始长度为0，右侧序列等于数组长度
	for i := 0; i <= hi; i++ {
		min := i
		// 遍历右侧序列，找到最小值的索引
		for j := i + 1; j <= hi; j++ {
			if arr[j] < arr[min] {
				min = j
			}
		}
		// 替换最小值到左侧序列末尾
		tmp := arr[i]
		arr[i] = arr[min]
		arr[min] = tmp
	}
}
