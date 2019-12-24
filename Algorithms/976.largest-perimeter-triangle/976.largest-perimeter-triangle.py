class Solution(object):
    def largestPerimeter(self, A):
        """
        :type A: List[int]
        :rtype: int
        """
        A.sort()  #先对A数组排序
        for i in xrange(len(A) - 3,-1,-1):
            if A[i] + A[i+1] > A[i+2]:  #判断A数组能否构成三角形：两边之和是否大于第三边
                return A[i] + A[i+1] + A[i+2]  #条件满足，返回周长
            
        return 0

