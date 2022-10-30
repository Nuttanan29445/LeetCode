class Solution:
    def findDifference(self,nums1: list[int], nums2: list[int]) -> list[list[int]]:
        a1 = []
        a2 = []
        for i in nums1:
            if i not in nums2 and i not in a1:
                a1.append(i)
        for i in nums2:
            if i not in nums1 and i not in a2:
                a2.append(i)
        return list([a1]+[a2])