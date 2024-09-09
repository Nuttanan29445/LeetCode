package leetcode

import "sort"

func findContentChildren(g []int, s []int) int {
	sort.Ints(s)
	sort.Ints(g)
    i,j := 0,0
	ans := 0
	for i < len(g) && j < len(s){
        if s[j] >= g[i]{
            ans+=1
            i+=1
        }
        j+=1
    }
    return ans
}