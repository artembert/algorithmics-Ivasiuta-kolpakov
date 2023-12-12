package main

import "fmt"

func RemoveDuplicates(nums []int) int {
	pivot := 0

	for i := 0; i < len(nums)-1; i++ {
		if nums[i] != nums[i+1] {
			nums[pivot] = nums[i]
			pivot++
		}
	}

	return pivot + 1
}

func main() {
	nums := []int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}
	fmt.Print(RemoveDuplicates(nums))
}
