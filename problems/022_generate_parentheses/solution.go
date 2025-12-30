package generateparentheses

// 22. 括号生成
// 难度：中等
// 标签：回溯
// 链接：https://leetcode.cn/problems/generate_parentheses/

// GenerateParenthesis 回溯
// 时间复杂度: O(4^n / sqrt(n))
// 空间复杂度: O(n)
func GenerateParenthesis(n int) []string {
	res := []string{}
	path := []byte{}

	var backtrack func(int, int)
	backtrack = func(left, right int) {
		// 左括号数量不能超过n，右括号数量不能超过左括号
		if left > n || right > left {
			return
		}
		if left == n && right == n {
			res = append(res, string(path))
			return
		}

		// 添加左括号
		path = append(path, '(')
		backtrack(left+1, right)
		path = path[:len(path)-1]

		// 添加右括号
		path = append(path, ')')
		backtrack(left, right+1)
		path = path[:len(path)-1]
	}

	backtrack(0, 0)
	return res
}
