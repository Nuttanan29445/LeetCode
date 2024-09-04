class Solution:
    def restoreString(self,s: str, indices: list[int]) -> str:
        z = {}
        ans = ''
        for i,j in enumerate(indices):
            z[j] = s[i]
        for i in sorted(z.keys()):
            ans+=z[i]
        return ans