package flattenbinarytree

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func flatten(root *TreeNode) {
	var pre *TreeNode //鍓嶉┍鎸囬拡
	var dfs func(node *TreeNode)
	dfs = func(node *TreeNode) {
		if node == nil {
			return
		}
		//瑕佹眰锛氬厛搴忛亶鍘嗛『搴忔槸 鏍瑰乏鍙?
		//鐜板湪瑕佷粠涓嬪線涓婃帴锛屽繀椤诲弽杩囨潵 鍙冲乏鏍?
		dfs(node.Right)
		dfs(node.Left)

		node.Right = pre
		node.Left = nil
		pre = node
	}
	dfs(root)
}
