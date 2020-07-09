package insertsort

// 插入排序
// 方法：左侧有序，选右侧第一个元素，插入到左侧有序列表，大于插入元素的向右移动
// 特征：最好情况比较N-1次、交换0次，最差情况比较N²/2次，交换N²/2次
// 分析：最好情况出现在序列已排序的状态时，最差情况出现在序列完全倒序时。
//		交换次数与倒置数量相等，当序列完全倒序时，任意两个元素之间都是倒置的
//		根据组合数公式有倒置数量为N*(N-1)/2 = N²/2 - N/2

// InsertSorter insertsorter
type InsertSorter struct{}

func (is *InsertSorter) Name() string {
	return "insert sort"
}

// Sort sort
func (is *InsertSorter) Sort(arr []int) {

}

func (is *InsertSorter) SortIndex(arr []int, lo, hi int) {
	// 左侧有序、右侧无序，左侧序列起始长度为1
	for i := 1; i <= hi; i++ {
		// 取右侧第一个元素与有序序列元素从后往前比较，所以是j--
		for j := i; j > 0; j-- {
			// 如果比前一个元素小，则与前一个交换，抵消一次倒置
			if arr[j] < arr[j-1] {
				tmp := arr[j]
				arr[j] = arr[j-1]
				arr[j-1] = tmp
			} else {
				// 如果大于等于前一个元素，则完成当前元素移动，跳出内循环
				break
			}
		}
		// 内循环完成，左侧保持有序
	}
}
