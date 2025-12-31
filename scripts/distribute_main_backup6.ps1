# 分发更多题解 - 二叉树

# 094_binary_tree_inorder_traversal
$content = @'
package inordertraversal

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func inorderTraversal(root *TreeNode) []int {
	res := make([]int, 0)
	var dfs func(node *TreeNode)
	dfs = func(node *TreeNode) {
		if node == nil {
			return
		}
		dfs(node.Left)
		res = append(res, node.Val)
		dfs(node.Right)
	}

	dfs(root)
	return res
}
'@
Set-Content -Path "problems/094_binary_tree_inorder_traversal/solution.go" -Value $content -Encoding UTF8
Write-Host "Created/Updated problems/094_binary_tree_inorder_traversal/solution.go"

$testContent = @'
package inordertraversal

import (
	"reflect"
	"testing"
)

func TestInorderTraversal(t *testing.T) {
	tests := []struct {
		name string
		root *TreeNode
		want []int
	}{
		{
			name: "Example_1",
			root: &TreeNode{Val: 1, Right: &TreeNode{Val: 2, Left: &TreeNode{Val: 3}}},
			want: []int{1, 3, 2},
		},
		{
			name: "Empty",
			root: nil,
			want: []int{},
		},
		{
			name: "Single",
			root: &TreeNode{Val: 1},
			want: []int{1},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := inorderTraversal(tt.root); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("inorderTraversal() = %v, want %v", got, tt.want)
			}
		})
	}
}
'@
Set-Content -Path "problems/094_binary_tree_inorder_traversal/solution_test.go" -Value $testContent -Encoding UTF8
Write-Host "Created/Updated problems/094_binary_tree_inorder_traversal/solution_test.go"

# 104_maximum_depth_of_binary_tree
$content = @'
package maximumdepth

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	left := maxDepth(root.Left)
	right := maxDepth(root.Right)
	return max(left, right) + 1
}
'@
Set-Content -Path "problems/104_maximum_depth_of_binary_tree/solution.go" -Value $content -Encoding UTF8
Write-Host "Created/Updated problems/104_maximum_depth_of_binary_tree/solution.go"

$testContent = @'
package maximumdepth

import "testing"

func TestMaxDepth(t *testing.T) {
	tests := []struct {
		name string
		root *TreeNode
		want int
	}{
		{
			name: "Example_1",
			root: &TreeNode{
				Val:  3,
				Left: &TreeNode{Val: 9},
				Right: &TreeNode{
					Val:   20,
					Left:  &TreeNode{Val: 15},
					Right: &TreeNode{Val: 7},
				},
			},
			want: 3,
		},
		{
			name: "Example_2",
			root: &TreeNode{Val: 1, Right: &TreeNode{Val: 2}},
			want: 2,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxDepth(tt.root); got != tt.want {
				t.Errorf("maxDepth() = %v, want %v", got, tt.want)
			}
		})
	}
}
'@
Set-Content -Path "problems/104_maximum_depth_of_binary_tree/solution_test.go" -Value $testContent -Encoding UTF8
Write-Host "Created/Updated problems/104_maximum_depth_of_binary_tree/solution_test.go"

# 101_symmetric_tree
$content = @'
package symmetrictree

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isSymmetric(root *TreeNode) bool {
	if root == nil {
		return true
	}
	var isMirror func(left, right *TreeNode) bool
	isMirror = func(left, right *TreeNode) bool {
		if left == nil && right == nil {
			return true
		}
		if left == nil || right == nil {
			return false
		}
		if left.Val != right.Val {
			return false
		}
		//外面和外面比 && 里面和里面比
		return isMirror(left.Left, right.Right) && isMirror(left.Right, right.Left)
	}
	return isMirror(root.Left, root.Right)
}
'@
Set-Content -Path "problems/101_symmetric_tree/solution.go" -Value $content -Encoding UTF8
Write-Host "Created/Updated problems/101_symmetric_tree/solution.go"

$testContent = @'
package symmetrictree

