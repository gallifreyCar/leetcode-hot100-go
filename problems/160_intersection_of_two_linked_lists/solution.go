package intersectionoftwolinkedlists

// 160. 相交链表
// 难度：简单
// 标签：链表、双指针
// 链接：https://leetcode.cn/problems/intersection_of_two_linked_lists/

type ListNode struct {
	Val  int
	Next *ListNode
}

// GetIntersectionNode 双指针
// 时间复杂度: O(m+n)
// 空间复杂度: O(1)
func GetIntersectionNode(headA, headB *ListNode) *ListNode {
	if headA == nil || headB == nil {
		return nil
	}

	pA, pB := headA, headB
	// 当pA和pB相遇时，就是相交节点
	for pA != pB {
		if pA == nil {
			pA = headB
		} else {
			pA = pA.Next
		}
		if pB == nil {
			pB = headA
		} else {
			pB = pB.Next
		}
	}
	return pA
}
