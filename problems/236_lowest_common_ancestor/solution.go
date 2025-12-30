package lowestcommonancestor

// 236. 二叉树的最近公共祖先
// 难度：中等
// 标签：树、深度优先搜索
// 链接：https://leetcode.cn/problems/lowest_common_ancestor/

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// LowestCommonAncestor 递归
// 时间复杂度: O(n)
// 空间复杂度: O(h)
func LowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	if root == nil || root == p || root == q {
		return root
	}

	left := LowestCommonAncestor(root.Left, p, q)
	right := LowestCommonAncestor(root.Right, p, q)

	if left != nil && right != nil {
		return root
	}
	if left != nil {
		return left
	}
	return right
}
