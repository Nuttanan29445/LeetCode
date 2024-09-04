class Solution:
    def findErrorNums(self, nums: list[int]) -> list[int]:
        n = len(nums)
        a = sum(nums)
        b = sum(set(nums))
        s = n*(n+1)//2
        
        return [a-b, s-b]