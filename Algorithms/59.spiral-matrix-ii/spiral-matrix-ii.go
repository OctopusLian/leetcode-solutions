func generateMatrix(n int) [][]int {
    if n <= 0 {
		return nil
	}
	var result = make([][]int, n)
	for i := 0; i < n; i++ {
		result[i] = make([]int, n)
	}
	num := 1
	nRound := n / 2
	if nRound*2 != n {
		nRound = nRound + 1
	}
	for nr := 0; nr < nRound; nr++ {
		tmpLen := n - 2*nr - 1
		if tmpLen == 0 { //奇数，圈只有一个
			result[nr][nr] = num
			continue
		}
		if tmpLen < 0 {
			break
		}
		for i := 0; i < tmpLen; i++ {
			result[nr][nr+i] = num
			num++
		}
		for i := 0; i < tmpLen; i++ {
			result[nr+i][n-nr-1] = num
			num++
		}
		for i := tmpLen; i > 0; i-- {
			result[n-nr-1][nr+i] = num
			num++
		}
		for i := tmpLen; i > 0; i-- {
			result[nr+i][nr] = num
			num++
		}
	}
	return result
}
