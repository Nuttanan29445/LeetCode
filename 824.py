class Solution:
    def toGoatLatin(self,sentence: str) -> str:
        sentence = sentence.split()
        z=[]
        vowel = 'aeiouAEIOU'
        for num,i in enumerate(sentence):
            if i[0] in vowel:
                i += 'ma'+((num+1)*'a')
            else:
                i = i[1:]+i[0]+'ma'+((num+1)*'a')
            z.append(i)
        return " ".join(z)