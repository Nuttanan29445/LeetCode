package leetcode

import "sort"

func lastStoneWeight(stones []int) int {

	for len(stones) > 1 {
		sort.Slice(stones, func(i, j int) bool { return stones[i] > stones[j] })
		max1, max2 := stones[0], stones[1]
		if max1 == max2 {
			stones = append(stones[:0], stones[2:]...)
		} else if max1 != max2 {
			stones[0] = stones[0] - stones[1]
			stones = append(stones[:1], stones[2:]...)
		}
	}

	if len(stones) == 1 {
		return stones[0]
	} else {
		return 0
	}

}