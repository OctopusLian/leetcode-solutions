class Solution(object):
    def trailingZeroes(self, n):
        """
        :type n: int
        :rtype: int
        """
        if n == 0:
            return 0
        elif n < 0:
            return -1
        else:
            return n / 5 + self.trailingZeroes(n / 5)
