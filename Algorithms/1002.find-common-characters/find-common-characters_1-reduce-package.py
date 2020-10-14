class Solution:
    def commonChars(self, A: List[str]) -> List[str]:
        from functools import reduce
        return list(reduce(lambda x, y: x & y, map(collections.Counter, A)).elements())