import "testing"

func TestIsSymmetric(t *testing.T) {
	tests := []struct {
		name string
		root *TreeNode
		want bool
	}{
		{
			name: "Example_1",
			root: &TreeNode{
				Val:   1,
				Left:  &TreeNode{Val: 2, Left: &TreeNode{Val: 3}, Right: &TreeNode{Val: 4}},
				Right: &TreeNode{Val: 2, Left: &TreeNode{Val: 4}, Right: &TreeNode{Val: 3}},
			},
			want: true,
		},
		{
			name: "Example_2",
			root: &TreeNode{
				Val:   1,
				Left:  &TreeNode{Val: 2, Right: &TreeNode{Val: 3}},
				Right: &TreeNode{Val: 2, Right: &TreeNode{Val: 3}},
			},
			want: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isSymmetric(tt.root); got != tt.want {
				t.Errorf("isSymmetric() = %v, want %v", got, tt.want)
			}
		})
	}
}
'@
Set-Content -Path "problems/101_symmetric_tree/solution_test.go" -Value $testContent -Encoding UTF8
Write-Host "Created/Updated problems/101_symmetric_tree/solution_test.go"

# 108_convert_sorted_array_to_bst
$content = @'
package convertsortedarray

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func sortedArrayToBST(nums []int) *TreeNode {
	if len(nums) == 0 {
		return nil
	}

	mid := len(nums) / 2
	root := &TreeNode{Val: nums[mid],
		Left:  sortedArrayToBST(nums[:mid]),
		Right: sortedArrayToBST(nums[mid+1:]),
	}
	return root
}
'@
Set-Content -Path "problems/108_convert_sorted_array_to_bst/solution.go" -Value $content -Encoding UTF8
Write-Host "Created/Updated problems/108_convert_sorted_array_to_bst/solution.go"

$testContent = @'
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
'@
Set-Content -Path "problems/108_convert_sorted_array_to_bst/solution_test.go" -Value $testContent -Encoding UTF8
Write-Host "Created/Updated problems/108_convert_sorted_array_to_bst/solution_test.go"

# 098_validate_binary_search_tree
$content = @'
package validatebst

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isValidBST(root *TreeNode) bool {
	var pre *TreeNode
	var inorder func(node *TreeNode) bool
	//BST 的中序遍历是严格递增序列
	inorder = func(node *TreeNode) bool {
		if node == nil {
			return true
		}
		//左节点
		if !inorder(node.Left) {
			return false
		}
		//和上一个访问的节点比较，看看是否合法（单调递增）
		if pre != nil && node.Val <= pre.Val {
			return false
		}
		pre = node
		//右节点
		return inorder(node.Right)
	}
	return inorder(root)
}
'@
Set-Content -Path "problems/098_validate_binary_search_tree/solution.go" -Value $content -Encoding UTF8
Write-Host "Created/Updated problems/098_validate_binary_search_tree/solution.go"

$testContent = @'
package validatebst

import "testing"

func TestIsValidBST(t *testing.T) {
	tests := []struct {
		name string
		root *TreeNode
		want bool
	}{
		{
			name: "Example_1",
			root: &TreeNode{
				Val:   2,
				Left:  &TreeNode{Val: 1},
				Right: &TreeNode{Val: 3},
			},
			want: true,
		},
		{
			name: "Example_2",
			root: &TreeNode{
				Val:   5,
				Left:  &TreeNode{Val: 1},
				Right: &TreeNode{Val: 4, Left: &TreeNode{Val: 3}, Right: &TreeNode{Val: 6}},
			},
			want: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isValidBST(tt.root); got != tt.want {
				t.Errorf("isValidBST() = %v, want %v", got, tt.want)
			}
		})
	}
}
'@
Set-Content -Path "problems/098_validate_binary_search_tree/solution_test.go" -Value $testContent -Encoding UTF8
Write-Host "Created/Updated problems/098_validate_binary_search_tree/solution_test.go"

Write-Host "`nAll files created/updated successfully!"
