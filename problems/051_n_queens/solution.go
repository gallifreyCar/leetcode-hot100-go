package nqueens

import "strings"

// 51. N皇后
// 难度：困难
// 标签：回溯
// 链接：https://leetcode.cn/problems/n_queens/
func solveNQueens(n int) [][]string {
	diagA := make([]bool, 2*n-1) //主对角线 i-j=常数 为了不越界 i-j+n-1
	diagB := make([]bool, 2*n-1) //副对角线 i+j=常数
	// 比如n=4，q在(0,1)，那么diagA(-2)=true diagB(1)=true
	// 因此 diagA可以标记(0,1),(1,2),(2,3) diagB可以过滤(0,1),(1,0)

	col := make([]bool, n) //标记这一列有没有皇后q

	//遍历所有的行
	var res [][]string
	var path []string
	var dfs func(row int)
	dfs = func(row int) {
		if row == n {
			tmp := append([]string(nil), path...)
			res = append(res, tmp)
			return
		}
		for i := 0; i < n; i++ {
			if diagA[row-i+n-1] || diagB[row+i] || col[i] {
				continue
			}
			diagA[row-i+n-1], diagB[row+i], col[i] = true, true, true
			rowStr := strings.Repeat(".", i) + "Q" + strings.Repeat(".", n-i-1)
			path = append(path, rowStr)
			dfs(row + 1)
			diagA[row-i+n-1], diagB[row+i], col[i] = false, false, false
			path = path[:len(path)-1]
		}
	}
	dfs(0)
	return res
}
