package main

import "fmt"

func main() {

	nums := []int{2, 7, 11, 15}
	target := 9

	fmt.Println(twoSum(nums, target))
}

func twoSum(nums []int, target int) []int {
	tmp := make([]int, 2)

	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {

			if nums[i]+nums[j] == target {
				tmp = []int{i, j}
				return tmp

			}
		}

	}
	return tmp
}
