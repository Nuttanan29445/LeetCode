class Solution:
    def maximum69Number (self,num: int) -> int:
        num = list(str(num))
        for n,j in enumerate(num):
            if j == '6':
                num[n] = '9'
                break
        return "".join(num)