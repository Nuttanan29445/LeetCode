class Solution:
    def maxSubsequence(self,nums: list[int], k: int) -> list[int]:
        n = sorted(nums, reverse = True)
        n = n[:k]
        ans = []
        c = 0
        for i in nums:
            if c == k:
                break
            else:
                if i in n:
                    ans.append(i)
                    n.remove(i)
                    c+=1
        return ans