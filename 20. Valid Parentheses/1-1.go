package main

import "fmt"

func main() {
	str := "()[]{}"
	fmt.Println(isValid(str))

	str = "([)]"
	fmt.Println(isValid(str))

}

func isValid(s string) bool {

	//括号集合
	brackets := map[rune]rune{')': '(', ']': '[', '}': '{'}

	var stack []rune

	for _, char := range s {
		if char == '(' || char == '[' || char == '{' { //匹配就入栈
			stack = append(stack, char)
		} else if brackets[char] == stack[len(stack)-1] { //如果有匹配，就出栈
			stack = stack[:len(stack)-1]
		} else {
			return false //如果不是指定的括号，那么就是错误的输入，返回false
		}
	}

	//栈全部清空为true
	return len(stack) == 0
}
