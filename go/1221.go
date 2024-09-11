package leetcode

func balancedStringSplit(s string) int {
	m := map[string]int{
		"R": 0,
		"L": 0,
	}
	ans := 0
	for _, val := range s {
		m[string(val)] += 1

		if m["R"] == m["L"] {
			ans += 1
		}
	}
	return ans
}