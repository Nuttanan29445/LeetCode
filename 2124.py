class Solution:
    def checkString(self,s: str) -> bool:
        z = []
        for i,j in enumerate(s):
            if j == 'b':
                z = s[i:]
                break
        if z == []:
            return True
        else:
            return z.count('a') == 0