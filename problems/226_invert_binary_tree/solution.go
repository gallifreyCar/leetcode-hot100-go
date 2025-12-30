package invertbinarytree

// 226. 翻转二叉树
// 难度：简单
// 标签：树、深度优先搜索
// 链接：https://leetcode.cn/problems/invert_binary_tree/

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// InvertTree 递归
// 时间复杂度: O(n)
// 空间复杂度: O(h)
func InvertTree(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	root.Left, root.Right = InvertTree(root.Right), InvertTree(root.Left)
	return root
}
