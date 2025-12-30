package subarraysumequalsk

// 560. 和为K的子数组
// 难度：中等
// 标签：前缀和、哈希表
// 链接：https://leetcode.cn/problems/subarray_sum_equals_k/

// SubarraySum 前缀和+哈希表
// 时间复杂度: O(n)
// 空间复杂度: O(n)
func SubarraySum(nums []int, k int) int {
	count := 0
	sum := 0
	prefixSum := make(map[int]int)
	prefixSum[0] = 1 // 前缀和为0的有一种情况

	for i := 0; i < len(nums); i++ {
		sum += nums[i]
		// 如果存在 sum - k，说明找到了和为k的子数组
		if val, exists := prefixSum[sum-k]; exists {
			count += val
		}
		prefixSum[sum]++
	}

	return count
}
