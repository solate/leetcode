package leetcode

//给定一个只包括 '('，')'，'{'，'}'，'['，']' 的字符串 s ，判断字符串是否有效。
//
// 有效字符串需满足：
//
//
// 左括号必须用相同类型的右括号闭合。
// 左括号必须以正确的顺序闭合。
//
//
//
//
// 示例 1：
//
//
//输入：s = "()"
//输出：true
//
//
// 示例 2：
//
//
//输入：s = "()[]{}"
//输出：true
//
//
// 示例 3：
//
//
//输入：s = "(]"
//输出：false
//
//
// 示例 4：
//
//
//输入：s = "([)]"
//输出：false
//
//
// 示例 5：
//
//
//输入：s = "{[]}"
//输出：true
//
//
//
// 提示：
//
//
// 1 <= s.length <= 10⁴
// s 仅由括号 '()[]{}' 组成
//
// Related Topics 栈 字符串 👍 2888 👎 0

//leetcode submit region begin(Prohibit modification and deletion)
func isValid(s string) bool {

	n := len(s)
	if n%2 == 1 { // 奇数个直接返回
		return false
	}

	dict := map[byte]byte{
		')': '(',
		']': '[',
		'}': '{',
	}

	stack := []byte{}
	for i := 0; i < n; i++ {
		if dict[s[i]] > 0 { // 如果找到右括号
			stackLen := len(stack)
			if stackLen == 0 || stack[stackLen-1] != dict[s[i]] {
				return false
			}
			stack = stack[:stackLen-1]
		} else {
			stack = append(stack, s[i]) //左括号入栈
		}
	}

	return len(stack) == 0
}

//leetcode submit region end(Prohibit modification and deletion)
