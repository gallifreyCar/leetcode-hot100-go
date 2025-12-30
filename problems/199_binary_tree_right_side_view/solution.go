package binarytreerightsideview

// 199. 二叉树的右视图
// 难度：中等
// 标签：树、广度优先搜索
// 链接：https://leetcode.cn/problems/binary_tree_right_side_view/

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// RightSideView BFS
// 时间复杂度: O(n)
// 空间复杂度: O(n)
func RightSideView(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}

	res := []int{}
	queue := []*TreeNode{root}

	for len(queue) > 0 {
		levelSize := len(queue)
		for i := 0; i < levelSize; i++ {
			node := queue[0]
			queue = queue[1:]

			// 每层最后一个节点就是右视图能看到的
			if i == levelSize-1 {
				res = append(res, node.Val)
			}

			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}
	}

	return res
}
