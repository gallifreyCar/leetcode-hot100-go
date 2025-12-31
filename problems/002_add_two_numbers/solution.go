package addtwonumbers

type ListNode struct {
	Val  int
	Next *ListNode
}

// AddTwoNumbers 妯℃嫙鍔犳硶
// 鏃堕棿澶嶆潅搴? O(max(m,n))
// 绌洪棿澶嶆潅搴? O(1)
func AddTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	cnt := 0 //杩涗綅
	res := &ListNode{}
	cur := res
	cur1 := l1
	cur2 := l2

	for cur1 != nil || cur2 != nil || cnt != 0 {
		sum := 0
		if cur1 != nil {
			sum += cur1.Val
			cur1 = cur1.Next
		}
		if cur2 != nil {
			sum += cur2.Val
			cur2 = cur2.Next
		}
		sum += cnt
		cnt = sum / 10
		cur.Next = &ListNode{Val: sum % 10}
		cur = cur.Next
	}

	return res.Next
}
