class Solution:
    def sortedSquares(self, nums: list[int]) -> list[int]:
        return sorted([int(x*x) for x in nums])
        