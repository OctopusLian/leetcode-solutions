class Solution:
    def arrangeCoins(self, n: int) -> int:
        res = ((1+8*n)**0.5-1)/2
        return int(res)