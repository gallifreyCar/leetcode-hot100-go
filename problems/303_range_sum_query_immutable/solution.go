package range_sum_query_immutable

// NumArray 前缀和
type NumArray struct {
	prefix []int
}

// Constructor 构造函数
func Constructor(nums []int) NumArray {
	n := NumArray{}
	n.prefix = make([]int, len(nums)+1, len(nums)+1)
	for i := 0; i < len(nums); i++ {
		n.prefix[i+1] = nums[i] + n.prefix[i]
	}
	return n
}

// SumRange 区域和检索
// 时间复杂度: O(1)
// 空间复杂度: O(n)
func (this *NumArray) SumRange(left int, right int) int {
	return this.prefix[right+1] - this.prefix[left]
}
