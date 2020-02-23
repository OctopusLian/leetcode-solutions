func printNumbers(n int) []int {
	res := make([]int, 0)
	limit := int(math.Pow(float64(10), float64(n))) - 1
	for i := 1; i <= limit; i++ {
		res = append(res, i)
	}
	return res
}
