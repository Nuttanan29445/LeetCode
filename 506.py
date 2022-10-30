class Solution:
    def findRelativeRanks(self,score: list[int]) -> list[str]:
        ss  = sorted(score, reverse=True)
        z = {}
        ans = []
        r = ["Gold Medal","Silver Medal","Bronze Medal"]
        for n,j in enumerate(ss):
            if n < 3:
                z[j] = r[n]
            else:
                z[j] = str(n+1)
        for i in score:
            ans.append(z[i])
        return ans