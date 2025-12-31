package sortlist

type ListNode struct {
	Val  int
	Next *ListNode
}

func sortList(head *ListNode) *ListNode {
	//閫掑綊缁撴潫鏉′欢
	if head == nil || head.Next == nil {
		return head
	}
	//1.鎵句腑鐐?
	mid := findMid(head) //鍋舵暟鍙栧乏杈癸紝濂囨暟鍙栦腑闂?
	right := mid.Next
	mid.Next = nil //鏂紑

	//2.褰掑苟
	left := sortList(head)
	right = sortList(right)

	//3.鍚堝苟
	return merge(left, right)

}

func findMid(head *ListNode) *ListNode {
	slow, fast := head, head.Next //fast鍜宻low璧风偣涓嶄竴鏍凤紝纭繚slow鍋滃湪宸﹁捣鐐?

	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}
	return slow
}

func merge(left *ListNode, right *ListNode) *ListNode {
	dummy := &ListNode{}
	cur := dummy
	for left != nil && right != nil {
		if left.Val < right.Val {
			cur.Next = left
			left = left.Next
		} else {
			cur.Next = right
			right = right.Next
		}
		cur = cur.Next
	}
	if left != nil {
		cur.Next = left
	}
	if right != nil {
		cur.Next = right
	}
	return dummy.Next
}
