func spiralOrder(matrix [][]int) []int {
    if len(matrix) == 0 || len(matrix[0]) == 0 {
		return nil
	}
	const right, down, left, up = 0, 1, 2, 3
	m, n := len(matrix), len(matrix[0])
	rMin, rMax, cMin, cMax := 0, m-1, 0, n-1 // 上下左右边界
	r, c, direct := 0, 0, right
	result := make([]int, m*n)
	for i := 0; i < len(result); i++ {
		result[i] = matrix[r][c]
		switch direct {
		case right:
			if c < cMax {
				c++
			} else {
				direct, rMin, r = down, rMin+1, r+1
			}
		case down:
			if r < rMax {
				r++
			} else {
				direct, cMax, c = left, cMax-1, c-1
			}
		case left:
			if c > cMin {
				c--
			} else {
				direct, rMax, r = up, rMax-1, r-1
			}
		case up:
			if r > rMin {
				r--
			} else {
				direct, cMin, c = right, cMin+1, c+1
			}
		}
	}
	return result
}
