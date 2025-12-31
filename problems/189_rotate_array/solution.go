package rotatearray

import "slices"

func rotate(nums []int, k int) {
	k %= len(nums)
	//涓夋鍙嶈浆锛屽師鍦扮畻娉?
	slices.Reverse(nums)
	slices.Reverse(nums[:k])
	slices.Reverse(nums[k:])
}
