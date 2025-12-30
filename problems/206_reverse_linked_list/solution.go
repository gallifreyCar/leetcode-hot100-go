package reverselinkedlist

// 206. 反转链表
// 难度：简单
// 标签：链表、递归
// 链接：https://leetcode.cn/problems/reverse_linked_list/

type ListNode struct {
	Val  int
	Next *ListNode
}

// ReverseList 迭代
// 时间复杂度: O(n)
// 空间复杂度: O(1)
func ReverseList(head *ListNode) *ListNode {
	var prev *ListNode
	current := head

	for current != nil {
		next := current.Next
		current.Next = prev
		prev = current
		current = next
	}

	return prev
}
