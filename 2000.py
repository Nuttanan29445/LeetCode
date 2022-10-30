class Solution:
    def reversePrefix(self,word: str, ch: str) -> str:
        z = []
        c = True
        for i in word:
            if i == ch and c:
                z.append(i)
                z = z[::-1]
                c = False
            else:
                z.append(i)
        return "".join(z)