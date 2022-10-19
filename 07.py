class Solution(object):
    def reverse(self,x):
        z = []
        temp = pow(2,31)
        for i in str(x):
            z.append(i)
        if x > 0:
            a = int("".join(z[-1::-1]))
            if a > temp or a < -temp:
                return 0
            else:
                return a
        else:
            j=[]
            j.append(z[0])
            for x in z[-1:0:-1]:
                j.append(x)
            a1 = int("".join(j))
            if a1 > temp or a1 < -temp:
                return 0
            else:
                return a1