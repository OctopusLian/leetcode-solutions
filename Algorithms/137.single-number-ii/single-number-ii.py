class Solution(object):
    def singleNumber(self, nums):
        """
        :type nums: List[int]
        :rtype: int
        """
        if nums is None:
            return 0

        result = 0
        for i in xrange(32):
            bit_i_sum = 0
            for num in nums:
                bit_i_sum += ((num >> i) & 1)
            result |= ((bit_i_sum % 3) << i)
        return self.twos_comp(result, 32)
    
    def twos_comp(self, val, bits):
        """
        compute the 2's compliment of int value val
        e.g. -4 ==> 11100 == -(10000) + 01100
        """
        return -(val & (1 << (bits - 1))) | (val & ((1 << (bits - 1)) - 1))
