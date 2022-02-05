package main

import "fmt"

func main() {

	arr := []string{"flower", "flow", "flight"}
	fmt.Println(longestCommonPrefix(arr))

	arr = []string{"dog", "racecar", "car"}
	fmt.Println(longestCommonPrefix(arr))

}

//水平扫描
func longestCommonPrefix(strs []string) string {

	length := len(strs)

	if length == 0 {
		return ""
	}

	str0Len := len(strs[0])

	//第一个字符串长度挨个遍历
	for i := 0; i < str0Len; i++ {
		char := strs[0][i]
		for j := 1; j < length; j++ {
			//数组中每个都取和第一个下标相同的进行比较，如果相等就跳过，比较下一个字母

			// i==len(strs[j]) 到了字符串的最后, 或者有元素的字母不相等，那么前缀就是i下标之前的字符串
			if i == len(strs[j]) || strs[j][i] != char {
				return strs[0][:i]
			}
		}
	}

	return strs[0]

}
