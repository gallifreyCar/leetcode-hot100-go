package lettercombinations

// 17. 电话号码的字母组合
// 难度：中等
// 标签：回溯
// 链接：https://leetcode.cn/problems/letter_combinations/

// LetterCombinations 回溯
// 时间复杂度: O(3^m * 4^n)
// 空间复杂度: O(m+n)
func LetterCombinations(digits string) []string {
	if len(digits) == 0 {
		return []string{}
	}

	phoneMap := map[byte]string{
		'2': "abc",
		'3': "def",
		'4': "ghi",
		'5': "jkl",
		'6': "mno",
		'7': "pqrs",
		'8': "tuv",
		'9': "wxyz",
	}

	res := []string{}
	path := []byte{}

	var backtrack func(int)
	backtrack = func(index int) {
		if index == len(digits) {
			res = append(res, string(path))
			return
		}

		letters := phoneMap[digits[index]]
		for i := 0; i < len(letters); i++ {
			path = append(path, letters[i])
			backtrack(index + 1)
			path = path[:len(path)-1]
		}
	}

	backtrack(0)
	return res
}
