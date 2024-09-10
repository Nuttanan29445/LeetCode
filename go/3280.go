package leetcode

import (
	"strconv"
	"strings"
)

func convertDateToBinary(date string) string {
	res := strings.Split(date, "-")
	var ans []string
	for _, val := range res {
		num, _ := strconv.Atoi(val)
		ans = append(ans, strconv.FormatInt(int64(num), 2))
	}
	return strings.Join(ans, "-")
}