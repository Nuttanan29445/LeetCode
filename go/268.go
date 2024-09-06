package leetcode

func missingNumber(nums []int) int {
	sum := ((len(nums) + 1) * len(nums)) / 2
	s := 0

	for _, val := range nums {
		s += val
	}
	return sum - s
}