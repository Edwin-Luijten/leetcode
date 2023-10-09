package median_of_two_sorted_arrays

import (
	"reflect"
	"testing"
)

type TestCase struct {
	nums1    []int
	nums2    []int
	expected float64
}

var tests []TestCase = []TestCase{
	{nums1: []int{1, 3}, nums2: []int{2}, expected: 2},
	{nums1: []int{1, 2}, nums2: []int{3, 4}, expected: 2.5},
}

var benchmark = TestCase{nums1: []int{1, 3}, nums2: []int{2}, expected: 2}

func TestFindMedianSortedArrays(t *testing.T) {
	for _, test := range tests {
		actual := findMedianSortedArrays(test.nums1, test.nums2)

		if !reflect.DeepEqual(test.expected, actual) {
			t.Errorf("findMedianSortedArrays(%v, %v) = %v; expected %v", test.nums1, test.nums2, actual, test.expected)
		}
	}
}

func BenchmarkTestFindMedianSortedArrays(b *testing.B) {
	for n := 0; n < b.N; n++ {
		findMedianSortedArrays(benchmark.nums1, benchmark.nums2)
	}
}
