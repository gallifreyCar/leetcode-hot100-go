package maximumpathsum

import "math"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func maxPathSum(root *TreeNode) int {
	//鏈€澶ц矾寰?褰撳墠璺緞鍊?宸﹀彸瀛愭爲鎻愪緵鐨勬渶澶?鍗曡竟璺緞鍊?锛坢ax(leftGain,RightGain)

	res := math.MinInt32
	var dfs func(node *TreeNode) int
	dfs = func(node *TreeNode) int {
		if node == nil {
			return 0
		}
		//宸﹀彸瀛愭爲鑳芥彁渚涚殑鏈€澶?鍗曡竟璺緞"
		leftGain := max(dfs(node.Left), 0) //璐熸暟鎶涘純
		rightGain := max(dfs(node.Right), 0)
		//褰撳墠鏈€澶ц矾寰勫€?
		curSum := node.Val + leftGain + rightGain
		res = max(res, curSum)
		//杩斿洖鍗曡竟璺緞
		return node.Val + max(leftGain, rightGain)
	}
	dfs(root)
	return res
}
