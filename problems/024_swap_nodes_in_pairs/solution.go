package swapnodesinpairs

type ListNode struct {
	Val  int
	Next *ListNode
}

// SwapPairs 杩唬娉?
// 鏃堕棿澶嶆潅搴? O(n)
// 绌洪棿澶嶆潅搴? O(1)
func SwapPairs(head *ListNode) *ListNode {
	dummy := &ListNode{Next: head}
	cur := dummy
	for cur.Next != nil && cur.Next.Next != nil {
		tmp := cur.Next           //1
		tmp2 := tmp.Next          //2
		tmp3 := tmp2.Next         //3
		cur.Next = tmp2           //0->2
		cur.Next.Next = tmp       //2->1
		cur.Next.Next.Next = tmp3 //1->3

		cur = cur.Next.Next //绉诲姩鍒颁笅涓€瀵?
	}

	return dummy.Next
}
