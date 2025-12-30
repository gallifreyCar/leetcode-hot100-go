package rottingoranges

// 994. 腐烂的橘子
// 难度：中等
// 标签：广度优先搜索
// 链接：https://leetcode.cn/problems/rotting_oranges/

// OrangesRotting BFS
// 时间复杂度: O(m*n)
// 空间复杂度: O(m*n)
func OrangesRotting(grid [][]int) int {
	m, n := len(grid), len(grid[0])
	queue := [][]int{}
	freshCount := 0

	// 统计腐烂桔子和新鲜桔子数量
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == 2 {
				queue = append(queue, []int{i, j})
			} else if grid[i][j] == 1 {
				freshCount++
			}
		}
	}

	if freshCount == 0 {
		return 0
	}

	minutes := 0
	directions := [][]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}

	for len(queue) > 0 {
		size := len(queue)
		for i := 0; i < size; i++ {
			pos := queue[0]
			queue = queue[1:]
			row, col := pos[0], pos[1]

			for _, dir := range directions {
				newRow, newCol := row+dir[0], col+dir[1]
				if newRow >= 0 && newRow < m && newCol >= 0 && newCol < n && grid[newRow][newCol] == 1 {
					grid[newRow][newCol] = 2
					freshCount--
					queue = append(queue, []int{newRow, newCol})
				}
			}
		}
		minutes++
	}

	if freshCount > 0 {
		return -1
	}
	return minutes - 1
}
