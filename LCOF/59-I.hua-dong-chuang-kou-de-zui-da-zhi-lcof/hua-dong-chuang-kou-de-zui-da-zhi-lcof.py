class Solution:
    def maxSlidingWindow(self, nums: List[int], k: int) -> List[int]:
        res = []
        if nums != []:
            tmp_max = max(nums[0:k])
            for i in range(len(nums) - (k-1)):
                tmp_max = max(nums[i:i+k])
                res.append(tmp_max)
            return res
        else:
            return []
