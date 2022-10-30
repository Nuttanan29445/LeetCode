class Solution:
    def countWords(self, words1: list[str], words2: list[str]) -> int:
        c=0
        for i in words1:
            if i in words2 and words1.count(i) == 1 and words2.count(i) == 1:
                c+=1
        return c
   
