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

func isPowerOfThree1(n int) bool {
	// The largest power of 3 within the 32-bit signed integer range is 3^19 = 1162261467
	return n > 0 && 1162261467%n == 0
}