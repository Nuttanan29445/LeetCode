package leetcode

func isPowerOfFour(n int) bool {
	if n == 1 {
		return true
	}
	for n%4 == 0 && n >= 1 {
		n = n / 4
		if n == 1 {
			return true
		}
	}
	return false
}