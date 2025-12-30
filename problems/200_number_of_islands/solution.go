package numberofislands

// 200. 岛屿数量
// 难度：中等
// 标签：深度优先搜索、广度优先搜索
// 链接：https://leetcode.cn/problems/number_of_islands/

// NumIslands DFS
// 时间复杂度: O(m*n)
// 空间复杂度: O(m*n)
func NumIslands(grid [][]byte) int {
	if len(grid) == 0 {
		return 0
	}

	count := 0
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			if grid[i][j] == '1' {
				count++
				dfs(grid, i, j)
			}
		}
	}
	return count
}

func dfs(grid [][]byte, i, j int) {
	if i < 0 || i >= len(grid) || j < 0 || j >= len(grid[0]) || grid[i][j] == '0' {
		return
	}
	grid[i][j] = '0' // 标记为已访问
	dfs(grid, i-1, j)
	dfs(grid, i+1, j)
	dfs(grid, i, j-1)
	dfs(grid, i, j+1)
}
