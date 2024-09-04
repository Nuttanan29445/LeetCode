package leetcode

func getFinalState(nums []int, k int, multiplier int) []int {
	for i := 0; i < k; i++ {
		minimum := findMinElement(nums)
		for i, val := range nums {
			if val == minimum {
				nums[i] = val * multiplier
				break
			}
		}
	}
	return nums
}
func findMinElement(arr []int) int {
	min_num := arr[0]

	for i := 0; i < len(arr); i++ {
		if arr[i] < min_num {
			min_num = arr[i]
		}
	}
	return min_num
}