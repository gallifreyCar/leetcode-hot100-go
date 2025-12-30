package max_consecutive_ones_iii

// LongestOnes 滑动窗口
// 时间复杂度: O(n)
// 空间复杂度: O(1)
func LongestOnes(nums []int, k int) int {
	start, res := 0, 0
	zeroCount := 0

	for end := 0; end < len(nums); end++ {
		// 1. 加入新元素
		if nums[end] == 0 {
			zeroCount++
		}

		// 2. 收缩左边界（窗口不合法时）
		for zeroCount > k {
			if nums[start] == 0 {
				zeroCount--
			}
			start++
		}

		// 3. 更新结果
		res = max(res, end-start+1)
	}

	return res
}
