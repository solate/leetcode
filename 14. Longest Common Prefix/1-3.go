package main

import "fmt"

func main() {

	arr := []string{"flower", "flow", "flight", "flggg", "fllll"}
	fmt.Println(longestCommonPrefix(arr))

	arr = []string{"dog", "racecar", "car"}
	fmt.Println(longestCommonPrefix(arr))

}

//分治
func longestCommonPrefix(strs []string) string {

	if len(strs) == 0 {
		return ""
	}

	return longestCommonPrefixDivide(strs, 0, len(strs)-1)

}

func longestCommonPrefixDivide(strs []string, l, r int) string {

	if l == r { //如果相等，那么就拆到了最小
		return strs[l]
	}

	mid := (l + r) / 2 //中间值
	lcpLeft := longestCommonPrefixDivide(strs, 1, mid)
	lcpRight := longestCommonPrefixDivide(strs, mid+1, r)
	return commonPrefix(lcpLeft, lcpRight)

}

func commonPrefix(left, right string) string {

	min := 0
	if len(left) >= len(right) {
		min = len(right)
	} else {
		min = len(left)
	}

	for i := 0; i < min; i++ {
		if left[i] != right[i] {
			return left[:i]
		}
	}

	return left[:min]

}
