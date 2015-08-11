package leetcode

import (
	"sort"
)

// https://leetcode.com/problems/two-sum/
// O(n)
func TwoSum(ints []int, sum int) [2]int {
	idxs := make([]int, 2)
	// key: value required
	// val: index of another part
	hash := map[int]int{}
	for i := range ints {
		hash[sum-ints[i]] = i + 1
	}
	for i := range ints {
		idx, ok := hash[ints[i]]
		if !ok {
			continue
		}
		idxs[0] = idx
		idxs[1] = i + 1
		sort.Ints(idxs)
		break
	}
	return [2]int{idxs[0], idxs[1]}
}
