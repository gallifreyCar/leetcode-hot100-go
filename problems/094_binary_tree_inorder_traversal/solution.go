package binarytreeinordertraversal

// 94. 二叉树的中序遍历
// 难度：简单
// 标签：树、栈
// 链接：https://leetcode.cn/problems/binary_tree_inorder_traversal/

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// InorderTraversal 递归
// 时间复杂度: O(n)
// 空间复杂度: O(n)
func InorderTraversal(root *TreeNode) []int {
	res := []int{}
	var inorder func(*TreeNode)
	inorder = func(node *TreeNode) {
		if node == nil {
			return
		}
		inorder(node.Left)
		res = append(res, node.Val)
		inorder(node.Right)
	}
	inorder(root)
	return res
}
