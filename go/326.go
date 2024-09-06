package leetcode

func isPowerOfThree(n int) bool {
	if n == 1 {
		return true
	}
	for n%3 == 0 && n >= 1 {
		n = n / 3
		if n == 1 {
			return true
		}
	}
	return false

}