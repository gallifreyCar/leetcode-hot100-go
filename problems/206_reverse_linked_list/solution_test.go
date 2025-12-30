package reverselinkedlist

import (
	"reflect"
	"testing"
)

func TestReverseList(t *testing.T) {
	tests := []struct {
		name string
		head *ListNode
		want []int
	}{
		{
			name: "示例1",
			head: &ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 3, Next: &ListNode{Val: 4, Next: &ListNode{Val: 5}}}}},
			want: []int{5, 4, 3, 2, 1},
		},
		{
			name: "空链表",
			head: nil,
			want: []int{},
		},
		{
			name: "单节点",
			head: &ListNode{Val: 1},
			want: []int{1},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := ReverseList(tt.head)
			result := []int{}
			for got != nil {
				result = append(result, got.Val)
				got = got.Next
			}
			if !reflect.DeepEqual(result, tt.want) {
				t.Errorf("ReverseList() = %v, want %v", result, tt.want)
			}
		})
	}
}
