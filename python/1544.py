def makeGood(s) -> str:
        st = []
        for i in range(len(s)):
            if st and abs(ord(st[-1]) - ord(s[i])) == 32:
                st.pop(-1)
            else:
                st.append(s[i])  
        return "".join(st)
            
print(makeGood("Pp"))