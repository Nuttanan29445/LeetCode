def relativeSortArray( arr1: list[int], arr2: list[int]) -> list[int]:
        ans = []
        for i in arr2:
            for j in range(arr1.count(i)):
                ans.append(i)
            arr1 = [s for s in arr1 if s != i]

        z = sorted(arr1)
        for i in z:
            ans.append(i)
        return ans
print(relativeSortArray(arr1 = [2,3,1,3,2,4,6,7,9,2,19], arr2 = [2,1,4,3,9,6]))