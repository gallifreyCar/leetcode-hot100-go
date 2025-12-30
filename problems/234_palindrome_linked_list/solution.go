package palindromelinkedlist

// 234. 回文链表
// 难度：简单
// 标签：链表、双指针
// 链接：https://leetcode.cn/problems/palindrome_linked_list/

type ListNode struct {
	Val  int
	Next *ListNode
}

// IsPalindrome 双指针+反转后半部分
// 时间复杂度: O(n)
// 空间复杂度: O(1)
func IsPalindrome(head *ListNode) bool {
	if head == nil || head.Next == nil {
		return true
	}

	// 找到中点
	slow, fast := head, head
	for fast.Next != nil && fast.Next.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}

	// 反转后半部分
	secondHalf := reverseList(slow.Next)

	// 比较前后两部分
	firstHalf := head
	for secondHalf != nil {
		if firstHalf.Val != secondHalf.Val {
			return false
		}
		firstHalf = firstHalf.Next
		secondHalf = secondHalf.Next
	}

	return true
}

func reverseList(head *ListNode) *ListNode {
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
