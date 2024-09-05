package leetcode

import "fmt"

func summaryRanges(nums []int) []string {
	var ans []string
	var t []int
	temp := ""
	if len(nums) == 1 {
		ans = append(ans, fmt.Sprintf("%v", nums[0]))
	} else {
		for i := 0; i < len(nums)-1; i++ {
			if temp == "" && len(t) == 0 {
				temp = fmt.Sprintf("%v", nums[i])
				t = append(t, nums[i])
			}
			if nums[i+1]-nums[i] != 1 {
				if len(t) > 1 {
					ans = append(ans, fmt.Sprintf("%v->%v", temp, nums[i]))
				} else {
					ans = append(ans, fmt.Sprintf("%v", temp))
				}
				if i+1 == len(nums)-1 {
					ans = append(ans, fmt.Sprintf("%v", nums[i+1]))
				}
				temp = ""
				t = []int{}
			} else {
				if i+1 == len(nums)-1 {
					ans = append(ans, fmt.Sprintf("%v->%v", temp, nums[i+1]))
				}
				t = append(t, nums[i+1])
			}
		}
	}

	return ans
}