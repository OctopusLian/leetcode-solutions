class Solution:
    def isOneBitCharacter(self, bits: List[int]) -> bool:
        return False if not bits or bits == [[1, 0]] else True if bits == [0] else self.isOneBitCharacter(bits[1:]) if bits[0] == 0 else self.isOneBitCharacter(bits[2:])