package leetcode

import (
	"math"
	"strconv"
)

func myAtoi(s string) int {
	t := ""
	zero := 0
	for _, val := range s {
		if val == 32 && zero == 0 && t == "" {
			continue
		}
		if val > 48 && val <= 57 {
			t += string(val)
			continue
		} else {
			if t == "" {
				if val == 48 {
					zero += 1
					continue
				} else if val == 43 && zero == 0 {
					zero += 1
					continue
				} else if val == 45 && zero == 0 {
					t += string(val)
					continue
				} else {
					return 0
				}
			} else if t == "-" {
				if val == 48 {
					zero += 1
					continue
				} else {
					break
				}
			} else {
				if val == 48 {
					zero += 1
					t += string(val)
				} else {
					break
				}

			}
		}
	}
	ans, _ := strconv.Atoi(t)
	max := int(math.Pow(2, 31)) - 1
	min := -int(math.Pow(2, 31))
	if ans < min {
		return min
	} else if ans > max {
		return max
	} else {
		return ans
	}
}

