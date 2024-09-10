class Solution:
    def numberOfPoints(self, nums: List[List[int]]) -> int:
        z = []
        for i in nums:
            for j in range(i[0],i[1]+1):
                z.append(j)
        return len(set(z))