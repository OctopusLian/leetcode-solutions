from functools import lru_cache

class Solution:
    @lru_cache(None)
    def numDecodings(self, s: str) -> int:
        if len(s) == 0: return 1
        cnt = 0

        if s[0] != '0': #必须从1~9，把0去掉
            cnt += self.numDecodings(s[1:])

        if 10 <= int(s[0:2]) <= 26: #必须小于等于26才能解码
            cnt += self.numDecodings(s[2:])

        return cnt