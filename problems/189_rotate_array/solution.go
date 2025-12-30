package rotatearray

import "slices"

// 189. 轮转数组
// 难度：中等
// 标签：数组
// 链接：https://leetcode.cn/problems/rotate_array/

// Rotate 使用额外数组
// 时间复杂度: O(n)
// 空间复杂度: O(n)
func Rotate(nums []int, k int) {
	newNums := make([]int, len(nums))
	for i := 0; i < len(nums); i++ {
		t := (i + k) % len(nums)
		newNums[t] = nums[i]
	}
	copy(nums, newNums)
}

// RotateV2 三次反转，原地算法
// 时间复杂度: O(n)
// 空间复杂度: O(1)
func RotateV2(nums []int, k int) {
	k %= len(nums)
	slices.Reverse(nums)
	slices.Reverse(nums[:k])
	slices.Reverse(nums[k:])
}
