package rightsideview

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func rightSideView(root *TreeNode) []int {
	if root == nil {
		return nil
	}
	que := make([]*TreeNode, 0)
	que = append(que, root)
	res := make([]int, 0)
	for len(que) > 0 {
		size := len(que)
		tmp := make([]int, 0)
		for i := 0; i < size; i++ {
			cur := que[0] //闃熷ご鍑哄垪
			que = que[1:]
			tmp = append(tmp, cur.Val)
			if cur.Left != nil {
				que = append(que, cur.Left)
			}
			if cur.Right != nil {
				que = append(que, cur.Right)
			}
		}
		res = append(res, tmp[len(tmp)-1]) //鍙闃熷熬
	}
	return res
}
