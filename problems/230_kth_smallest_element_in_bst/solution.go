package kthsmallestelementinbst

// 230. 二叉搜索树中第K小的元素
// 难度：中等
// 标签：树、二叉搜索树
// 链接：https://leetcode.cn/problems/kth_smallest_element_in_bst/

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// KthSmallest 中序遍历
// 时间复杂度: O(n)
// 空间复杂度: O(h)
func KthSmallest(root *TreeNode, k int) int {
	count := 0
	result := 0
	var inorder func(*TreeNode)
	inorder = func(node *TreeNode) {
		if node == nil || count >= k {
			return
		}
		inorder(node.Left)
		count++
		if count == k {
			result = node.Val
			return
		}
		inorder(node.Right)
	}
	inorder(root)
	return result
}
