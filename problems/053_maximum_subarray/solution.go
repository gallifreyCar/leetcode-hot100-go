package maximumsubarray

import "math"

// 53. 最大子数组和
// 难度：中等
// 标签：动态规划、分治、数组
// 链接：https://leetcode.cn/problems/maximum_subarray/

// MaxSubArray 动态规划
// 时间复杂度: O(n)
// 空间复杂度: O(n)
func MaxSubArray(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	dp := make([]int, len(nums))
	dp[0] = nums[0]
	maxSum := nums[0]

	for i := 1; i < len(nums); i++ {
		dp[i] = max(dp[i-1]+nums[i], nums[i])
		maxSum = max(maxSum, dp[i])
	}

	return maxSum
}

// MaxSubArrayV2 空间优化
// 时间复杂度: O(n)
// 空间复杂度: O(1)
func MaxSubArrayV2(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	prev := nums[0]
	maxSum := nums[0]

	for i := 1; i < len(nums); i++ {
		prev = max(prev+nums[i], nums[i])
		maxSum = max(maxSum, prev)
	}

	return maxSum
}

// MaxSubArrayV3 分治算法
// 时间复杂度: O(n log n)
// 空间复杂度: O(log n)
func MaxSubArrayV3(nums []int) int {
	return maxSubArrayHelper(nums, 0, len(nums)-1)
}

func maxSubArrayHelper(nums []int, left, right int) int {
	if left == right {
		return nums[left]
	}

	mid := left + (right-left)/2
	leftMax := maxSubArrayHelper(nums, left, mid)
	rightMax := maxSubArrayHelper(nums, mid+1, right)
	crossMax := maxCrossingSum(nums, left, mid, right)

	return max(max(leftMax, rightMax), crossMax)
}

func maxCrossingSum(nums []int, left, mid, right int) int {
	leftSum := math.MinInt32
	sum := 0
	for i := mid; i >= left; i-- {
		sum += nums[i]
		if sum > leftSum {
			leftSum = sum
		}
	}

	rightSum := math.MinInt32
	sum = 0
	for i := mid + 1; i <= right; i++ {
		sum += nums[i]
		if sum > rightSum {
			rightSum = sum
		}
	}

	return leftSum + rightSum
}
