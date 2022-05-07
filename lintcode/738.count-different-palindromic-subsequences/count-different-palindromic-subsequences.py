class Solution:
    """
    @param str: a string S
    @return: the number of different non-empty palindromic subsequences in S
    """
    def countPalindSubseq(self, str):
        # write your code here
        mod = 1000000007
        ans = 0
        n = len(str)
        if n== 0 : 
            return 0
        dp=[[[0 for i in range(4)] for i in range(n)]for i in  range (3)]
        for leng in range(1, n  + 1) :  #枚举区间长度
            for i in range(0, n- leng + 1) :    #枚举区间起点
                for x in range(0, 4) :	#枚举开头字母
                    j = i + leng - 1   #区间终点
                    c = chr(97 + x)
                    if leng == 1 :       #如果长度为1
                        if str[i] == c  :  #当前区间开头字母和枚举字母相同，单个字母组成回文子序列
                            ans = 1
                        else :
                            ans = 0
                    else :
                        if str[i] != c   :    #如果当前字母与枚举的不相同
                            ans = dp[1][i+1][x]  #当前答案则为上一轮中(长度为leng-1)时的i+1位置至j位置以'a'+x开头的回文子序列数量。
                        
                        elif str[j] != c  : #如果终点字母与枚举的不相同
                            ans = dp[1][i][x]	#当前答案则为上一轮中(长度为leng-1)时的i位置至j-1位置的以'a'+x开头的回文子序列数量。
                        
                        else :	 #如果起点和终点字母均与枚举字母相同
                            ans = 2
                            if leng > 2 :
                                for y in range(0, 4) :		#每次加入长度为leng-2时，i+1位置后的回文子序列数量(即i+1至j-1区间的回文子序列数量)。
                                    ans += dp[0][i+1][y]
                                    ans %= mod
                    
                      
                    dp[2][i][x] = ans
            for i in range(0, 2) :
                for j in range(0, n) :
                    for x in range(0, 4) :
                        dp[i][j][x] = dp[i+1][j][x]	#每次将本次计算的答案dp[2][i][x]保存到dp[1][i][x]中，则原来的dp[1][i][x]保存到dp[0][i][x]中
        ret = 0
        for x in range(0, 4) :
            ret = (ret + dp[2][0][x]) % mod  #最后统计四种字母开头的回文子序列数量求和即可
        return ret