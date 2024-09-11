package leetcode

import (
	"strconv"
	"strings"
)

func minBitFlips(start int, goal int) int {
	binS := strconv.FormatInt(int64(start), 2)
	binG := strconv.FormatInt(int64(goal), 2)
	maxLen := max(len(binS), len(binG))
	s := strings.Repeat("0", maxLen-len(binS)) + binS
	g := strings.Repeat("0", maxLen-len(binG)) + binG
	ans := 0
	for i := 0; i < maxLen; i++ {
		if s[i] != g[i] {
			ans += 1
		}
	}
	return ans
}