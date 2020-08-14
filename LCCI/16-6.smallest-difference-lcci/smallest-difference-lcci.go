func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}
func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

func smallestDifference(a []int, b []int) int {
	sort.Ints(a)
	sort.Ints(b)

	lengthA := len(a)
	lengthB := len(b)
	i, j, res := 0, 0, math.MaxInt32
	for i < lengthA && j < lengthB {
		res = min(res, abs(a[i] - b[j]))
		if a[i] > b[j] {
			j++
		}else {
			i++
		}
	}
	return res
}