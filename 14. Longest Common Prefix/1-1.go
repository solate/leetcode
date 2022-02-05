package main

import (
	"fmt"
	"strings"
)

func main() {

	arr := []string{"flower", "flow", "flight"}
	fmt.Println(longestCommonPrefix(arr))

	arr = []string{"dog", "racecar", "car"}
	fmt.Println(longestCommonPrefix(arr))

}

//水平扫描法
func longestCommonPrefix(strs []string) string {

	if len(strs) == 0 {
		return ""
	}

	prefix := strs[0]

	for i := 1; i < len(strs); i++ {
		for strings.Index(strs[i], prefix) != 0 {
			prefix = string(prefix[:len(prefix)-1])
			if prefix == "" {
				return ""
			}
		}
	}

	return prefix

}
