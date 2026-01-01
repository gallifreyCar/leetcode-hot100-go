package searchinsertposition

// 35. 搜索插入位置
// 难度：简单
// 标签：二分查找
// 链接：https://leetcode.cn/problems/search_insert_position/
func searchInsert(nums []int, target int) int {

	left := 0
	right := len(nums) - 1
	for left <= right {
		mid := left + (right-left)/2
		if nums[mid] == target {
			return mid
		} else if nums[mid] < target {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	return left
}
