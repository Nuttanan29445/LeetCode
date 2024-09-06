package leetcode

import (
	"math"
	"strconv"
	"unicode"
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


//better solution
func myAtoi1(s string) int {
    i := 0
    n := len(s)
    
    // Step 1: Skip leading whitespaces
    for i < n && unicode.IsSpace(rune(s[i])) {
        i++
    }
    
    // Step 2: Check for optional sign
    sign := 1
    if i < n && (s[i] == '+' || s[i] == '-') {
        if s[i] == '-' {
            sign = -1
        }
        i++
    }
    
    // Step 3: Convert digits to integer
    result := 0
    for i < n && unicode.IsDigit(rune(s[i])) {
        digit := int(s[i] - '0')
        
        // Check for overflow
        if result > (math.MaxInt32-digit)/10 {
            if sign == 1 {
                return math.MaxInt32
            }
            return math.MinInt32
        }
        
        result = result*10 + digit
        i++
    }
    
    return sign * result
}

