# 分发更多题解 - 更多二叉树和图论

# 199_binary_tree_right_side_view
$content = @'
package rightsideview

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func rightSideView(root *TreeNode) []int {
	if root == nil {
		return nil
	}
	que := make([]*TreeNode, 0)
	que = append(que, root)
	res := make([]int, 0)
	for len(que) > 0 {
		size := len(que)
		tmp := make([]int, 0)
		for i := 0; i < size; i++ {
			cur := que[0] //队头出列
			que = que[1:]
			tmp = append(tmp, cur.Val)
			if cur.Left != nil {
				que = append(que, cur.Left)
			}
			if cur.Right != nil {
				que = append(que, cur.Right)
			}
		}
		res = append(res, tmp[len(tmp)-1]) //只要队尾
	}
	return res
}
'@
Set-Content -Path "problems/199_binary_tree_right_side_view/solution.go" -Value $content -Encoding UTF8
Write-Host "Created/Updated problems/199_binary_tree_right_side_view/solution.go"

$testContent = @'
package rightsideview

import (
	"reflect"
	"testing"
)

func TestRightSideView(t *testing.T) {
	tests := []struct {
		name string
		root *TreeNode
		want []int
	}{
		{
			name: "Example_1",
			root: &TreeNode{
				Val:   1,
				Left:  &TreeNode{Val: 2, Right: &TreeNode{Val: 5}},
				Right: &TreeNode{Val: 3, Right: &TreeNode{Val: 4}},
			},
			want: []int{1, 3, 4},
		},
		{
			name: "Example_2",
			root: &TreeNode{Val: 1, Right: &TreeNode{Val: 3}},
			want: []int{1, 3},
		},
		{
			name: "Empty",
			root: nil,
			want: nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := rightSideView(tt.root); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("rightSideView() = %v, want %v", got, tt.want)
			}
		})
	}
}
'@
Set-Content -Path "problems/199_binary_tree_right_side_view/solution_test.go" -Value $testContent -Encoding UTF8
Write-Host "Created/Updated problems/199_binary_tree_right_side_view/solution_test.go"

# 114_flatten_binary_tree_to_linked_list
$content = @'
package flattenbinarytree

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func flatten(root *TreeNode) {
	var pre *TreeNode //前驱指针
	var dfs func(node *TreeNode)
	dfs = func(node *TreeNode) {
		if node == nil {
			return
		}
		//要求：先序遍历顺序是 根左右
		//现在要从下往上接，必须反过来 右左根
		dfs(node.Right)
		dfs(node.Left)

		node.Right = pre
		node.Left = nil
		pre = node
	}
	dfs(root)
}
'@
Set-Content -Path "problems/114_flatten_binary_tree_to_linked_list/solution.go" -Value $content -Encoding UTF8
Write-Host "Created/Updated problems/114_flatten_binary_tree_to_linked_list/solution.go"

$testContent = @'
package flattenbinarytree

import "testing"

func TestFlatten(t *testing.T) {
	tests := []struct {
		name string
		root *TreeNode
		want []int
	}{
		{
			name: "Example_1",
			root: &TreeNode{
				Val:   1,
				Left:  &TreeNode{Val: 2, Left: &TreeNode{Val: 3}, Right: &TreeNode{Val: 4}},
				Right: &TreeNode{Val: 5, Right: &TreeNode{Val: 6}},
			},
			want: []int{1, 2, 3, 4, 5, 6},
		},
		{
			name: "Empty",
			root: nil,
			want: []int{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			flatten(tt.root)
			got := treeToList(tt.root)
			if !equal(got, tt.want) {
				t.Errorf("flatten() = %v, want %v", got, tt.want)
			}
		})
	}
}

func treeToList(root *TreeNode) []int {
	var result []int
	for root != nil {
		result = append(result, root.Val)
		root = root.Right
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
'@
Set-Content -Path "problems/114_flatten_binary_tree_to_linked_list/solution_test.go" -Value $testContent -Encoding UTF8
Write-Host "Created/Updated problems/114_flatten_binary_tree_to_linked_list/solution_test.go"

# 105_construct_binary_tree
$content = @'
package constructbinarytree

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func buildTree(preorder []int, inorder []int) *TreeNode {
	if len(preorder) == 0 || len(inorder) == 0 {
		return nil
	}
	// 前序遍历第一个节点是根
	rootVal := preorder[0]
	root := &TreeNode{Val: rootVal}

	// 在中序遍历中找到根的位置（因为无重复值）
	var mid int
	for i, v := range inorder {
		if v == rootVal {
			mid = i
			break
		}
	}

	//中序遍历，根的左边是左节点，根的右边是右节点
	root.Left = buildTree(preorder[1:mid+1], inorder[:mid])
	root.Right = buildTree(preorder[mid+1:], inorder[mid+1:])

	return root
}
'@
Set-Content -Path "problems/105_construct_binary_tree/solution.go" -Value $content -Encoding UTF8
Write-Host "Created/Updated problems/105_construct_binary_tree/solution.go"

$testContent = @'
package constructbinarytree

import (
	"reflect"
	"testing"
)

func TestBuildTree(t *testing.T) {
	tests := []struct {
		name     string
		preorder []int
		inorder  []int
		want     *TreeNode
	}{
		{
			name:     "Example_1",
			preorder: []int{3, 9, 20, 15, 7},
			inorder:  []int{9, 3, 15, 20, 7},
			want: &TreeNode{
				Val:   3,
				Left:  &TreeNode{Val: 9},
				Right: &TreeNode{Val: 20, Left: &TreeNode{Val: 15}, Right: &TreeNode{Val: 7}},
			},
		},
		{
			name:     "Example_2",
			preorder: []int{-1},
			inorder:  []int{-1},
			want:     &TreeNode{Val: -1},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := buildTree(tt.preorder, tt.inorder); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("buildTree() = %v, want %v", got, tt.want)
			}
		})
	}
}
'@
Set-Content -Path "problems/105_construct_binary_tree/solution_test.go" -Value $testContent -Encoding UTF8
Write-Host "Created/Updated problems/105_construct_binary_tree/solution_test.go"

