func minimumAbsDifference(arr []int) [][]int {
    sort.Ints(arr)
	l := len(arr)
	var result [][]int
	min := arr[l-1] - arr[0]
	for i := 1; i < l; i++ {
		cha := arr[i] - arr[i-1]
		if cha < min {
			min = cha
			result = [][]int{}
			result = append(result, []int{arr[i-1], arr[i]})
		} else if cha == min {
			result = append(result, []int{arr[i-1], arr[i]})
		}
	}
	return result
}