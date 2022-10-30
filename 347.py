from collections import defaultdict
class Solution:
    def topKFrequent(self,nums: list[int], k: int) -> list[int]:
        n = set(nums)
        z = sorted([nums.count(x) for x in n], reverse = True)
        z = z[:k]
        d = defaultdict(list)
        for i in n:
            if nums.count(i) in z:
                d[nums.count(i)].append(i)
        ans = []
        for i in sorted(set(z),reverse=True):
            ans.append(d[i])
        a = []
        for i in ans:
            for j in i:
                a.append(j)
        return a[:k]

        