package rotateimage

import "slices"

// 48. 旋转图像
// 难度：中等
// 标签：数组、矩阵
// 链接：https://leetcode.cn/problems/rotate_image/

// Rotate 先转置再翻转每一行
// 时间复杂度: O(n²)
// 空间复杂度: O(1)
func Rotate(matrix [][]int) {
	n := len(matrix)
	// 转置矩阵
	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			matrix[i][j], matrix[j][i] = matrix[j][i], matrix[i][j]
		}
	}
	// 翻转每一行
	for i := 0; i < n; i++ {
		slices.Reverse(matrix[i])
	}
}
