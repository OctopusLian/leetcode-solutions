func findMaxConsecutiveOnes(nums []int) (maxCnt int) {
	cnt := 0
	for _, v := range nums {
		if v == 1 {
			cnt++
		} else {
			maxCnt = max(maxCnt, cnt)
			cnt = 0
		}
	}
	maxCnt = max(maxCnt, cnt)
	return
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}