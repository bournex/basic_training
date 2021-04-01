package of15

func hammingWeight(num uint32) int {
	var n int
	for num > 0 {
		if num&0x00000001 == 0x00000001 {
			n++
		}
		num = num >> 1
	}
	return n
}
