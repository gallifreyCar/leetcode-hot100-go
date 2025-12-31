package intersectionoftwolinkedlists

type ListNode struct {
	Val  int
	Next *ListNode
}

// GetIntersectionNode 鍙屾寚閽堣В娉?
// 鏃堕棿澶嶆潅搴? O(m+n)
// 绌洪棿澶嶆潅搴? O(1)
func GetIntersectionNode(headA, headB *ListNode) *ListNode {
	if headA == nil || headB == nil {
		return nil
	}
	curA := headA
	curB := headB
	for curB != curA {
		if curA != nil {
			curA = curA.Next
		} else {
			curA = headB
		}

		if curB != nil {
			curB = curB.Next
		} else {
			curB = headA
		}
	}
	return curA
}
