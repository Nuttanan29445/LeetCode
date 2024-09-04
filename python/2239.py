class Solution:
    def findClosestNumber(self,nums: list[int]) -> int:
        n = min([abs(x) for x in nums])
        j = []
        for i in nums:
            if abs(i) == n:
                j.append(i)
        return max(j)
