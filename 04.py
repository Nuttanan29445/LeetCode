class Solution(object):
    def findMedianSortedArrays(nums1, nums2):
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