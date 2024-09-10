class Solution:
    def countBits(self, n: int) -> List[int]:
        z = []
        j = ''
        for i in range(n+1):
            j = bin(i).replace('0b','')
            sum = 0
            for k in j:
                sum += int(k)
            z.append(sum)
        return z
