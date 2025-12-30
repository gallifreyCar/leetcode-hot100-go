package trappingrainwater

// 42. 接雨水
// 难度：困难
// 标签：双指针、动态规划、栈
// 链接：https://leetcode.cn/problems/trapping_rain_water/

// Trap 朴素解法：每列储水高度只要考虑每列的最左和最右的储水高度
// 时间复杂度: O(n²)
// 空间复杂度: O(1)
func Trap(height []int) int {
	res := 0
	//最左和最右存储不了水，直接从 1遍历到 len(height)-2
	for i := 1; i < len(height)-1; i++ {
		maxLeft := 0
		maxRight := 0
		cur := height[i]
		// 向左遍历找出除自己外最高的柱子
		for j := i - 1; j >= 0; j-- {
			maxLeft = max(maxLeft, height[j])
		}
		// 向右遍历找出除自己外最高的柱子
		for k := i + 1; k <= len(height)-1; k++ {
			maxRight = max(maxRight, height[k])
		}
		// 两边的最高柱均高于当前柱子高度，该列可储水
		if min(maxLeft, maxRight) > cur {
			res += min(maxLeft, maxRight) - cur
		}
	}
	return res
}

// TrapV2 动态规划优化：预先计算每个位置的左右最大值
// 时间复杂度: O(n)
// 空间复杂度: O(n)
func TrapV2(height []int) int {
	res := 0
	LeftMax := make([]int, len(height))
	RightMax := make([]int, len(height))
	LeftMax[0] = 0
	RightMax[len(height)-1] = 0

	for i := 1; i < len(height)-1; i++ {
		LeftMax[i] = max(LeftMax[i-1], height[i-1])
	}
	for i := len(height) - 2; i > 0; i-- {
		RightMax[i] = max(RightMax[i+1], height[i+1])
	}
	//最左和最右存储不了水
	for i := 1; i < len(height)-1; i++ {
		cur := height[i]
		minHeight := min(LeftMax[i], RightMax[i])
		if minHeight > cur {
			res += minHeight - cur
		}
	}
	return res
}

// TrapV3 双指针优化：空间优化到O(1)
// 时间复杂度: O(n)
// 空间复杂度: O(1)
func TrapV3(height []int) int {
	res := 0
	LeftMax := 0
	RightMax := 0
	l := 1
	r := len(height) - 2

	for l <= r {
		if height[l-1] < height[r+1] {
			LeftMax = max(LeftMax, height[l-1])
			cur := height[l]
			if LeftMax > cur {
				res += LeftMax - cur
			}
			l++
		} else {
			RightMax = max(RightMax, height[r+1])
			cur := height[r]
			if RightMax > cur {
				res += RightMax - cur
			}
			r--
		}
	}
	return res
}
