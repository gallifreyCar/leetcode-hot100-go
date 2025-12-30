package longestsubstringwithoutrepeating

// 3. 无重复字符的最长子串
// 难度：中等
// 标签：滑动窗口、哈希表
// 链接：https://leetcode.cn/problems/longest_substring_without_repeating/

// LengthOfLongestSubstring 滑动窗口+哈希表
// 时间复杂度: O(n)
// 空间复杂度: O(字符集大小)
func LengthOfLongestSubstring(s string) int {
	if len(s) == 0 {
		return 0
	}
	l, r := 0, 0
	freq := make(map[byte]int)
	maxLen := 0
	for r < len(s) {
		freq[s[r]]++
		// 如果出现重复，移动左指针
		for freq[s[r]] > 1 {
			freq[s[l]]--
			l++
		}
		maxLen = max(maxLen, r-l+1)
		r++
	}
	return maxLen
}

// LengthOfLongestSubstringV2 优化：使用map存储索引
// 时间复杂度: O(n)
// 空间复杂度: O(字符集大小)
func LengthOfLongestSubstringV2(s string) int {
	if len(s) == 0 {
		return 0
	}
	l, r := 0, 0
	indexMap := make(map[byte]int)
	maxLen := 0
	for r < len(s) {
		// 如果当前字符出现过，移动左指针到重复字符的下一个位置
		if idx, exists := indexMap[s[r]]; exists && idx >= l {
			l = idx + 1
		}
		indexMap[s[r]] = r
		maxLen = max(maxLen, r-l+1)
		r++
	}
	return maxLen
}
