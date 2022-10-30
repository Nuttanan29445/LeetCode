class Solution:
    def heightChecker(self,heights: list[int]) -> int:
        l = sorted(heights)
        c = 0
        for i in range(len(heights)):
            if l[i] != heights[i]:
                c+=1
        return c