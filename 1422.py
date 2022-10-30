class Solution:
    def maxScore(self,s: str) -> int:
        j = []
        for i in range(1,len(s)):
            l = s[:i]
            r = s[i:]
            j.append(l.count('0')+r.count('1'))
        return max(j)