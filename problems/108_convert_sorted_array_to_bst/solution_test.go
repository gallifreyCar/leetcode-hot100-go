package convertsortedarray

import "testing"

func TestSortedArrayToBST(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want bool // just check if it is a valid BST
	}{
		{
			name: "Example_1",
			nums: []int{-10, -3, 0, 5, 9},
			want: true,
		},
		{
			name: "Example_2",
			nums: []int{1, 3},
			want: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			root := sortedArrayToBST(tt.nums)
			if got := isValidBST(root, nil, nil); got != tt.want {
				t.Errorf("sortedArrayToBST() created invalid BST")
			}
		})
	}
}

func isValidBST(node *TreeNode, min, max *int) bool {
	if node == nil {
		return true
	}
	if min != nil && node.Val <= *min {
		return false
	}
	if max != nil && node.Val >= *max {
		return false
	}
	return isValidBST(node.Left, min, &node.Val) && isValidBST(node.Right, &node.Val, max)
}