# 124_binary_tree_maximum_path_sum
$content = @'
package maximumpathsum

import "math"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func maxPathSum(root *TreeNode) int {
	//最大路径=当前路径值+左右子树提供的最大"单边路径值"（max(leftGain,RightGain)

	res := math.MinInt32
	var dfs func(node *TreeNode) int
	dfs = func(node *TreeNode) int {
		if node == nil {
			return 0
		}
		//左右子树能提供的最大"单边路径"
		leftGain := max(dfs(node.Left), 0) //负数抛弃
		rightGain := max(dfs(node.Right), 0)
		//当前最大路径值
		curSum := node.Val + leftGain + rightGain
		res = max(res, curSum)
		//返回单边路径
		return node.Val + max(leftGain, rightGain)
	}
	dfs(root)
	return res
}
'@
Set-Content -Path "problems/124_binary_tree_maximum_path_sum/solution.go" -Value $content -Encoding UTF8
Write-Host "Created/Updated problems/124_binary_tree_maximum_path_sum/solution.go"

$testContent = @'
package maximumpathsum

import "testing"

func TestMaxPathSum(t *testing.T) {
	tests := []struct {
		name string
		root *TreeNode
		want int
	}{
		{
			name: "Example_1",
			root: &TreeNode{
				Val:   1,
				Left:  &TreeNode{Val: 2},
				Right: &TreeNode{Val: 3},
			},
			want: 6,
		},
		{
			name: "Example_2",
			root: &TreeNode{
				Val:   -10,
				Left:  &TreeNode{Val: 9},
				Right: &TreeNode{Val: 20, Left: &TreeNode{Val: 15}, Right: &TreeNode{Val: 7}},
			},
			want: 42,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxPathSum(tt.root); got != tt.want {
				t.Errorf("maxPathSum() = %v, want %v", got, tt.want)
			}
		})
	}
}
'@
Set-Content -Path "problems/124_binary_tree_maximum_path_sum/solution_test.go" -Value $testContent -Encoding UTF8
Write-Host "Created/Updated problems/124_binary_tree_maximum_path_sum/solution_test.go"

# 200_number_of_islands
$content = @'
package numberofislands

func numIslands(grid [][]byte) int {
	if len(grid) == 0 || len(grid[0]) == 0 {
		return 0
	}

	m, n := len(grid), len(grid[0])
	cnt := 0
	var dfs func(i, j int)
	dfs = func(i, j int) {
		//越界or不是陆地，直接返回
		if i < 0 || i >= m || j < 0 || j >= n || grid[i][j] != '"'"'1'"'"' {
			return
		}
		//把陆地变成水
		grid[i][j] = '"'"'0'"'"'
		//往四个方向走
		dfs(i-1, j)
		dfs(i, j-1)
		dfs(i, j+1)
		dfs(i+1, j)
	}
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == '"'"'1'"'"' {
				cnt++
				dfs(i, j)
			}

		}
	}
	return cnt
}
'@
Set-Content -Path "problems/200_number_of_islands/solution.go" -Value $content -Encoding UTF8
Write-Host "Created/Updated problems/200_number_of_islands/solution.go"

$testContent = @'
package numberofislands

import "testing"

func TestNumIslands(t *testing.T) {
	tests := []struct {
		name string
		grid [][]byte
		want int
	}{
		{
			name: "Example_1",
			grid: [][]byte{
				{'"'"'1'"'"', '"'"'1'"'"', '"'"'1'"'"', '"'"'1'"'"', '"'"'0'"'"'},
				{'"'"'1'"'"', '"'"'1'"'"', '"'"'0'"'"', '"'"'1'"'"', '"'"'0'"'"'},
				{'"'"'1'"'"', '"'"'1'"'"', '"'"'0'"'"', '"'"'0'"'"', '"'"'0'"'"'},
				{'"'"'0'"'"', '"'"'0'"'"', '"'"'0'"'"', '"'"'0'"'"', '"'"'0'"'"'},
			},
			want: 1,
		},
		{
			name: "Example_2",
			grid: [][]byte{
				{'"'"'1'"'"', '"'"'1'"'"', '"'"'0'"'"', '"'"'0'"'"', '"'"'0'"'"'},
				{'"'"'1'"'"', '"'"'1'"'"', '"'"'0'"'"', '"'"'0'"'"', '"'"'0'"'"'},
				{'"'"'0'"'"', '"'"'0'"'"', '"'"'1'"'"', '"'"'0'"'"', '"'"'0'"'"'},
				{'"'"'0'"'"', '"'"'0'"'"', '"'"'0'"'"', '"'"'1'"'"', '"'"'1'"'"'},
			},
			want: 3,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			grid := make([][]byte, len(tt.grid))
			for i := range tt.grid {
				grid[i] = make([]byte, len(tt.grid[i]))
				copy(grid[i], tt.grid[i])
			}
			if got := numIslands(grid); got != tt.want {
				t.Errorf("numIslands() = %v, want %v", got, tt.want)
			}
		})
	}
}
'@
Set-Content -Path "problems/200_number_of_islands/solution_test.go" -Value $testContent -Encoding UTF8
Write-Host "Created/Updated problems/200_number_of_islands/solution_test.go"

Write-Host "`nAll files created/updated successfully!"
