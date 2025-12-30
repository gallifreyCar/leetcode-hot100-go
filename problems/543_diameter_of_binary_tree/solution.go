package diameterofbinarytree

// 543. 二叉树的直径
// 难度：简单
// 标签：树、深度优先搜索
// 链接：https://leetcode.cn/problems/diameter_of_binary_tree/

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// DiameterOfBinaryTree 递归
// 时间复杂度: O(n)
// 空间复杂度: O(h)
func DiameterOfBinaryTree(root *TreeNode) int {
	maxDiameter := 0
	var depth func(*TreeNode) int
	depth = func(node *TreeNode) int {
		if node == nil {
			return 0
		}
		left := depth(node.Left)
		right := depth(node.Right)
		maxDiameter = max(maxDiameter, left+right)
		return max(left, right) + 1
	}
	depth(root)
	return maxDiameter
}
