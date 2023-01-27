def pivotInteger(n: int) -> int:
        s = [int(i) for i in range(1,n+1)]
        for i in range(len(s)):
            if sum(s[0:i+1]) == sum(s[i:]):
                return i+1
        return -1
print(pivotInteger(n = 8))
