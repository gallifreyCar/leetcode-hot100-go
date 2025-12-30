package firstmissingpositive

// 41. 缺失的第一个正数
// 难度：困难
// 标签：数组、哈希表
// 链接：https://leetcode.cn/problems/first_missing_positive/

// FirstMissingPositive 原地哈希
// 时间复杂度: O(n)
// 空间复杂度: O(1)
func FirstMissingPositive(nums []int) int {
	n := len(nums)
	// 第一次遍历：将所有负数、0和大于n的数改为n+1
	for i := 0; i < n; i++ {
		if nums[i] <= 0 || nums[i] > n {
			nums[i] = n + 1
		}
	}

	// 第二次遍历：使用数组索引标记出现的正数
	for i := 0; i < n; i++ {
		abs := nums[i]
		if abs < 0 {
			abs = -abs
		}
		if abs <= n {
			if nums[abs-1] > 0 {
				nums[abs-1] = -nums[abs-1]
			}
		}
	}

	// 第三次遍历：找到第一个正数的位置
	for i := 0; i < n; i++ {
		if nums[i] > 0 {
			return i + 1
		}
	}
	return n + 1
}
