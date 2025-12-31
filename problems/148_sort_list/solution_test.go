package sortlist

import (
	"testing"
)

func TestSortList(t *testing.T) {
	tests := []struct {
		name   string
		values []int
		want   []int
	}{
		{
			name:   "Example_1",
			values: []int{4, 2, 1, 3},
			want:   []int{1, 2, 3, 4},
		},
		{
			name:   "Example_2",
			values: []int{-1, 5, 3, 4, 0},
			want:   []int{-1, 0, 3, 4, 5},
		},
		{
			name:   "Empty",
			values: []int{},
			want:   []int{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			head := buildList(tt.values)
			got := sortList(head)
			gotValues := listToSlice(got)
			if !equal(gotValues, tt.want) {
				t.Errorf("sortList() = %v, want %v", gotValues, tt.want)
			}
		})
	}
}

func buildList(values []int) *ListNode {
	if len(values) == 0 {
		return nil
	}
	head := &ListNode{Val: values[0]}
	cur := head
	for i := 1; i < len(values); i++ {
		cur.Next = &ListNode{Val: values[i]}
		cur = cur.Next
	}
	return head
}

func listToSlice(head *ListNode) []int {
	var result []int
	for head != nil {
		result = append(result, head.Val)
		head = head.Next
	}
	return result
}

func equal(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}
