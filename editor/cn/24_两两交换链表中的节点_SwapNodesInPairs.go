package leetcode

//给你一个链表，两两交换其中相邻的节点，并返回交换后链表的头节点。你必须在不修改节点内部的值的情况下完成本题（即，只能进行节点交换）。
//
//
//
// 示例 1：
//
//
//输入：head = [1,2,3,4]
//输出：[2,1,4,3]
//
//
// 示例 2：
//
//
//输入：head = []
//输出：[]
//
//
// 示例 3：
//
//
//输入：head = [1]
//输出：[1]
//
//
//
//
// 提示：
//
//
// 链表中节点的数目在范围 [0, 100] 内
// 0 <= Node.val <= 100
//
// Related Topics 递归 链表 👍 1198 👎 0

//leetcode submit region begin(Prohibit modification and deletion)
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func swapPairs(head *ListNode) *ListNode {

	if head == nil || head.Next == nil {
		return head
	}

	newHead := head.Next                // 2  4
	head.Next = swapPairs(newHead.Next) // 3.next = 4  4.next = 3  1->2->3<->4
	newHead.Next = head
	return newHead
}

//leetcode submit region end(Prohibit modification and deletion)

type ListNode struct {
	Val  int
	Next *ListNode
}

func swapPairs1(head *ListNode) *ListNode {

	dummyHead := &ListNode{0, head}
	tmp := dummyHead
	for tmp.Next != nil && tmp.Next.Next != nil {
		node1 := tmp.Next
		node2 := tmp.Next.Next

		tmp.Next = node2
		node1.Next = node2.Next
		node2.Next = node1
		tmp = node1
	}

	return dummyHead.Next
}
