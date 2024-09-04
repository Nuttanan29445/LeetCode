class Solution:
    def singleNumber(self, nums: list[int]) -> int:
        z = set(nums)-set(nums-set(nums))
        for i in z:
            if nums.count(i) == 1:
                return i