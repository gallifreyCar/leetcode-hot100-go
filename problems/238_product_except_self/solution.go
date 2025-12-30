package productexceptself

// 238. 除自身以外数组的乘积
// 难度：中等
// 标签：数组、前缀和
// 链接：https://leetcode.cn/problems/product_except_self/

// ProductExceptSelf 前缀积和后缀积
// 时间复杂度: O(n)
// 空间复杂度: O(n)
func ProductExceptSelf(nums []int) []int {
	prefix := make([]int, len(nums)+1)
	suffix := make([]int, len(nums))
	res := make([]int, 0, len(nums))
	suffix[len(nums)-1] = 1
	for i := len(nums) - 1; i > 0; i-- {
		suffix[i-1] = suffix[i] * nums[i]
	}

	prefix[0] = 1
	for i := 0; i < len(nums); i++ {
		prefix[i+1] = prefix[i] * nums[i]
		res = append(res, prefix[i]*suffix[i])
	}

	return res
}
