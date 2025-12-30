package symmetrictree

// 101. 对称二叉树
// 难度：简单
// 标签：树、深度优先搜索
// 链接：https://leetcode.cn/problems/symmetric_tree/

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// IsSymmetric 递归
// 时间复杂度: O(n)
// 空间复杂度: O(h)
func IsSymmetric(root *TreeNode) bool {
	if root == nil {
		return true
	}
	return isMirror(root.Left, root.Right)
}

func isMirror(left, right *TreeNode) bool {
	if left == nil && right == nil {
		return true
	}
	if left == nil || right == nil {
		return false
	}
	return left.Val == right.Val && isMirror(left.Left, right.Right) && isMirror(left.Right, right.Left)
}
