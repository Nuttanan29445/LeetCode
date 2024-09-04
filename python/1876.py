def countGoodSubstrings(s: str) -> int:
        t = []
        ans = []
        for i in range(len(s)-2):
            for k in s[i:i+3]:
                if k not in t:
                    t.append(k)
                else:
                    t = []
                    break
            if t!=[]:
                ans.append(''.join(t))
                t = []
        return ans
print(countGoodSubstrings(s="aababcabc"))