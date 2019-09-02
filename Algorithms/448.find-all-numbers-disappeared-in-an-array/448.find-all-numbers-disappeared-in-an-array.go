/*
 * @lc app=leetcode id=448 lang=golang
 *
 * [448] Find All Numbers Disappeared in an Array
 */
func findDisappearedNumbers(nums []int) []int {
	n := len(nums)  //获取数组的长度
	for i := 0;i < n;i++ {
		for nums[i]-1 != i && nums[i] != nums[nums[i]-1] {
			nums[i],nums[nums[i]-1] = nums[nums[i]-1],nums[i]  //交换位置
		}
	}

	ans := []int{}
	for i := 0;i < n;i++ {
		if nums[i]-1 != i {
			ans = append(ans,i+1)
		}
	}

	return ans

}

