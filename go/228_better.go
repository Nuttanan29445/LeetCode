package leetcode

import "fmt"

func summaryRanges2(nums []int) []string {
    var ans []string
    i := 0
    for i < len(nums) {
        start := i
        
        for i+1 < len(nums) && nums[i+1] == nums[i]+1 {
            i++
        }
  
        if start == i {
           
            ans = append(ans, fmt.Sprintf("%v", nums[start]))
        } else {
           
            ans = append(ans, fmt.Sprintf("%v->%v", nums[start], nums[i]))
        }
        i++
    }
    return ans
}