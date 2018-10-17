package main

import (
	"fmt"
	"strconv"
)

func main() {

	num := 121
	fmt.Println(isPalindrome3(num))

}

//是否是回文数
func isPalindrome3(x int) bool {

	if x < 0 || (x%10 == 0 && x != 0) {
		return false
	}

	charSlice := []byte(strconv.Itoa(x))
	length := len(charSlice)

	for i := 0; i < length/2; i++ {
		if charSlice[i] != charSlice[length-i-1] {
			return false
		}
	}

	return true
}
