package leetcode

func getSmallestString(s string) string {

	runes := []rune(s)

	for i := 0; i < len(runes)-1; i++ {
		n1 := int64(runes[i] - '0')
		n2 := int64(runes[i+1] - '0')
		if n1 > n2 && ((n1 % 2) == (n2 % 2)) {
			runes[i], runes[i+1] = runes[i+1], runes[i]
			break
		}
	}
	return string(runes)
}