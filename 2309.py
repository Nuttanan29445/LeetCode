class Solution:
    def greatestLetter(self,s: str) -> str:
        z = []
        for i in s:
            if i.islower():
                if i.upper() in s:
                    z.append(i.upper())
            else:
                if i.lower() in s:
                    z.append(i.upper()) 
        if z == []:
            return ""
        else:
            z = set(z)
            ans = max(z)
            return ans
   
