class Solution(object):
    def topKFrequent(self,words, k):
        words = sorted(words)
        countW=[]
        ans=[]
        for i in words:
            countW.append(words.count(i))
        sortZ = sorted(countW,reverse=True)
        for i in sortZ:
            for m in words:
                if words.count(m) == i:
                    if len(ans) < k:
                        if m not in ans:
                            ans.append(m)
                    else:
                        break
        return ans