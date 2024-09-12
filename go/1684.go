package leetcode

import "strings"

func countConsistentStrings(allowed string, words []string) int {
	ans := 0
	for _, val := range words {
		res := val
		for _, allow := range allowed {
			res = strings.ReplaceAll(res, string(allow), "")
		}
		if res == "" {
			ans += 1
		}
	}
	return ans
}