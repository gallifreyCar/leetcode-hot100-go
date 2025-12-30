package minimumwindowsubstring

// 76. 最小覆盖子串
// 难度：困难
// 标签：滑动窗口、哈希表
// 链接：https://leetcode.cn/problems/minimum_window_substring/

// MinWindow 滑动窗口+哈希表
// 时间复杂度: O(n)
// 空间复杂度: O(字符集大小)
func MinWindow(s string, t string) string {
	if len(s) < len(t) {
		return ""
	}

	tCount := make(map[byte]int)
	sCount := make(map[byte]int)

	// 统计t中字符频率
	for i := 0; i < len(t); i++ {
		tCount[t[i]]++
	}

	l, r := 0, 0
	minLen := len(s) + 1
	minStart := 0
	valid := 0 // 符合要求的字符个数

	for r < len(s) {
		// 扩大窗口
		c := s[r]
		r++
		if _, exists := tCount[c]; exists {
			sCount[c]++
			if sCount[c] == tCount[c] {
				valid++
			}
		}

		// 缩小窗口
		for valid == len(tCount) {
			// 更新最小覆盖子串
			if r-l < minLen {
				minStart = l
				minLen = r - l
			}

			// 移除左边字符
			d := s[l]
			l++
			if _, exists := tCount[d]; exists {
				if sCount[d] == tCount[d] {
					valid--
				}
				sCount[d]--
			}
		}
	}

	if minLen == len(s)+1 {
		return ""
	}
	return s[minStart : minStart+minLen]
}

// MinWindowV2 优化：使用need记录还需要的字符数
// 时间复杂度: O(n)
// 空间复杂度: O(字符集大小)
func MinWindowV2(s string, t string) string {
	if len(s) < len(t) {
		return ""
	}

	need := make(map[byte]int)
	window := make(map[byte]int)

	for i := 0; i < len(t); i++ {
		need[t[i]]++
	}

	l, r := 0, 0
	minLen := len(s) + 1
	start := 0
	valid := 0

	for r < len(s) {
		c := s[r]
		r++
		if _, exists := need[c]; exists {
			window[c]++
			if window[c] == need[c] {
				valid++
			}
		}

		for valid == len(need) {
			if r-l < minLen {
				start = l
				minLen = r - l
			}

			d := s[l]
			l++
			if _, exists := need[d]; exists {
				if window[d] == need[d] {
					valid--
				}
				window[d]--
			}
		}
	}

	if minLen == len(s)+1 {
		return ""
	}
	return s[start : start+minLen]
}

// MinWindowV3 再优化：只在必要时检查
// 时间复杂度: O(n)
// 空间复杂度: O(字符集大小)
func MinWindowV3(s string, t string) string {
	if len(s) < len(t) {
		return ""
	}

	need := make(map[byte]int)
	for i := 0; i < len(t); i++ {
		need[t[i]]++
	}

	l, r := 0, 0
	minLen := len(s) + 1
	start := 0
	required := len(t)

	for r < len(s) {
		if need[s[r]] > 0 {
			required--
		}
		need[s[r]]--
		r++

		for required == 0 {
			if r-l < minLen {
				start = l
				minLen = r - l
			}

			need[s[l]]++
			if need[s[l]] > 0 {
				required++
			}
			l++
		}
	}

	if minLen == len(s)+1 {
		return ""
	}
	return s[start : start+minLen]
}
