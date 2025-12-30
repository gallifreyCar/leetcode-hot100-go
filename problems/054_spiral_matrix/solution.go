package spiralmatrix

// 54. 螺旋矩阵
// 难度：中等
// 标签：数组、矩阵
// 链接：https://leetcode.cn/problems/spiral_matrix/

// SpiralOrder 模拟法
// 时间复杂度: O(m*n)
// 空间复杂度: O(1)
func SpiralOrder(matrix [][]int) []int {
	m := len(matrix)
	n := len(matrix[0])
	res := make([]int, 0, m*n)
	directions := [][]int{{0, 1}, {1, 0}, {0, -1}, {-1, 0}} // 右下左上
	direction := 0
	x, y := 0, 0

	for i := 0; i < m*n; i++ {
		res = append(res, matrix[x][y])
		matrix[x][y] = 101 // 标记已访问
		nextX := x + directions[direction][0]
		nextY := y + directions[direction][1]
		// 判断是否需要转向
		if nextX < 0 || nextX >= m || nextY < 0 || nextY >= n || matrix[nextX][nextY] == 101 {
			direction = (direction + 1) % 4
			nextX = x + directions[direction][0]
			nextY = y + directions[direction][1]
		}
		x = nextX
		y = nextY
	}
	return res
}

// SpiralOrderV2 按层模拟
// 时间复杂度: O(m*n)
// 空间复杂度: O(1)
func SpiralOrderV2(matrix [][]int) []int {
	if len(matrix) == 0 {
		return []int{}
	}

	res := make([]int, 0, len(matrix)*len(matrix[0]))
	top, bottom := 0, len(matrix)-1
	left, right := 0, len(matrix[0])-1

	for top <= bottom && left <= right {
		// 从左到右
		for i := left; i <= right; i++ {
			res = append(res, matrix[top][i])
		}
		top++

		// 从上到下
		for i := top; i <= bottom; i++ {
			res = append(res, matrix[i][right])
		}
		right--

		// 从右到左
		if top <= bottom {
			for i := right; i >= left; i-- {
				res = append(res, matrix[bottom][i])
			}
			bottom--
		}

		// 从下到上
		if left <= right {
			for i := bottom; i >= top; i-- {
				res = append(res, matrix[i][left])
			}
			left++
		}
	}
	return res
}
