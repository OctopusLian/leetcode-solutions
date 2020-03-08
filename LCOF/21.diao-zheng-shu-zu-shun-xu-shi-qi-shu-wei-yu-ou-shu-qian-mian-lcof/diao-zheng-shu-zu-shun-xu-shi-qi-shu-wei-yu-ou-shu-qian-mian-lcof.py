class Solution:
    def exchange(self, nums: List[int]) -> List[int]:
        return sorted(nums,key=lambda x:1-x%2)
