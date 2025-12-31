package movezeroes

// 283. 移动零
// 难度：简单
// 标签：双指针、数组
// 链接：https://leetcode.cn/problems/move_zeroes/
// MoveZeroes 双指针解法
// 时间复杂度: O(n)
// 空间复杂度: O(1)
func MoveZeroes(nums []int) {
	// 双指针
	r := 0
	for l := 0; l < len(nums); l++ {
		if nums[l] != 0 {
			nums[l], nums[r] = nums[r], nums[l]
			r++
		}
	}
}
