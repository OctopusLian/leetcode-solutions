from functools import lru_cache

class Solution:
    @lru_cache(None)
    
    def isMatch(self, s: str, p: str) -> bool:
        # 没有 . 和 *
        # 推荐: if not p: return not s
        if len(p) == 0: return len(s) == 0
        # process first char and "."
        first_match = bool(s and p[0] in {s[0],'.'})

        if len(p) >= 2 and p[1] == '*':
            # 处理 *
            # process "*": 0个或者多个prev char
            # s: "abc", p: "a*abc"
            return self.isMatch(s,p[2:]) or first_match and self.isMatch(s[1:],p)

        return first_match and self.isMatch(s[1:],p[1:])