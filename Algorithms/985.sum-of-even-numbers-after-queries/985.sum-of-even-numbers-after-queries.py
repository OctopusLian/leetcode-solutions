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
        
