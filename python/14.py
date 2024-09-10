class Solution:
    def longestCommonPrefix(self, strs: List[str]) -> str:
        s = sorted(strs)
        f = s[0]
        l = s[-1]
        ans = ''
        for i in range(min(len(f),len(l))):
            if f[i] != l[i]:
                return ans
            ans += f[i]
        return ans