class Solution(object):
    def addDigits(self,num):
        """
        :type num: int
        :rtype: int
        """
        z = []
        for i in str(num):
            z.append(int(i))
        j = sum(z)
        z = []
        while(j > 9):
            for i in str(j):
                z.append(int(i))
            j = sum(z)
            z = []
        return j