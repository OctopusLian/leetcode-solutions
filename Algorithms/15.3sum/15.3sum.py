#
# @lc app=leetcode id=15 lang=python3
#
# [15] 3Sum
#
# https://leetcode.com/problems/3sum/description/
#
# algorithms
# Medium (23.14%)
# Total Accepted:    470.2K
# Total Submissions: 2M
# Testcase Example:  '[-1,0,1,2,-1,-4]'
#
# Given an array nums of n integers, are there elements a, b, c in nums such
# that a + b + c = 0? Find all unique triplets in the array which gives the sum
# of zero.
# 
# Note:
# 
# The solution set must not contain duplicate triplets.
# 
# Example:
# 
# 
# Given array nums = [-1, 0, 1, 2, -1, -4],
# 
# A solution set is:
# [
# ⁠ [-1, 0, 1],
# ⁠ [-1, -1, 2]
# ]
# 
# 
#
class Solution:
    def threeSum(self, nums: 'List[int]') -> 'List[List[int]]':
        res = []
        nums.sort()
        length = len(nums)
        for i in xrange(length-2):
            if nums[i]>0:break
            if i>0 and nums[i]==nums[i-1]:continue
            
            l,r = i+1,length-1
            while l<r:
                total = nums[i]+nums[l]+nums[r]

                if total<0:
                    l+=1
                elif total>0:
                    r-=1
                else:
                    res.append([nums[i], nums[l], nums[r]])
                    while l<r and nums[l]==nums[l+1]:
                        l+=1
                    while l<r and nums[r]==nums[r-1]:
                        r-=1
                    l+=1
                    r-=1
        return res
        
