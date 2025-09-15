package main

import (
	"fmt"
	"reflect"
)

type TestCase struct {
	// first sorted array + its size
	nums1 []int
	m     int

	// second sorted array + its size
	nums2 []int
	n     int

	// expected merged array
	expected []int

	// result array, initially just a copy of nums1
	nums1Mutable []int
	nums2Mutable []int
}

func main() {
	fmt.Println("Search Insert Position Problem")

	var testCases = []TestCase{
		{[]int{1, 2, 3, 0, 0, 0}, 3, []int{2, 5, 6}, 3, []int{1, 2, 2, 3, 5, 6}, []int{1, 2, 3, 0, 0, 0}, []int{2, 5, 6}},
		{[]int{1}, 1, []int{}, 0, []int{1}, []int{1}, []int{}},
		{[]int{0}, 0, []int{1}, 1, []int{1}, []int{0}, []int{1}},
	}

	fmt.Println("Total Test Cases:", len(testCases))

	for i, tc := range testCases {

		fmt.Println("\n*************************************************************")
		fmt.Println("\nTest Case:", i+1)
		merge(tc.nums1Mutable, tc.m, tc.nums2Mutable, tc.n)

		fmt.Println("\nNums1: ", tc.nums1)
		fmt.Println("\nNums2: ", tc.nums2)
		fmt.Println("\nExpected: ", tc.expected)
		fmt.Println("\nResult: ", tc.nums1Mutable)

		if reflect.DeepEqual(tc.nums1Mutable, tc.expected) {
			fmt.Printf("\nTest case %d PASSED\n", i+1)
		} else {
			fmt.Printf("\nTest case %d FAILED", i+1)
		}
	}

}

func merge(nums1 []int, m int, nums2 []int, n int) {
	// TODO: IMplement
}
