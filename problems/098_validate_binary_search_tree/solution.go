package validatebinarysearchtree

import "math"

// 98. 验证二叉搜索树
// 难度：中等
// 标签：树、深度优先搜索、二叉搜索树
// 链接：https://leetcode.cn/problems/validate_binary_search_tree/

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// IsValidBST 递归验证
// 时间复杂度: O(n)
// 空间复杂度: O(h)
func IsValidBST(root *TreeNode) bool {
	return isValidBSTHelper(root, math.MinInt64, math.MaxInt64)
}

func isValidBSTHelper(node *TreeNode, min, max int) bool {
	if node == nil {
		return true
	}
	if node.Val <= min || node.Val >= max {
		return false
	}
	return isValidBSTHelper(node.Left, min, node.Val) && isValidBSTHelper(node.Right, node.Val, max)
}
