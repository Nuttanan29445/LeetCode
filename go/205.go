package leetcode

import (
	"fmt"
	"strconv"
)

func isIsomorphic(s string, t string) bool {
	temp1 := make(map[string]int)
	var t1 []string

	temp2 := make(map[string]int)
	var t2 []string

	for i := range s {
		if !stringInMap(string(s[i]), temp1) {
			temp1[string(s[i])] = i
			t1 = append(t1, fmt.Sprintf("%v", i))
		} else {
			t1 = append(t1, strconv.Itoa(temp1[string(s[i])]))
		}

		if !stringInMap(string(t[i]), temp2) {
			temp2[string(t[i])] = i
			t2 = append(t2, fmt.Sprintf("%v", i))
		} else {
			t2 = append(t2, strconv.Itoa(temp2[string(s[i])]))
		}
	}

	
	for index, val := range t {
		if !stringInMap(string(val), temp2) {
			temp2[string(val)] = index
			t2 = append(t2, fmt.Sprintf("%v", index))
		} else {
			t2 = append(t2, strconv.Itoa(temp2[string(val)]))
		}
	}

	for i := range t1 {
		if t1[i] != t2[i] {
			return false
		}
	}

	return true
}

func stringInMap(a string, map1 map[string]int) bool {
    _, ok := map1[a]
    return ok
}