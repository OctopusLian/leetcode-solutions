/*
 * @lc app=leetcode id=66 lang=golang
 *
 * [66] Plus One
 */
func plusOne(digits []int) []int {
	digits = append(digits,1)
	for i := len(digits) - 1;i > 0;i-- {
		digits[i] = digits[i] + digits[i-1]
		digits[i-1] = digits[i]/10  //取除数
		digits[i] %= 10  //取余数
	}
	if digits[0] == 0 {
		return digits[1:]
	} else {
		return digits
	}
}

