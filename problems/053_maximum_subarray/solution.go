package maximumsubarray

import "math"

func maxSubArray(nums []int) int {

	prefixMin := 0
	prefixSum := 0
	res := math.MinInt32

	//绠楀墠缂€,璁＄畻褰撳墠鐨勬暟-涔嬪墠鐨勬渶灏忓€?
	for i := 0; i < len(nums); i++ {
		prefixSum += nums[i]
		res = max(res, prefixSum-prefixMin)
		prefixMin = min(prefixMin, prefixSum)
	}

	return res
}
