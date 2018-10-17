package main

import "fmt"

func main() {

	nums := []int{2, 7, 11, 15}
	target := 9

	fmt.Println(twoSum3(nums, target))
}

func twoSum3(nums []int, target int) []int {
	tmp := make([]int, 2)

	m := make(map[int]int)
	for k, v := range nums {
		complement := target - v
		if _, ok := m[complement]; ok && m[complement] != k {
			tmp = []int{m[complement], k}
			return tmp
		}
		m[v] = k
	}

	return tmp
}
