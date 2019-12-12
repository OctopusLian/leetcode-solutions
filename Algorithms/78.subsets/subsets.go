//复制数组
func copySlice(src []int) []int {
	temp := []int{}
	for _, i := range src {
		temp = append(temp, i)
	}
	return temp
}

func subsets(nums []int) [][]int {
	result := [][]int{}
	result = append(result, []int{}) //第一步 包含空集
	for i := 0; i < len(nums); i++ {
        length := len(result)
		for j := 0; j < length; j++ {
			//设定两个下标元素,从0开始遍历nums的值
			//里层复制数组,接着append进去下标为j的值组成nums的一个子集,再append进result数组
			result = append(result, append(copySlice(result[j]), nums[i]))
		}
	}
	return result
}
