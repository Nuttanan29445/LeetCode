class Solution:
    def minMaxGame(self,nums: list[int]) -> int:
        z = []
        z1 = [int(x) for x in nums]
        c = 0
        while(len(z1) != 1):
            for i in range(0,len(z1),2):
                c+=1
                if c % 2 !=0:
                    z.append(min(z1[i],z1[i+1]))
                else:
                    z.append(max(z1[i],z1[i+1]))
            z1 = z
            z = []
            c=0
        return z1[0]
    
    
            
        