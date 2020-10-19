class Solution:
    def isOneBitCharacter(self, bits: List[int]) -> bool:
        for inx, i in enumerate(bits):
            if i:
                bits[inx] = bits[inx+1] = None
        return bits[-1] == 0