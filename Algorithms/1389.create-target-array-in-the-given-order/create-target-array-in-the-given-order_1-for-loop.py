class Solution:
    def createTargetArray(self, nums: List[int], index: List[int]) -> List[int]:
        target=[0 for x in range(len(index))]
        for i in range(len(nums)):
            target.insert(index[i],nums[i])
        return target[:len(index):]