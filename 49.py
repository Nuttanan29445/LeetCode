from collections import defaultdict


class Solution:
    def groupAnagrams(self,strs: list[str]) -> list[list[str]]:
        z = []
        d = defaultdict(list)
        ans = []
        for i in strs:
            if sorted(i) not in z:
                z.append(sorted(i))
        for i,j in enumerate(strs):
            # print(sorted(j))
            if sorted(j) in z:
                d[str(sorted(j))].append(j)
        for i in z:
            ans.append(d[str(i)])
        return sorted(ans, key=len)