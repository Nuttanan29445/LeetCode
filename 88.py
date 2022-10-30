class Solution:
    def merge(self,nums1: list[int], m: int, nums2: list[int], n: int) -> None:
        """
        Do not return anything, modify nums1 in-place instead.
        """
        for i in range(len(nums1)-m):
            nums1.remove(nums1[m])
        for i in range(len(nums2)-n):
            nums2.remove(nums2[n])
        for i in nums2:
            nums1.append(i)
        nums1.sort()
    

   
        