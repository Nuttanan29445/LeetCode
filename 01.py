class Solution(object):
    def twoSum(self, nums, target):
        z = []
        for j,i in enumerate(nums):
            v = j
            for k in nums[j+1:len(nums)]:
                v+=1
                if int(i)+int(k) == target:
                    return[j,v]