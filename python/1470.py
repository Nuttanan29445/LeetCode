class Solution:
    def shuffle(self,nums: list[int], n: int) -> list[int]:
        l = len(nums) / 2
        x = []
        y = []
        ans = []
        for i,j in enumerate(nums):
            if i >= l:
                y.append(j)
            else:
                x.append(j)
        for i in range(len(x)):
            ans.append(x[i])
            ans.append(y[i])
        return ans




        