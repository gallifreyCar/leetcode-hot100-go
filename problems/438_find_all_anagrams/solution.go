package findallanagrams

// 438. 找到字符串中所有字母异位词
// 难度：中等
// 标签：滑动窗口、哈希表
// 链接：https://leetcode.cn/problems/find_all_anagrams/

// FindAnagrams 滑动窗口
// 时间复杂度: O(n)
// 空间复杂度: O(字符集大小)
func FindAnagrams(s string, p string) []int {
	if len(s) < len(p) {
		return []int{}
	}

	result := []int{}
	pCount := make(map[byte]int)
	sCount := make(map[byte]int)

	// 统计p中字符频率
	for i := 0; i < len(p); i++ {
		pCount[p[i]]++
	}

	// 初始化窗口
	for i := 0; i < len(p); i++ {
		sCount[s[i]]++
	}

	// 检查第一个窗口
	if isAnagram(pCount, sCount) {
		result = append(result, 0)
	}

	// 滑动窗口
	for i := len(p); i < len(s); i++ {
		// 添加新字符
		sCount[s[i]]++
		// 移除旧字符
		oldChar := s[i-len(p)]
		sCount[oldChar]--
		if sCount[oldChar] == 0 {
			delete(sCount, oldChar)
		}

		if isAnagram(pCount, sCount) {
			result = append(result, i-len(p)+1)
		}
	}

	return result
}

func isAnagram(map1, map2 map[byte]int) bool {
	if len(map1) != len(map2) {
		return false
	}
	for k, v := range map1 {
		if map2[k] != v {
			return false
		}
	}
	return true
}
