package convertsortedarraytobst

// 108. 将有序数组转换为二叉搜索树
// 难度：简单
// 标签：树、二叉搜索树
// 链接：https://leetcode.cn/problems/convert_sorted_array_to_bst/

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// SortedArrayToBST 递归分治
// 时间复杂度: O(n)
// 空间复杂度: O(log n)
func SortedArrayToBST(nums []int) *TreeNode {
	if len(nums) == 0 {
		return nil
	}
	mid := len(nums) / 2
	return &TreeNode{
		Val:   nums[mid],
		Left:  SortedArrayToBST(nums[:mid]),
		Right: SortedArrayToBST(nums[mid+1:]),
	}
}
