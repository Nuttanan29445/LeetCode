class Solution:
    def decrypt(self, code: List[int], k: int) -> List[int]:
        ans = []
        t = 0
        if k > 0:
            for i in range(-len(code),0):
                for j in range(1,k+1):
                    t += code[i+j]
                ans.append(t)
                t = 0
            return ans
        elif k == 0:
            for i in range(len(code)):
                ans.append(0)
            return ans
        elif k < 0:
            for i in range(len(code)): 
                for j in range(1,-k+1): 
                    t += code[i-j]
                    print(t)
                ans.append(t)
                t = 0
            return ans
            