package validparentheses

// 20. 有效的括号
// 难度：简单
// 标签：栈
// 链接：https://leetcode.cn/problems/valid_parentheses/

// IsValid 栈
// 时间复杂度: O(n)
// 空间复杂度: O(n)
func IsValid(s string) bool {
	if len(s)%2 != 0 {
		return false
	}

	stack := []rune{}
	pairs := map[rune]rune{
		')': '(',
		'}': '{',
		']': '[',
	}

	for _, ch := range s {
		if ch == '(' || ch == '{' || ch == '[' {
			stack = append(stack, ch)
		} else {
			if len(stack) == 0 || stack[len(stack)-1] != pairs[ch] {
				return false
			}
			stack = stack[:len(stack)-1]
		}
	}

	return len(stack) == 0
}
