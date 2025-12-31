package linkedlistcycleii

type ListNode struct {
	Val  int
	Next *ListNode
}

// DetectCycle 蹇參鎸囬拡
// 缁撹锛氳捣鐐瑰埌鐜叆鍙ｈ窛绂?鐩搁亣鐐瑰埌鐜叆鍙ｈ窛绂?
// 鏃堕棿澶嶆潅搴? O(n)
// 绌洪棿澶嶆潅搴? O(1)
func DetectCycle(head *ListNode) *ListNode {
	fast := head
	slow := head
	var meetPoint *ListNode
	//鎵惧埌鐩搁亣鐐?
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
		if fast == slow {
			meetPoint = fast
			break
		}
	}
	if meetPoint == nil {
		return nil
	}

	// 鍐嶈蛋涓€娆?
	cur := head
	for slow != cur {
		slow = slow.Next
		cur = cur.Next
	}
	return cur
}
