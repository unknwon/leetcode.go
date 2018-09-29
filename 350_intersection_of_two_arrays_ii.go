import "sort"

func intersect(nums1 []int, nums2 []int) []int {
	if len(nums1) == 0 || len(nums2) == 0 {
		return []int{}
	}

	sort.Ints(nums1)
	sort.Ints(nums2)

	nums := []int{}
	i1 := 0
	i2 := 0
	for i1 < len(nums1) && i2 < len(nums2) {
		if nums1[i1] == nums2[i2] {
			nums = append(nums, nums1[i1])
			i1++
			i2++
		} else if nums1[i1] > nums2[i2] {
			i2++
		} else {
			i1++
		}
	}
	return nums
}