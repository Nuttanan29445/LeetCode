class Solution:
    def countEven(self,num: int) -> int:
        c = 0
        for i in range(1,num+1):
            a = list(str(i))
            b = sum([int(x) for x in a])
            if b % 2 == 0:
                c+=1
        return c