func maxProduct(nums []int) int {
    //正整数数组，只需要记录最大和次大，否则还需记录最小和次小
	max1, max2 := math.MinInt64, math.MinInt64
	for _, num := range nums {
		if num > max1 {
			max2 = max1
			max1 = num
		} else if num > max2 {
			max2 = num
		}
	}
	max1--
	max2--
	return max1 * max2
}