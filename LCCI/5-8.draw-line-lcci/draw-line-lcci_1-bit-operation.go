func drawLine(length int, w int, x1 int, x2 int, y int) []int {
	//解法1，位运算法
    ret := make([]int, length)
	for i := range ret {
		var i32 int32

		for k := 0 + i*32; k < 32+i*32; k++ {
			if k >= x1+y*w && k <= x2+y*w {
				i32 |= 1 << (31 - uint(k-i*32))
			}
		}
		ret[i] = int(i32)
	}
	return ret
}