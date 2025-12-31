package validatebst

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isValidBST(root *TreeNode) bool {
	var pre *TreeNode
	var inorder func(node *TreeNode) bool
	//BST 鐨勪腑搴忛亶鍘嗘槸涓ユ牸閫掑搴忓垪
	inorder = func(node *TreeNode) bool {
		if node == nil {
			return true
		}
		//宸﹁妭鐐?
		if !inorder(node.Left) {
			return false
		}
		//鍜屼笂涓€涓闂殑鑺傜偣姣旇緝锛岀湅鐪嬫槸鍚﹀悎娉曪紙鍗曡皟閫掑锛?
		if pre != nil && node.Val <= pre.Val {
			return false
		}
		pre = node
		//鍙宠妭鐐?
		return inorder(node.Right)
	}
	return inorder(root)
}
