package setmatrixzeroes

// 73. 矩阵置零
// 难度：中等
// 标签：数组、矩阵
// 链接：https://leetcode.cn/problems/set_matrix_zeroes/

// SetZeroes 使用两个数组记录
// 时间复杂度: O(m*n)
// 空间复杂度: O(m+n)
func SetZeroes(matrix [][]int) {
	rows := make([]bool, len(matrix))
	cols := make([]bool, len(matrix[0]))
	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[0]); j++ {
			if matrix[i][j] == 0 {
				rows[i] = true
				cols[j] = true
			}
		}
	}
	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[0]); j++ {
			if rows[i] || cols[j] {
				matrix[i][j] = 0
			}
		}
	}
}

// SetZeroesV2 使用第一行和第一列作为标记
// 时间复杂度: O(m*n)
// 空间复杂度: O(1)
func SetZeroesV2(matrix [][]int) {
	// 判断第一行和第一列是否有0
	firstRowHasZero := false
	firstColHasZero := false
	for j := 0; j < len(matrix[0]); j++ {
		if matrix[0][j] == 0 {
			firstRowHasZero = true
			break
		}
	}
	for i := 0; i < len(matrix); i++ {
		if matrix[i][0] == 0 {
			firstColHasZero = true
			break
		}
	}

	// 使用第一行和第一列记录其他位置的0
	for i := 1; i < len(matrix); i++ {
		for j := 1; j < len(matrix[0]); j++ {
			if matrix[i][j] == 0 {
				matrix[i][0] = 0
				matrix[0][j] = 0
			}
		}
	}

	// 根据第一行和第一列的标记，置零其他位置
	for i := 1; i < len(matrix); i++ {
		for j := 1; j < len(matrix[0]); j++ {
			if matrix[i][0] == 0 || matrix[0][j] == 0 {
				matrix[i][j] = 0
			}
		}
	}

	// 最后处理第一行和第一列
	if firstRowHasZero {
		for j := 0; j < len(matrix[0]); j++ {
			matrix[0][j] = 0
		}
	}
	if firstColHasZero {
		for i := 0; i < len(matrix); i++ {
			matrix[i][0] = 0
		}
	}
}
