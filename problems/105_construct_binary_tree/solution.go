package constructbinarytree

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func buildTree(preorder []int, inorder []int) *TreeNode {
	if len(preorder) == 0 || len(inorder) == 0 {
		return nil
	}
	// 鍓嶅簭閬嶅巻绗竴涓妭鐐规槸鏍?
	rootVal := preorder[0]
	root := &TreeNode{Val: rootVal}

	// 鍦ㄤ腑搴忛亶鍘嗕腑鎵惧埌鏍圭殑浣嶇疆锛堝洜涓烘棤閲嶅鍊硷級
	var mid int
	for i, v := range inorder {
		if v == rootVal {
			mid = i
			break
		}
	}

	//涓簭閬嶅巻锛屾牴鐨勫乏杈规槸宸﹁妭鐐癸紝鏍圭殑鍙宠竟鏄彸鑺傜偣
	root.Left = buildTree(preorder[1:mid+1], inorder[:mid])
	root.Right = buildTree(preorder[mid+1:], inorder[mid+1:])

	return root
}
