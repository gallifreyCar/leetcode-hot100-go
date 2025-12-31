package linkedlistcycle

type ListNode struct {
	Val  int
	Next *ListNode
}

// HasCycle 蹇參鎸囬拡瑙ｆ硶
// 鏃堕棿澶嶆潅搴? O(n)
// 绌洪棿澶嶆潅搴? O(1)
func HasCycle(head *ListNode) bool {
	//蹇參鎸囬拡
	fast := head
	slow := head
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
		if fast == slow {
			return true
		}
	}
	return false
}
