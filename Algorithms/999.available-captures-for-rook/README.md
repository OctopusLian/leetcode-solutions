# 解法  

## python一行解决  

代码的思想就是将二维数组摊开成一条字符串之后，删除其中的'.'，再搜索其中的'pR'和'Rp'即可，写的有点繁琐，不知道有没有更优雅的解决方案。  

```py
class Solution:
    def numRookCaptures(self, board: List[List[str]]) -> int:
        return (lambda x:x('pR')+x('Rp'))([''.join(board[x]+[' ']+[i[y] for i in board]).replace('.','') for x in range(8) for y in range(8) if board[x][y]=='R'][0].count)
```

链接：https://leetcode-cn.com/problems/available-captures-for-rook/solution/python-yi-xing-jie-jue-by-zhu-mang-4/



