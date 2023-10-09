package median_of_two_sorted_arrays

import (
	"slices"
)

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	merged := append(nums1, nums2...)
	slices.Sort(merged)

	length := len(merged)

	if length%2 != 0 {
		return float64(merged[length/2])
	} else {
		a := float64(merged[length/2])
		b := float64(merged[length/2-1])

		return (a + b) / 2
	}
}
