func removeDuplicates(nums []int) int {
	//解法2，双指针-快慢指针
	if len(nums) == 0 { //考虑特殊情况
		return 0
	}
	i := 0                           //慢指针
	for j := 1; j < len(nums); j++ { //j是快指针
		if nums[j] != nums[i] {
			i++
			nums[i] = nums[j]
		}
	}
	return i + 1
}