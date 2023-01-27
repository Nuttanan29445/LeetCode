def thirdMax(nums: list[int]) -> int:
        s = sorted(set(nums), reverse=True)
        print(s)
        if len(s) < 3:
            return (max(s))
        else:
            return s[2]
print(thirdMax(nums = [3,2,1]))
        