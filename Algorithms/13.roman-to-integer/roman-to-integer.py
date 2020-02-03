class Solution:
    def romanToInt(self, s: str) -> int:
        d = {'I': 1, 'IV': -1, 'V': 5, 'IX': -1,
             'X': 10, 'XL': -10, 'L': 50, 'XC': -10,
             'C': 100, 'CD': -100, 'D': 500, 'CM': -100,
             'M': 1000}
        return sum(d.get(s[i: i + 2], d[s[i]]) for i in range(len(s)))
