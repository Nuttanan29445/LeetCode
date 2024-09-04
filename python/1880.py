def isSumEqual( firstWord: str, secondWord: str, targetWord: str) -> bool:
        f = []
        s = []
        t = []
        for i in firstWord:
            f.append(str(ord(i)-97))
        for j in secondWord:
            s.append(str(ord(j)-97))
        for k in targetWord:
            t.append(str(ord(k)-97))
        return int(''.join(f))+int(''.join(s)) == int(''.join(t))
print(isSumEqual(firstWord = "acb", secondWord = "cba", targetWord = "cdb"))
