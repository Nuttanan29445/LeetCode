package leetcode

func isUgly(n int) bool {
	for {
		if n == 0 {
			return false
		}
		if n == 1 {
			return true
		}
		if n%2 == 0 {
			n = n / 2
		} else if n%3 == 0 {
			n = n / 3
		} else if n%5 == 0 {
			n = n / 5
		} else {
			return false
		}
	}
}

func isUglyRecursive(n int) bool {
	for {
		if n == 0 {
			return false
		}
		if n == 1 {
			return true
		}

		if n%2 == 0 {
			return isUgly(n / 2)
		} else if n%3 == 0 {
			return isUgly(n / 3)
		} else if n%5 == 0 {
			return isUgly(n / 5)
		}
		return false
	}
}