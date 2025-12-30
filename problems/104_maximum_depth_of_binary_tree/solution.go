package maximumdepthofbinarytree

// 104. 二叉树的最大深度
// 难度：简单
// 标签：树、深度优先搜索
// 链接：https://leetcode.cn/problems/maximum_depth_of_binary_tree/

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// MaxDepth 递归
// 时间复杂度: O(n)
// 空间复杂度: O(h)
func MaxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	return max(MaxDepth(root.Left), MaxDepth(root.Right)) + 1
}
