/*
 * @lc app=leetcode id=1 lang=golang
 *
 * [1] Two Sum
 *
 * https://leetcode.com/problems/two-sum/description/
 *
 * algorithms
 * Easy (40.20%)
 * Total Accepted:    1.4M
 * Total Submissions: 3.5M
 * Testcase Example:  '[2,7,11,15]\n9'
 *
 * Given an array of integers, return indices of the two numbers such that they
 * add up to a specific target.
 * 
 * You may assume that each input would have exactly one solution, and you may
 * not use the same element twice.
 * 
 * Example:
 * 
 * 
 * Given nums = [2, 7, 11, 15], target = 9,
 * 
 * Because nums[0] + nums[1] = 2 + 7 = 9,
 * return [0, 1].
 * 
 * 
 * 
 * 
 */
 func twoSum(nums []int, target int) []int {
    var result = [2]int {0,0}
    if len(nums) < 2 {
        return nil
    }
    
    for i := 0 ; i < len(nums) - 1; i++ {
        for j := i + 1; j < len(nums); j++ {  
            if(nums[i] + nums[j] == target){
                result[0] = i
                result[1] = j
                return result[:]  //返回结果
            }
        }
    }
    return nil    
}

