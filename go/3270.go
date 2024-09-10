package leetcode

import (
	"strconv"
	"strings"
)

func generateKey(num1 int, num2 int, num3 int) int {
	n1 := strconv.Itoa(num1)
	n2 := strconv.Itoa(num2)
	n3 := strconv.Itoa(num3)

	maxlen := Max(len(n1), len(n2), len(n3))

	binN1 := strings.Repeat("0", maxlen-len(n1)) + n1
	binN2 := strings.Repeat("0", maxlen-len(n2)) + n2
	binN3 := strings.Repeat("0", maxlen-len(n3)) + n3
	ans := ""

	for i := 0; i < maxlen; i++ {
		b1, _ := strconv.Atoi(string(binN1[i]))
		b2, _ := strconv.Atoi(string(binN2[i]))
		b3, _ := strconv.Atoi(string(binN3[i]))
		ans += strconv.Itoa(min(b1, b2, b3))
	}
	result, _ := strconv.Atoi(ans)
	return result
}

func Max(number1, number2, number3 int) int {
	var localLargest int
	
	
	if number1 >= number2 && number1 >= number3 {
	   localLargest = number1
	} else if number2 >= number1 && number2 >= number3 {
	   localLargest = number2
	} else {
	   localLargest = number3
	}
	return localLargest
 }