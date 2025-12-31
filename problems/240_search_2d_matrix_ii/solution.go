package search2dmatrixii

// 240. 鎼滅储浜岀淮鐭╅樀 II
// 闅惧害锛氫腑绛?
// 鏍囩锛氭暟缁勩€佷簩鍒嗘煡鎵俱€佺煩闃?
// 閾炬帴锛歨ttps://leetcode.cn/problems/search-a-2d-matrix-ii/
// SearchMatrix 浠庡彸涓婅寮€濮嬫悳绱?
// 鏃堕棿澶嶆潅搴? O(m+n)
// 绌洪棿澶嶆潅搴? O(1)
func SearchMatrix(matrix [][]int, target int) bool {
	m := len(matrix)
	n := len(matrix[0])

	j := n - 1
	i := 0
	for i < m && j >= 0 {
		//浠庡彸涓婅寮€濮嬮亶鍘?
		if matrix[i][j] == target {
			return true
		}
		if matrix[i][j] > target {
			j--
		} else if matrix[i][j] < target {
			i++
		}
	}
	return false
}
