class Solution:
    def plusOne(self,digits: list[int]) -> list[int]:
        z=[str(x) for x in digits]
        ans = ''.join(z)
        ans = int(ans)+1
        ans1 = [x for x in str(ans)]
        return ans1
        
 
    
        