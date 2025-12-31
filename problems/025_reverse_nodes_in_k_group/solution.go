package reversenodesingroup

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseList(head *ListNode) *ListNode {
	cur := head
	var pre *ListNode
	for cur != nil {
		tmp := cur.Next
		cur.Next = pre
		pre = cur
		cur = tmp
	}
	return pre
}

// ReverseKGroup K涓竴缁勭炕杞摼琛?
// 鏃堕棿澶嶆潅搴? O(n)
// 绌洪棿澶嶆潅搴? O(1)
func ReverseKGroup(head *ListNode, k int) *ListNode {
	dummy := &ListNode{Next: head}
	start := head //缈昏浆鍚庡ご灏句細鍙嶈浆
	cur := dummy  //鐢ㄤ簬閾炬帴鍙嶈浆鍚庣殑閾捐〃
	end := dummy  //end浠庡墠涓€涓妭鐐瑰嚭鍙?
	for {
		// end鍏堢Щ鍔ㄥ埌k
		for i := 0; i < k; i++ {
			if end.Next == nil {
				return dummy.Next
			}
			end = end.Next
		}
		nextGroup := end.Next         //涓嬩竴缁?
		end.Next = nil                //鏂紑
		newHead := reverseList(start) //缈昏浆閾捐〃
		cur.Next = newHead
		cur = start            //cur鍥炲埌闃熷熬甯繖閾炬帴
		end = start            //end鍥炲埌闃熷熬閲嶆柊閬嶅巻
		start.Next = nextGroup //鍐嶉摼鎺ヤ笅涓€缁?
		start = nextGroup
	}
}
