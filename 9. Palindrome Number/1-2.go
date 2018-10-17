package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {

	num := 121
	fmt.Println(isPalindrome2(num))

}

//是否是回文数
func isPalindrome2(x int) bool {

	if x < 0 || (x%10 == 0 && x != 0) {
		return false
	}

	str := strconv.Itoa(x)
	rev := reverseString(str)
	if strings.EqualFold(str, rev) { //反转字符串
		return true
	}

	return false
}

// 反转字符串
func reverseString(s string) string {
	runes := []rune(s)
	for from, to := 0, len(runes)-1; from < to; from, to = from+1, to-1 {
		runes[from], runes[to] = runes[to], runes[from]
	}
	return string(runes)
}
