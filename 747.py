class Solution:
    def dominantIndex(self,nums: list[int]) -> int:
        m = max(nums)
        l = []
        for i,j in enumerate(nums):
            if j != m and j*2 > m:
                return -1
            if j == m:
                l.append(str(i))
        return "".join(l)