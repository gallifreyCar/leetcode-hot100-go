package containerwithmostwater

// 11. 盛最多水的容器
// 难度：中等
// 标签：双指针、贪心
// 链接：https://leetcode.cn/problems/container_with_most_water/
// MaxArea 双指针解法
// 时间复杂度: O(n)
// 空间复杂度: O(1)
func MaxArea(height []int) int {
	res := 0
	l := 0
	r := len(height) - 1
	for l < r {
		if height[l] < height[r] {
			res = max(res, height[l]*(r-l))
			l++
		} else {
			res = max(res, height[r]*(r-l))
			r--
		}
	}
	return res
}
