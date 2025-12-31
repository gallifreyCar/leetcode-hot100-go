package copylistwithrandompointer

type Node struct {
	Val    int
	Next   *Node
	Random *Node
}

// CopyRandomList 鍝堝笇琛?
// 鏃堕棿澶嶆潅搴? O(n)
// 绌洪棿澶嶆潅搴? O(n)
func CopyRandomList(head *Node) *Node {
	//鍏堢敓鎴愬搴旂殑鑺傜偣鏀惧叆鍝堝笇琛?
	cur := head
	m := make(map[*Node]*Node)
	for cur != nil {
		m[cur] = &Node{Val: cur.Val}
		cur = cur.Next
	}

	//鍐嶈祴鍊糿ext鍜宺andom
	cur = head
	for cur != nil {
		m[cur].Random = m[cur.Random]
		m[cur].Next = m[cur.Next]
		cur = cur.Next
	}

	return m[head]
}
