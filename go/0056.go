package leetcode

import "sort"

func merge(intervals [][]int) [][]int {
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})
	var ans [][]int
	i := 0
	for i < len(intervals) {
		start := i
		valStart := intervals[start][0]
		valEnd := intervals[start][1]
		for i+1 < len(intervals) && (intervals[i+1][0] == valEnd || intervals[i+1][0] < valEnd) {
			valEnd = max(valEnd, intervals[i+1][1])
			i++
		}
		ans = append(ans, []int{valStart, valEnd})
		i++
	}
	return ans
}


func merge2(intervals [][]int) [][]int {
    if len(intervals) == 0 {
        return nil
    }

    sort.Slice(intervals, func(i, j int) bool {
        return intervals[i][0] < intervals[j][0]
    })

    var merged [][]int
    merged = append(merged, intervals[0])

    for i := 1; i < len(intervals); i++ {
        lastMerged := merged[len(merged)-1]

        if intervals[i][0] <= lastMerged[1] {
            lastMerged[1] = max(lastMerged[1], intervals[i][1])
        } else {
            merged = append(merged, intervals[i])
        }
    }

    return merged
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}