class Solution:
    def isPerfectSquare(self, num: int) -> bool:
        if num**(1/2) == int(num**(1/2)):
            return True
        else:
            return False
        