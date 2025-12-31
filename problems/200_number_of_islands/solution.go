package numberofislands

func numIslands(grid [][]byte) int {
	if len(grid) == 0 || len(grid[0]) == 0 {
		return 0
	}

	m, n := len(grid), len(grid[0])
	cnt := 0
	var dfs func(i, j int)
	dfs = func(i, j int) {
		//越界or不是陆地，直接返回
		if i < 0 || i >= m || j < 0 || j >= n || grid[i][j] != '1' {
			return
		}
		//把陆地变成水
		grid[i][j] = '0'
		//往四个方向走
		dfs(i-1, j)
		dfs(i, j-1)
		dfs(i, j+1)
		dfs(i+1, j)
	}
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == '1' {
				cnt++
				dfs(i, j)
			}

		}
	}
	return cnt
}
