#
# @lc app=leetcode id=985 lang=python
#
# [985] Sum of Even Numbers After Queries
#
# https://leetcode.com/problems/sum-of-even-numbers-after-queries/description/
#
# algorithms
# Easy (68.44%)
# Total Accepted:    8.9K
# Total Submissions: 13.1K
# Testcase Example:  '[1,2,3,4]\n[[1,0],[-3,1],[-4,0],[2,3]]'
#
# We have an array A of integers, and an array queries of queries.
# 
# For the i-th query val = queries[i][0], index = queries[i][1], we add val to
# A[index].  Then, the answer to the i-th query is the sum of the even values
# of A.
# 
# (Here, the given index = queries[i][1] is a 0-based index, and each query
# permanently modifies the array A.)
# 
# Return the answer to all queries.  Your answer array should have answer[i] as
# the answer to the i-th query.
# 
# 
# 
# Example 1:
# 
# 
# Input: A = [1,2,3,4], queries = [[1,0],[-3,1],[-4,0],[2,3]]
# Output: [8,6,2,4]
# Explanation: 
# At the beginning, the array is [1,2,3,4].
# After adding 1 to A[0], the array is [2,2,3,4], and the sum of even values is
# 2 + 2 + 4 = 8.
# After adding -3 to A[1], the array is [2,-1,3,4], and the sum of even values
# is 2 + 4 = 6.
# After adding -4 to A[0], the array is [-2,-1,3,4], and the sum of even values
# is -2 + 4 = 2.
# After adding 2 to A[3], the array is [-2,-1,3,6], and the sum of even values
# is -2 + 6 = 4.
# 
# 
# 
# 
# Note:
# 
# 
# 1 <= A.length <= 10000
# -10000 <= A[i] <= 10000
# 1 <= queries.length <= 10000
# -10000 <= queries[i][0] <= 10000
# 0 <= queries[i][1] < A.length
# 
# 
#
class Solution(object):
    def sumEvenAfterQueries(self, A, queries):
        """
        :type A: List[int]
        :type queries: List[List[int]]
        :rtype: List[int]
        """
        S = sum(x for x in A if x % 2 == 0) #计算数组A中的所有偶数和
        ans = []
        
        for x,k in queries:
            if A[k] % 2 == 0:  #如果要替换的数是一个偶数
                S -= A[k]  #先减去这个将要替换的数
            A[k] += x 
            if A[k] % 2 == 0: #如果新替换的数是一个偶数
                S += A[k] #加上这个偶数
            ans.append(S)
            
        return ans
    
    
    #给定一个数组A，对其中的数做以下操作
    #1,更新：A[index] += val
    #2,查询：问A中所有偶数的和
        
