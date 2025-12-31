package subarraysumequals

// 560. 鍜屼负 K 鐨勫瓙鏁扮粍
// 闅惧害锛氫腑绛?
// 鏍囩锛氭暟缁勩€佸搱甯岃〃銆佸墠缂€鍜?
// 閾炬帴锛歨ttps://leetcode.cn/problems/subarray-sum-equals-k/
// SubarraySum 鍓嶇紑鍜?鍝堝笇琛ㄤ紭鍖栵紙鏈€浼橈級
// 鏃堕棿澶嶆潅搴? O(n)
// 绌洪棿澶嶆潅搴? O(n)
func SubarraySum(nums []int, k int) int {
	// 鍓嶇紑鍜?鍝堝笇琛紙涓ゆ暟涔嬪拰鐨勬€濇兂锛?
	// 浼樺寲鎺夊墠缂€鍜屾暟缁勶紝鐩存帴鐢ㄤ竴涓彉閲忚褰曞綋鍓嶅墠缂€鍜?
	m := make(map[int]int)
	m[0] = 1 // 鍓嶇紑鍜屼负0鍑虹幇涓€娆?
	res := 0
	prefixSum := 0

	for _, num := range nums {
		prefixSum += num
		if v, ok := m[prefixSum-k]; ok {
			res += v
		}
		m[prefixSum]++
	}

	return res
}
