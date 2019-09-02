/*
 * @lc app=leetcode id=41 lang=c
 *
 * [41] First Missing Positive
 *
 * https://leetcode.com/problems/first-missing-positive/description/
 *
 * algorithms
 * Hard (28.05%)
 * Total Accepted:    187.1K
 * Total Submissions: 666.7K
 * Testcase Example:  '[1,2,0]'
 *
 * Given an unsorted integer array, find the smallest missing positive
 * integer.
 * 
 * Example 1:
 * 
 * 
 * Input: [1,2,0]
 * Output: 3
 * 
 * 
 * Example 2:
 * 
 * 
 * Input: [3,4,-1,1]
 * Output: 2
 * 
 * 
 * Example 3:
 * 
 * 
 * Input: [7,8,9,11,12]
 * Output: 1
 * 
 * 
 * Note:
 * 
 * Your algorithm should run in O(n) time and uses constant extra space.
 * 
 */
int firstMissingPositive(int* nums, int numsSize) {
    if(numsSize == 0)
        return 1;
    
    //第i位存放i+1的数值
    for(int i = 0;i < numsSize;i++){
        if(nums[i] > 0){
            //如果交换的数据大于0且小于i+1且数据不相等，则进行交换（放在合适的位置上）
            while(nums[i] > 0 && nums[i] < i+1 && nums[i] != nums[nums[i] - 1]){
                int temp = nums[nums[i] - 1];
                nums[nums[i] - 1] = nums[i];
                nums[i] = temp;
            }
        }
    }

    //循环寻找不符合要求的数据
    for(int i = 0;i < numsSize;i++){
        if(nums[i] != i+1){
            return i+1;
        }
    }

    //如果都符合要求，则返回长度+1的值
    return numsSize + 1;
}
