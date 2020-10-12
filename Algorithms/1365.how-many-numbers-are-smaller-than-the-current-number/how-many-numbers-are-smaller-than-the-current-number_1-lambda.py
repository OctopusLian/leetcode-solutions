class Solution:
    def smallerNumbersThanCurrent(self, nums: List[int]) -> List[int]:
        return [len(list(filter(lambda x: x < i, nums))) for i in nums]