package intersectionoftwolinkedlists

import "testing"

func TestGetIntersectionNode(t *testing.T) {
	// 创建相交的链表
	intersect := &ListNode{Val: 8, Next: &ListNode{Val: 4, Next: &ListNode{Val: 5}}}
	listA1 := &ListNode{Val: 4, Next: &ListNode{Val: 1, Next: intersect}}
	listB1 := &ListNode{Val: 5, Next: &ListNode{Val: 6, Next: &ListNode{Val: 1, Next: intersect}}}

	// 创建不相交的链表
	listA2 := &ListNode{Val: 1, Next: &ListNode{Val: 2}}
	listB2 := &ListNode{Val: 3, Next: &ListNode{Val: 4}}

	tests := []struct {
		name  string
		headA *ListNode
		headB *ListNode
		want  *ListNode
	}{
		{
			name:  "有相交节点",
			headA: listA1,
			headB: listB1,
			want:  intersect,
		},
		{
			name:  "无相交节点",
			headA: listA2,
			headB: listB2,
			want:  nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetIntersectionNode(tt.headA, tt.headB); got != tt.want {
				t.Errorf("GetIntersectionNode() = %v, want %v", got, tt.want)
			}
		})
	}
}
