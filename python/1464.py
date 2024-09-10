class Solution:
    def maxProduct(self, nums: List[int]) -> int:
        n = sorted(nums, reverse = True)
        return (n[0]-1)*(n[1]-1)