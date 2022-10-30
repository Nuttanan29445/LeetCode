class Solution:
    def findMedianSortedArrays(self, nums1: list[int], nums2: list[int]) -> float:
        for i in nums2:
            nums1.append(i)
        z = sorted(nums1)
        valueMid = len(z)//2
        if len(z) % 2 == 0:
            f = (z[valueMid-1]+z[valueMid])/2
            return f
        else:
            f1 = z[valueMid]
            return f1