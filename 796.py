class Solution:
    def rotateString(self, s: str, goal: str) -> bool:
        a1 = []
        a2 = []
        for i in s:
            a1.append(i)
        for i in goal:
            a2.append(i)
        for j,i in enumerate(a1):
            if a1 == a2:
                return True
            else:
                a1.remove(s[j])
                a1.append(s[j])
        return False