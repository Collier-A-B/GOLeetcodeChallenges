package main

import "fmt"

type TestCase struct {
	nums     []int
	target   int
	expected int
}

func main() {
	fmt.Println("Search Insert Position Problem")

	var testCases = []TestCase{
		{[]int{1, 3, 5, 6}, 5, 2},
		{[]int{1, 3, 5, 6}, 2, 1},
		{[]int{1, 3, 5, 6}, 7, 4},
	}

	fmt.Println("Total Test Cases:", len(testCases))

	for i, tc := range testCases {
		result := searchInsert(tc.nums, tc.target)

		fmt.Println("\n*************************************************************")
		fmt.Println("\nTest Case:", i+1)

		fmt.Println("Input: nums =", tc.nums, ", target =", tc.target)
		fmt.Println("Expected Output:", tc.expected)
		fmt.Println("Actual Output:", result)

		if result == tc.expected {
			fmt.Printf("Test case %d PASSED\n", i+1)
		} else {
			fmt.Printf("Test case %d FAILED: expected %d, got %d\n", i+1, tc.expected, result)
		}
	}

}

func searchInsert(nums []int, target int) int {
	return -1 // TODO: Implement
}
