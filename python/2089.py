class Solution:
    def targetIndices(self, nums: list[int], target: int) -> list[int]:
        nums = sorted(nums)
        ans = []
        for i,j in enumerate(nums):
            if j == target:
                ans.append(i)
        return ans
