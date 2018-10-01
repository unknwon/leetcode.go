package leetcode

func twoSum(nums []int, target int) []int {
	m := make(map[int]int) // num -> index
	for i := range nums {
		left := target - nums[i]
		if idx, ok := m[left]; ok {
			return []int{i, idx}
		}
		m[nums[i]] = i
	}
	return []int{-1, -1}
}
