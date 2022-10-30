from collections import defaultdict
class Solution:
    def countPoints(self,rings: str) -> int:
        d = defaultdict(list)
        temp = sorted(['R','G','B'])
        c = 0
        for i,j in enumerate(rings) :
            if j.isnumeric():
                d[j].append(rings[i-1])
        for i in d.keys():
            if sorted(set(d[i])) == temp:
                c+=1
        return c