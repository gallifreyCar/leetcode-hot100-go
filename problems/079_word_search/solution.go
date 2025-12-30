package wordsearch

// 79. 单词搜索
// 难度：中等
// 标签：回溯、矩阵
// 链接：https://leetcode.cn/problems/word_search/

// Exist 回溯+DFS
// 时间复杂度: O(m*n*3^L)
// 空间复杂度: O(L)
func Exist(board [][]byte, word string) bool {
	if len(board) == 0 || len(word) == 0 {
		return false
	}

	m, n := len(board), len(board[0])
	visited := make([][]bool, m)
	for i := range visited {
		visited[i] = make([]bool, n)
	}

	var backtrack func(int, int, int) bool
	backtrack = func(i, j, k int) bool {
		if k == len(word) {
			return true
		}
		if i < 0 || i >= m || j < 0 || j >= n || visited[i][j] || board[i][j] != word[k] {
			return false
		}

		visited[i][j] = true
		result := backtrack(i+1, j, k+1) || backtrack(i-1, j, k+1) ||
			backtrack(i, j+1, k+1) || backtrack(i, j-1, k+1)
		visited[i][j] = false
		return result
	}

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if backtrack(i, j, 0) {
				return true
			}
		}
	}
	return false
}
