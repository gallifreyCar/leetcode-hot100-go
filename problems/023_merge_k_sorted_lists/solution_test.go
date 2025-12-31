package mergeksortedlists

import "testing"

func TestMergeKLists(t *testing.T) {
	tests := []struct {
		name   string
		lists  [][]int
		want   []int
	}{
		{
			name:   "Example_1",
			lists:  [][]int{{1, 4, 5}, {1, 3, 4}, {2, 6}},
			want:   []int{1, 1, 2, 3, 4, 4, 5, 6},
		},
		{
			name:   "Empty",
			lists:  [][]int{},
			want:   []int{},
		},
		{
			name:   "EmptyLists",
			lists:  [][]int{{}},
			want:   []int{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			lists := make([]*ListNode, len(tt.lists))
			for i, vals := range tt.lists {
				lists[i] = buildList(vals)
			}
			got := mergeKLists(lists)
			gotValues := listToSlice(got)
			if !equal(gotValues, tt.want) {
				t.Errorf("mergeKLists() = %v, want %v", gotValues, tt.want)
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
