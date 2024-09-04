class Solution:
    def getLucky(self,s: str, k: int) -> int:
        j = ""
        c = 0
        for i in s:
            j+=str(ord(i)-96)
        print(j)
        a = []
        while c < k:
            for i in str(j):
                a.append(int(i))
            j = sum(a)
            a = []
            c+=1
        return j