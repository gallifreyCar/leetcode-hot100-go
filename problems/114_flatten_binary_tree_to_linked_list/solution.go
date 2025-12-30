package flattenbinarytreetolinkedlist

// 114. 二叉树展开为链表
// 难度：中等
// 标签：树、深度优先搜索
// 链接：https://leetcode.cn/problems/flatten_binary_tree_to_linked_list/

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// Flatten 递归
// 时间复杂度: O(n)
// 空间复杂度: O(h)
func Flatten(root *TreeNode) {
	if root == nil {
		return
	}

	Flatten(root.Left)
	Flatten(root.Right)

	// 保存右子树
	rightSubtree := root.Right
	// 将左子树移动到右子树
	root.Right = root.Left
	root.Left = nil

	// 找到当前右子树的最右节点
	current := root
	for current.Right != nil {
		current = current.Right
	}
	// 将原来的右子树连接到最右节点
	current.Right = rightSubtree
}
