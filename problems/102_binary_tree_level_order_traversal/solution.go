package binarytreelevelordertraversal

// 102. 二叉树的层序遍历
// 难度：中等
// 标签：树、广度优先搜索
// 链接：https://leetcode.cn/problems/binary_tree_level_order_traversal/

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// LevelOrder BFS
// 时间复杂度: O(n)
// 空间复杂度: O(n)
func LevelOrder(root *TreeNode) [][]int {
	if root == nil {
		return [][]int{}
	}

	res := [][]int{}
	queue := []*TreeNode{root}

	for len(queue) > 0 {
		levelSize := len(queue)
		level := []int{}

		for i := 0; i < levelSize; i++ {
			node := queue[0]
			queue = queue[1:]
			level = append(level, node.Val)

			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}
		res = append(res, level)
	}

	return res
}
