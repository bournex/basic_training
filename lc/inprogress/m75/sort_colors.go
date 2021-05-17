package m75

// 思路
//	意识到是双指针了，但是对于细节思考不到位，本不是难题，还是参阅了答案
// 时候用i从0遍历到n-1，左右各维护一个指针，如果i位置为2，后面的值交换，但是有可能right位置本身就是2，所以交换完了之后，还不能递增i，需要再判断i位置是否为2
// 当i位置不为2时，判断是否为0，如果为0，则与left位置交换，官方答案没有使用for循环处理left指针，下面的图表示处理过程中的指针位置
//	| 都是0 | 都是1 | 未知  | 都是2 |
//        left     i    right
// left在循环中可能前进也可能不前进，但是i每次都前进，所以有i>=left，当i==left时，有
//  | 都是0 | 未知 | 都是2 |
//        left	right
//         i
// 此时i迭代过程中还没有发现1，如果遇到1了，i会移动，left不会移动，如果遇到的是0，则left和i都会移动一位，如果遇到2了，则i会加一，right会减一，left不动
// 所以left和i之间只会有1，而left移动也只会因为i碰到了0
// 有一种情形是，i开始遇到的都是2
// | 都是1 | 未知 | 都是2 |
//left    i    right
// 这种情况下，i遇到0时，与left交换，可以视为将小于i的所有1，右移了一位，将i处的0插入到空出来的left位置
//
func sortColors(nums []int) {
	left, right := 0, len(nums)-1

	for i := 0; i <= right; i++ {
		for i <= right && nums[i] == 2 {
			nums[i], nums[right] = nums[right], nums[i]
			right--
		}
		if nums[i] == 0 {
			nums[i], nums[left] = nums[left], nums[i]
			left++
		}
	}
}
