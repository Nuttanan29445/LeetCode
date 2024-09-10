class Solution:
    def findSpecialInteger(self, arr: List[int]) -> int:
        for i in arr:
            if arr.count(i) > (25/100)*len(arr):
                return i