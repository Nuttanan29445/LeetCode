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

func isPowerOfFourO1(n int) bool {
	//note 0x55555555 the only set bit is in an even position.
	return n > 0 && (n&(n-1)) == 0 && (n&0x55555555) != 0
}