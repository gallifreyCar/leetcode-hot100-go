package threesum

import "sort"

// 15. 三数之和
// 难度：中等
// 标签：双指针、数组、排序
// 链接：https://leetcode.cn/problems/three_sum/

// ThreeSum 双指针解法
// 时间复杂度: O(n²)
// 空间复杂度: O(1)
func ThreeSum(nums []int) [][]int {
	res := [][]int{}
	// 先排序
	sort.Ints(nums)

	// i 是第一个数字
	for i := 0; i < len(nums); i++ {
		l := i + 1         // 第二个数字
		r := len(nums) - 1 // 第三个数字
		// 第一个数字去重
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		for l < r {
			// 找个解
			if nums[i]+nums[l]+nums[r] == 0 {
				// 先去重(找最中间的数）
				for l < r && nums[l] == nums[l+1] {
					l++
				}
				for l < r && nums[r] == nums[r-1] {
					r--
				}
				res = append(res, []int{nums[i], nums[l], nums[r]})
				l++
				r--
			} else if nums[i]+nums[l]+nums[r] > 0 {
				r-- // 太大了
			} else if nums[i]+nums[l]+nums[r] < 0 {
				l++ // 太小了
			}
		}
	}
	return res
}
