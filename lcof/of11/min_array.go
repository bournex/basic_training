package of11

// 线索
//	1.数组中随机取一中间点，如果比最右侧值大，则中间点左侧值一定都大于等于最右侧值，最小值存在于中间点右侧
//	2.数组中随机取一中间点，如果比最右侧值小，则中间点右侧值一定都小于等于最右侧值，最小值存在于中间点左侧或中间点本身
//	3.数组中随机取一中间点，如果与最右侧值相等，则最小值可能在中间点的两侧
// 思路
//	二分查找
//	以最右侧元素为参照点
//	如果left和right的中间点大于参照点，则说明最小值在有半边，不包括中间点（线索2）
//	如果left和right的中间点小于参照点，则说明最小值在左半边，可能包括中间点（线索1）
//	如果以上两条都不满足，则说明中间点等于参照点，此时最小值可能mid两边中的任意一边（线索3）
//	但由于mid等于参照点，所以收缩右侧是安全的
//	为什么收缩左侧不安全：因为左边的值可能小于mid([1,3,5])，也可能等于mid([2,2,2])，也可能大于mid[5,1,2,3,4]，当left小于mid时，收缩会丢失最小值
//	收缩右侧等于间接左移了mid，所以循环是趋向于关闭的，对于循环也是安全的

func minArray(numbers []int) int {
	n := len(numbers)
	if n <= 0 {
		return 0
	}

	left, right := 0, n-1
	for left < right {
		mid := left + (right-left)>>1
		if numbers[mid] < numbers[right] {
			right = mid
		} else if numbers[mid] > numbers[right] {
			left = mid + 1
		} else {
			right--
		}
	}
	return numbers[left]
}

// 下面是一次失败的尝试，曾努力考虑所有细节

/*
func min(numbers []int, left, right int) int {
	if right-left == 0 {
		return numbers[left]
	} else if right-left == 1 {
		minf := func(a, b int) int {
			if a > b {
				return b
			}
			return a
		}
		return minf(numbers[left], numbers[right])
	}

	iMid := (right - left) >> 1

	if numbers[iMid] > numbers[right] { // [2,2,2,3,1,2,2]
		// left = iMid + 1
		return min(numbers, iMid+1, right)
	} else if numbers[iMid] < numbers[right] { // [2,2,3,1,2,2,2]
		// right = iMid - 1
		return min(numbers, left, iMid-1)
	} else {
		if numbers[iMid] > numbers[left] { // non-exist
			return min(numbers, left, iMid-1)
		} else if numbers[iMid] < numbers[left] { // [3,1,2,2,2,2,2]
			// right = iMid - 1
			return min(numbers, left, iMid-1)
		} else { // [2,3,1,2,2,2,2],[2,2,2,2,3,1,2]
			left = iMid + 1
			n1 := min(numbers, left, iMid-1)
			n2 := min(numbers, iMid+1, right)
			if n1 > n2 {
				return n2
			}
			return n1
		}
	}
}

*/
