/*
 * @lc app=leetcode id=485 lang=golang
 *
 * [485] Max Consecutive Ones
 */
func findMaxConsecutiveOnes(nums []int) int {
	var sum, max int
	for _, v := range nums {
		sum += v
		sum *= v
		if sum > max {
			max = sum
		}
	}
	return max
}

