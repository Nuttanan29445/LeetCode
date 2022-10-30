class Solution:
    def isMonotonic(self, nums: list[int]) -> bool:
        if nums[0] == min(nums):
            z = sorted(nums)
        else:
            z = sorted(nums,reverse=True)
        return z == nums