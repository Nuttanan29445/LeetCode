package leetcode

func checkTwoChessboards(coordinate1 string, coordinate2 string) bool {
	s1 := int64(coordinate1[0]) + int64(coordinate1[1])
	s2 := int64(coordinate2[0]) + int64(coordinate2[1])

	if ((s1%2 == 0) && (s2%2 == 0)) || ((s1%2 != 0) && (s2%2 != 0)) {
		return true
	} else {
		return false
	}
}