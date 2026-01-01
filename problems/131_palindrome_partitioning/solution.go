package palindromepartitioning

// 131. 分割回文串
// 难度：中等
// 标签：回溯
// 链接：https://leetcode.cn/problems/palindrome_partitioning/
func partition(s string) [][]string {
	dp := make([][]bool, len(s)) //动态规划判断是否回文串
	for i := range dp {
		dp[i] = make([]bool, len(s))
	}
	n := len(s)

	for i := n - 1; i >= 0; i-- {
		for j := i; j < n; j++ {
			// 首尾字母相等且内部也是回文则为回文
			// 特殊情况：首位字母相等，但是长度是 1，2，3 也是回文串 例子：a,aa,aba
			if s[i] == s[j] && (j-i <= 2 || dp[i+1][j-1]) {
				dp[i][j] = true
			}
		}
	}

	res := make([][]string, 0)
	path := make([]string, 0)
	var dfs func(start int)
	dfs = func(start int) {
		if start == n {
			var tmp []string
			tmp = append(tmp, path...) //path是切片，本质是指针，必须copy
			res = append(res, tmp)
			return
		}
		for i := start; i < n; i++ {
			if dp[start][i] {
				path = append(path, s[start:i+1])
				dfs(i + 1)
				path = path[:len(path)-1]
			}
		}
	}
	dfs(0)

	return res
}
