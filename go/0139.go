package leetcode

func wordBreak(s string, wordDict []string) bool {
	wordSet := make(map[string]struct{})
	for _, word := range wordDict {
		wordSet[word] = struct{}{}
	}

	n := len(s)
	memo := make([]bool, n+1)
	memo[0] = true

	for end := 1; end <= n; end++ {
		for start := 0; start < end; start++ {
			if memo[start] {
				if _, exists := wordSet[s[start:end]]; exists {
					memo[end] = true
					break
				}
			}
		}
	}
	return memo[n]
}