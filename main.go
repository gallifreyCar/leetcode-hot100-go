package main

import (
	"math"
	"slices"
	"sort"
)

//TIP <p>To run your code, right-click the code and select <b>Run</b>.</p> <p>Alternatively, click
// the <icon src="AllIcons.Actions.Execute"/> icon in the gutter and select the <b>Run</b> menu item from here.</p>

func threeSum(nums []int) [][]int {

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
				l--
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

func trap(height []int) int {
	// 朴素解法一  每列储水高度只要考虑每列的最左和最右的储水高度就行
	// LeftMax，cur,RightMax waterHeight= if(max(LeftMax,RightMax)>cur) max()-cur

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

func trapV2(height []int) int {
	// 朴素解法思路： 每列储水高度只要考虑每列的左边和右边的储水高度就行
	// LeftMax,cur,RightMax
	// curWaterHeight= if(min(LeftMax,RightMax)>cur) min-cur
	// 动态规划: 不用每次遍历里面再for找LeftMax，RightMax
	// 可以只遍历一次把结果用数组存储下来

	res := 0
	LeftMax := make([]int, len(height))
	RightMax := make([]int, len(height))
	LeftMax[0] = 0              // 最左边柱子的除自己外的左边柱子最高高度不存在
	RightMax[len(height)-1] = 0 //右边同理同意

	for i := 1; i < len(height)-1; i++ {
		LeftMax[i] = max(LeftMax[i-1], height[i-1])
	}
	for i := len(height) - 2; i > 0; i-- {
		RightMax[i] = max(RightMax[i+1], height[i+1])
	}
	//最左和最右存储不了水，直接从 1遍历到 len(height)-2
	for i := 1; i < len(height)-1; i++ {
		cur := height[i]
		minHeight := min(LeftMax[i], RightMax[i])
		// 两边的最高柱均高于当前柱子高度，该列可储水（最低一方也高于当前柱高度）
		if minHeight > cur {
			res += minHeight - cur
		}
	}
	return res
}

func trapV3(height []int) int {
	// 朴素解法思路： 每列储水高度只要考虑每列的左边和右边的储水高度就行
	// LeftMax,cur,RightMax
	// curWaterHeight= if(min(LeftMax,RightMax)>cur) min-cur
	// 动态规划: 不用每次遍历里面再for找LeftMax，RightMax
	// 可以只遍历一次把结果用数组存储下来
	// 双指针是动态规划上优化， 把左右两边最高高度位置用两个指针存储就可以了

	res := 0
	LeftMax := 0
	RightMax := 0
	l := 1               //左指针 最左和最右存储不了水，直接从 1遍历到 len(height)-2
	r := len(height) - 2 //右指针

	// 两个指针都香中间靠拢
	for l <= r {
		// 重要！两边同时开工 （l指针）最左iLeftMax是必定小于等于（r指针）最右jLeftMax（最右右的左拥有更大的范围）
		// 反过来 （l指针）最左的右jRightMax 必须大于等于（r指针）最右jRightMax
		// 现在我们只要考虑两边的短板的就可以了（也就是最左(l指针)的iLeftMax和最右(r指针)的jRightMax
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

func lengthOfLongestSubstring(s string) int {
	// 滑动窗口
	// start不动，end向后移动
	// 当end遇到重复字符，start应该放在上一个重复字符的位置的后一位，同时记录最长的长度
	// key判断是否重复，value用找到重复字符的下一个位置

	res := 0
	m := make(map[byte]int)
	for start, end := 0, 0; end < len(s); end++ {
		c := s[end]
		//如果当前字符已经出现过，而且在我们窗口内，则为视为重复字符
		if loc, ok := m[c]; ok && loc >= start {
			start = loc + 1
		}
		m[c] = end
		res = max(res, end-start+1)
	}
	return res
}

func lengthOfLongestSubstringV(s string) int {
	res := 0
	m := make(map[byte]int)
	for start, end := 0, 0; end < len(s); end++ {
		c := s[end]
		if loc, ok := m[c]; ok && loc >= start {
			start = loc + 1 //边界右移
		}
		m[c] = end
		res = max(res, end-start+1)
	}
	return res
}

func findAnagrams(s string, p string) []int {
	res := make([]int, 0)
	start := 0
	need := [26]int{}
	window := [26]int{}
	for i := 0; i < len(p); i++ {
		need[p[i]-'a']++
	}

	for end := 0; end < len(s); end++ {
		//固定窗口的板子（通用三步走）
		// 1. 加入新元素
		window[s[end]-'a']++

		// 2. 窗口长度超过need固定长度，则收缩左边界
		if end-start+1 > len(p) {
			window[s[start]-'a']--
			start++
		}

		// 3. 窗口长度恰好等于need固定长度且符合need条件，处理结果
		if end-start+1 == len(p) && need == window {
			res = append(res, start)
		}
	}
	return res
}

func checkInclusion(s1 string, s2 string) bool {
	res := false
	need := [26]int{}
	window := [26]int{}

	for i := 0; i < len(s1); i++ {
		need[s1[i]-'a']++
	}

	for start, end := 0, 0; end < len(s2); end++ {
		//固定窗口的板子（通用三步走）
		// 1. 加入新元素
		window[s2[end]-'a']++
		// 2. 窗口长度超过need固定长度，则收缩左边界
		if end-start+1 > len(s1) {
			window[s2[end]-'a']--
			start++
		}
		// 3. 窗口长度恰好等于need固定长度且符合need条件，处理结果
		if end-start+1 == len(s1) && need == window {
			res = true
		}
	}
	return res

}

func findMaxAverage(nums []int, k int) float64 {

	var windowSum int
	res := math.Inf(-1)

	for start, end := 0, 0; end < len(nums); end++ {
		//固定窗口的板子（通用三步走）
		// 1. 加入新元素
		windowSum += nums[end]

		// 2. 窗口长度超过need固定长度，则收缩左边界
		if end-start+1 > k {
			windowSum -= nums[start]
			start++
		}
		// 3. 窗口长度恰好等于need固定长度且符合need条件，处理结果
		if end-start+1 == k {
			res = max(res, float64(windowSum)/float64(k))
		}
	}
	return res
}

func longestOnes(nums []int, k int) int {
	start, res := 0, 0
	zeroCount := 0

	for end := 0; end < len(nums); end++ {
		// 1. 加入新元素
		if nums[end] == 0 {
			zeroCount++
		}

		// 2. 收缩左边界（窗口不合法时）
		for zeroCount > k {
			if nums[start] == 0 {
				zeroCount--
			}
			start++
		}

		// 3. 更新结果
		res = max(res, end-start+1)
	}

	return res
}

func lengthOfLongestSubstringV2(s string) int {

	// 滑动窗口（可变窗口）板子
	res := 0
	window := make(map[byte]int)
	for start, end := 0, 0; end < len(s); end++ {
		// 1.新增元素
		c := s[end]
		window[c]++
		// 2.判断是否需要收缩
		for window[c] > 1 {
			window[s[start]]--
			start++
		}
		// 3.更新结果
		res = max(res, end-start+1)

	}
	return res
}

func maxSlidingWindow(nums []int, k int) []int {

	var deque []int                      //单调队列 deque[0]永远最大 value记录num的下标
	res := make([]int, 0, len(nums)-k+1) //结果

	for i := 0; i < len(nums); i++ {
		// 1. 加入新元素
		// 这里加入前，要维护单调队列的单调性，删除所有队列中小于新进入数的数
		for len(deque) > 0 && nums[deque[len(deque)-1]] <= nums[i] {
			deque = deque[:len(deque)-1]
		}
		deque = append(deque, i) //加入新数字的下标

		// 2. 收缩左边界（窗口不合法时）
		// 在这题里面就是队列头要超出窗口了
		left := i - k + 1
		if deque[0] < left {
			deque = deque[1:]
		}

		// 3. 更新结果
		// left>=0 才能形成窗口
		if left >= 0 {
			res = append(res, nums[deque[0]])
		}
	}

	return res
}

type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	listRes := &ListNode{}
	cur := listRes
	for list1 != nil && list2 != nil {
		if list1.Val < list2.Val {
			cur.Next = list1
			list1 = list1.Next
		} else {
			cur.Next = list2
			list2 = list2.Next
		}
		cur = cur.Next
	}
	if list1 != nil {
		cur.Next = list1
	}
	if list2 != nil {
		cur.Next = list2
	}

	return listRes.Next
}

// 判断S不是覆盖了T
func isCovered(cntS, cntT []int) bool {
	for i := 'A'; i <= 'Z'; i++ {
		if cntS[i] < cntT[i] {
			return false
		}
	}
	for i := 'a'; i <= 'z'; i++ {
		if cntS[i] < cntT[i] {
			return false
		}
	}
	return true
}

func minWindow(s string, t string) string {
	cntT := [128]int{} //统计T
	cntW := [128]int{} //统计窗口的

	for _, tc := range t {
		cntT[tc]++
	}
	res := ""
	resLen := len(s) + 1

	for start, end := 0, 0; end < len(s); end++ {
		//1.新增元素
		cntW[s[end]]++
		//2.收缩窗口
		for isCovered(cntW[:], cntT[:]) {
			//更新结果
			if end-start+1 < resLen {
				res = s[start : end+1]
				resLen = end - start + 1
			}
			//收缩边界
			cntW[s[start]]--
			start++
		}
	}

	//找不到
	if resLen == len(s)+1 {
		return ""
	}
	return res

}

func minWindowV2(s string, t string) string {
	cntT := [128]int{}   // 统计T
	cntW := [128]int{}   // 统计窗口的
	needKinds := 0       // 需要的字符种类数
	validKinds := 0      // 符合条件的字符种类数
	res := ""            // 结果
	resLen := len(s) + 1 // 用于判断最短条件

	// T 中需要满足的字符种类数
	for i := 0; i < len(t); i++ {
		// 某个字符第一次出现 needKinds++
		if cntT[t[i]] == 0 {
			needKinds++
		}
		cntT[t[i]]++
	}

	for start, end := 0, 0; end < len(s); end++ {
		//1.新增元素
		cntW[s[end]]++
		// 需要该字符且数量也达标
		if cntT[s[end]] > 0 && cntT[s[end]] == cntW[s[end]] {
			//有一个字符种类数量达标 +1
			//比如 A需要2个，窗口内刚好有2个，达标
			validKinds++
		}

		//2.收缩窗口 （所有字符种类需要的数量都满足要求）
		for validKinds == needKinds {
			//更新结果
			if end-start+1 < resLen {
				res = s[start : end+1]
				resLen = end - start + 1
			}

			//收缩边界
			//预判断，如果窗口左边界刚好是需要的字符种类且数量达标，出窗口后，达标字符种类数-1
			if cntT[s[start]] > 0 && cntT[s[start]] == cntW[s[start]] {
				//窗口左边界字母减少后，可能就不达标了
				validKinds--
			}
			cntW[s[start]]--
			start++
		}
	}

	//找不到
	if resLen == len(s)+1 {
		return ""
	}
	return res

}

func minWindowV3(s string, t string) string {
	need := [128]int{}   // 统计T
	window := [128]int{} // 统计窗口的
	needKinds := 0       // 需要的字符种类数
	validKinds := 0      // 符合条件的字符种类数
	res := ""            // 结果
	resLen := len(s) + 1 // 用于判断最短条件

	// T 中需要满足的字符种类数
	for _, c := range t {
		// 某个字符第一次出现 needKinds++
		if need[c] == 0 {
			needKinds++
		}
		need[c]++
	}

	start := 0
	for end, ec := range s {
		//1.新增元素
		window[ec]++
		// 需要该字符且数量也达标
		if need[ec] > 0 && need[ec] == window[ec] {
			//有一个字符种类数量达标 +1
			//比如 A需要2个，窗口内刚好有2个，达标
			validKinds++
		}

		//2.收缩窗口 （所有字符种类需要的数量都满足要求）
		for validKinds == needKinds {
			//更新结果
			if end-start+1 < resLen {
				res = s[start : end+1]
				resLen = end - start + 1
			}
			sc := s[start]
			//收缩边界
			//预判断，如果窗口左边界刚好是需要的字符种类且数量达标，出窗口后，达标字符种类数-1
			if need[sc] > 0 && need[sc] == window[sc] {
				//窗口左边界字母减少后，可能就不达标了
				validKinds--
			}
			window[sc]--
			start++
		}
	}

	//找不到
	if resLen == len(s)+1 {
		return ""
	}
	return res

}

func minSubArrayLen(target int, nums []int) int {
	res := len(nums) + 1 //最后 res不可能大于len(nums)
	sum := 0
	for start, end := 0, 0; end < len(nums); end++ {
		//1.新增元素
		sum += nums[end]
		//2.收缩窗口
		for sum >= target {

			// 3.更新结果（这题要收缩前更新结果）
			res = min(res, end-start+1)

			sum -= nums[start]
			start++
		}
	}
	// 没找到答案
	if res == len(nums)+1 {
		return 0
	}

	return res
}

func maxSubArray(nums []int) int {
	res := nums[0]
	sum := 0

	for start, end := 0, 0; end < len(nums); end++ {
		//1.新增元素
		sum += nums[end]

		//2.记录结果
		res = max(res, sum)
		res = max(res, nums[end])

		//3.缩小区间
		for sum < 0 {
			sum -= nums[start]
			start++
		}
	}

	return res

}

func maxSubArrayV2(nums []int) int {

	prefixMin := 0
	prefixSum := 0
	res := math.MinInt32

	//算前缀,计算当前的数-之前的最小值
	for i := 0; i < len(nums); i++ {
		prefixSum += nums[i]
		res = max(res, prefixSum-prefixMin)
		prefixMin = min(prefixMin, prefixSum)
	}

	return res
}

func maxSubArrayV3(nums []int) int {

	dpSum := nums[0]
	res := nums[0]

	//动态规划 当前元素接前面还是重新开始
	for i := 1; i < len(nums); i++ {
		dpSum = max(nums[i], dpSum+nums[i])
		println(dpSum)
		res = max(res, dpSum)
	}

	return res
}

//func merge(intervals [][]int) [][]int {
//	//1.先按照区左端点进行排序
//	sort.Slice(intervals, func(i, j int) bool {
//		return intervals[i][0] < intervals[j][0]
//	})
//	//2.遍历所有区间
//	res := [][]int{intervals[0]} //把第一个塞进先
//	for i := 1; i < len(intervals); i++ {
//		cur := intervals[i]     //当前区间
//		last := res[len(res)-1] //结果里面最后一个合并区间
//		if last[1] >= cur[0] {
//			//重合了 直接更新
//			last[1] = max(last[1], cur[1])
//		} else {
//			//不重合
//			res = append(res, cur)
//		}
//	}
//	return res
//}

func rotate(nums []int, k int) {
	newNums := make([]int, len(nums))
	for i := 0; i < len(nums); i++ {
		t := (i + k) % len(nums)
		newNums[t] = nums[i]
	}
	copy(nums, newNums)
}

func rotateV2(nums []int, k int) {
	k %= len(nums)
	//三次反转，原地算法
	slices.Reverse(nums)
	slices.Reverse(nums[:k])
	slices.Reverse(nums[k:])

}

func reverse(nums []int, l, r int) []int {
	for l < r {
		nums[l], nums[r] = nums[r], nums[l]
		l++
		r--
	}
	return nums
}

func productExceptSelf(nums []int) []int {

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

func setZeroes(matrix [][]int) {
	//用map记录i,j是否是0
	rowM := make(map[int]bool, len(matrix))
	colM := make(map[int]bool, len(matrix[0]))

	//记录原矩阵的0
	for i, row := range matrix {
		for j, v := range row {
			if v == 0 {
				rowM[i] = true
				colM[j] = true
			}
		}
	}

	//原地算法
	for i, row := range matrix {
		for j := range row {
			if rowM[i] || colM[j] {
				//该数的同列/同行存在0则置为0
				matrix[i][j] = 0
			}
		}
	}

}

func setZeroesV2(matrix [][]int) {
	//类型excel 用第一列和第一行来存储0信息

	//记录第一列和第一行的是否存在0
	rowF := false
	colF := false

	for _, col := range matrix[0] {
		if col == 0 {
			rowF = true
			break
		}
	}

	for _, row := range matrix {
		if row[0] == 0 {
			colF = true
			break
		}
	}

	//记录0 跳过第一行/列的判断
	for i := 1; i < len(matrix); i++ {
		for j := 1; j < len(matrix[i]); j++ {
			if matrix[i][j] == 0 {
				matrix[0][j] = 0
				matrix[i][0] = 0
			}
		}
	}

	//原地算法
	for i := 1; i < len(matrix); i++ {
		for j := 1; j < len(matrix[i]); j++ {
			if matrix[i][0] == 0 || matrix[0][j] == 0 {
				matrix[i][j] = 0
			}
		}
	}

	if rowF {
		clear(matrix[0])
	}
	if colF {
		for _, row := range matrix {
			row[0] = 0
		}
	}

}

func main() {
	//maxSlidingWindow([]int{1, 3, -1, -3, 5, 3, 6, 7}, 3)
	//
	//println(minSubArrayLen(11, []int{1, 2, 3, 4, 5}))

	// s = "ADOBECODEBANC", t = "ABC
	//minWindow("ADOBECODEBANC", "ABC")

	//nums := []int{1, 3, -1, 2, -3, 5, 3, 6, 7}
	//println(maxSubArrayV3(nums))
	//[[1,6],[8,10],[15,18]]
	//merge([][]int{{1, 3}, {2, 6}})
	//rotate([]int{1, 2, 3, 4, 5, 6, 7}, 3)

	productExceptSelf([]int{1, 2, 0})

	firstMissingPositive([]int{1, 2, 0})

	setZeroesV2([][]int{{0, 1, 2, 0}, {3, 4, 5, 2}, {1, 3, 1, 5}})

	spiralOrder([][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}})
}

type NumArray struct {
	prefix []int
}

//func Constructor(nums []int) NumArray {
//	n := NumArray{}
//	n.prefix = make([]int, len(nums)+1, len(nums)+1)
//	for i := 0; i < len(nums); i++ {
//		n.prefix[i+1] = nums[i] + n.prefix[i]
//	}
//	return n
//}

func (this *NumArray) SumRange(left int, right int) int {
	return this.prefix[right] - this.prefix[left]
}

func spiralOrder(matrix [][]int) []int {
	cnt := 0 //计数器
	m := len(matrix)
	n := len(matrix[0])

	res := make([]int, 0, m)
	visited := make([][]int, m)
	for k := 0; k < m; k++ {
		visited[k] = make([]int, n)
	}

	i := 0
	j := -1 //从界外出发
	for cnt < n*m {
		// 右-下-左-上
		// 下一个步没越界且没访问过
		for j+1 < n && visited[i][j+1] == 0 {
			j++
			visited[i][j]++
			res = append(res, matrix[i][j])
			cnt++
		}
		for i+1 < m && visited[i+1][j] == 0 {
			i++
			visited[i][j]++
			res = append(res, matrix[i][j])
			cnt++
		}
		for j-1 >= 0 && visited[i][j-1] == 0 {
			j--
			visited[i][j]++
			res = append(res, matrix[i][j])
			cnt++
		}
		for i-1 >= 0 && visited[i-1][j] == 0 {
			i--
			visited[i][j]++
			res = append(res, matrix[i][j])
			cnt++
		}
	}
	return res
}

func spiralOrderV2(matrix [][]int) []int {
	m, n := len(matrix), len(matrix[0])
	left, right, top, bottom := 0, n-1, 0, m-1
	res := make([]int, 0, m)
	//边界法，走完收缩边界
	for {
		// 右
		for i := left; i <= right; i++ {
			res = append(res, matrix[top][i])
		}
		top++
		if top > bottom {
			break
		}

		// 下
		for i := top; i <= bottom; i++ {
			res = append(res, matrix[i][right])
		}
		right--
		if right < left {
			break
		}

		// 左
		for i := right; i >= left; i-- {
			res = append(res, matrix[bottom][i])
		}
		bottom--
		if top > bottom {
			break
		}

		// 上
		for i := bottom; i >= top; i-- {
			res = append(res, matrix[i][left])
		}
		left++
		if right < left {
			break
		}
	}

	return res

}

func rotate48(matrix [][]int) {
	//对角线反转
	for i := 0; i < len(matrix); i++ {
		for j := 0; j < i; j++ {
			matrix[i][j], matrix[j][i] = matrix[j][i], matrix[i][j]
		}
	}
	//左右翻转
	for i := 0; i < len(matrix); i++ {
		left, right := 0, len(matrix[0])-1
		for left < right {
			matrix[i][left], matrix[i][right] = matrix[i][right], matrix[i][left]
			left++
			right--
		}
	}
}

func firstMissingPositive(nums []int) int {
	//原地算法，用下标来标记这个数有没有出现过
	//用正负数来标记（下标i-1对应数字i)

	// 原来的0或负数直接置为用边界外的任意一个正数（不用处理
	for i, num := range nums {
		if num <= 0 {
			//num = len(nums) x 错误值传递不能赋值
			nums[i] = len(nums) + 1
		}
	}

	//如果num的下标num-1存在，则置为负数
	for _, num := range nums {
		value := abs(num) //这里出现的负数是因为前置操作导致的
		if value-1 < len(nums) {
			nums[value-1] = -abs(nums[value-1])
		}
	}

	//存在正数就是缺少的那个数字
	for i, num := range nums {
		if num > 0 {
			return i + 1
		}
	}
	return len(nums) + 1
}

func abs(num int) int {
	if num < 0 {
		return -num
	}
	return num
}

func sortList(head *ListNode) *ListNode {
	//递归结束条件
	if head == nil || head.Next == nil {
		return head
	}
	//1.找中点
	mid := findMid(head) //偶数取左边，奇数取中间
	right := mid.Next
	mid.Next = nil //断开

	//2.归并
	left := sortList(head)
	right = sortList(right)

	//3.合并
	return merge(left, right)

}

func findMid(head *ListNode) *ListNode {
	slow, fast := head, head.Next //fast和slow起点不一样，确保slow停在左起点

	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}
	return slow
}

func merge(left *ListNode, right *ListNode) *ListNode {
	dummy := &ListNode{}
	cur := dummy
	for left != nil && right != nil {
		if left.Val < right.Val {
			cur.Next = left
			left = left.Next
		} else {
			cur.Next = right
			right = right.Next
		}
		cur = cur.Next
	}
	if left != nil {
		cur.Next = left
	}
	if right != nil {
		cur.Next = right
	}
	return dummy.Next
}

func mergeKLists(lists []*ListNode) *ListNode {
	if len(lists) == 0 {
		return nil
	}
	if len(lists) == 1 {
		return lists[0]
	}
	left := mergeKLists(lists[:len(lists)/2])
	right := mergeKLists(lists[len(lists)/2:])
	return merge(left, right)
}

type Node struct {
	key, val   int
	prev, next *Node
}

type LRUCache struct {
	capacity int
	dummy    *Node //虚拟头节点
	tail     *Node //虚拟尾节点
	cache    map[int]*Node
}

func Constructor(capacity int) LRUCache {

	dummy := &Node{}
	tail := &Node{}
	dummy.next = tail
	tail.prev = dummy

	return LRUCache{
		capacity: capacity,
		cache:    make(map[int]*Node),
		dummy:    dummy,
		tail:     tail,
	}
}

func (this *LRUCache) Get(key int) int {
	if node, ok := this.cache[key]; ok {
		this.moveToHead(node)
		return node.val
	}
	return -1
}

// 新元素添加到头节点
func (this *LRUCache) addToHead(node *Node) {
	node.prev = this.dummy
	node.next = this.dummy.next
	this.dummy.next.prev = node
	this.dummy.next = node
}

// 移除节点
func (this *LRUCache) removeNode(node *Node) {
	node.prev.next = node.next
	node.next.prev = node.prev
}

// 移除尾节点(要返回给哈希表删除)
func (this *LRUCache) removeTail() *Node {
	node := this.tail.prev
	this.removeNode(node)
	return node
}

// 移动到头节点
func (this *LRUCache) moveToHead(node *Node) {
	this.removeNode(node)
	this.addToHead(node)
}

func (this *LRUCache) Put(key int, value int) {
	//存在
	if node, ok := this.cache[key]; ok {
		node.val = value
		this.moveToHead(node)
		return
	}

	node := &Node{key: key, val: value}
	this.cache[key] = node
	this.addToHead(node)

	if len(this.cache) > this.capacity {
		removeTail := this.removeTail()
		delete(this.cache, removeTail.key)
	}

}

// TreeNode Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func inorderTraversal(root *TreeNode) []int {
	res := make([]int, 0)
	var dfs func(node *TreeNode)
	dfs = func(node *TreeNode) {
		if node == nil {
			return
		}
		dfs(node.Left)
		res = append(res, node.Val)
		dfs(node.Right)
	}

	dfs(root)
	return res
}

func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	left := maxDepth(root.Left)
	right := maxDepth(root.Right)
	return max(left, right) + 1
}

func invertTree(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	invertTree(root.Left)
	invertTree(root.Right)
	root.Left, root.Right = root.Right, root.Left
	return root
}

func isSymmetric(root *TreeNode) bool {
	if root == nil {
		return true
	}
	var isMirror func(left, right *TreeNode) bool
	isMirror = func(left, right *TreeNode) bool {
		if left == nil && right == nil {
			return true
		}
		if left == nil || right == nil {
			return false
		}
		if left.Val != right.Val {
			return false
		}
		//外面和外面比 && 里面和里面比
		return isMirror(left.Left, right.Right) && isMirror(left.Right, right.Left)
	}
	return isMirror(root.Left, root.Right)
}

func diameterOfBinaryTree(root *TreeNode) int {

	var res int
	var depth func(node *TreeNode) int
	depth = func(node *TreeNode) int {
		if node == nil {
			return 0
		}
		left := depth(node.Left)
		right := depth(node.Right)

		// 更新直径（边数）
		res = max(res, left+right)

		// 返回高度（节点数）
		return max(left, right) + 1
	}
	depth(root)
	return res
}

func sortedArrayToBST(nums []int) *TreeNode {
	if len(nums) == 0 {
		return nil
	}

	mid := len(nums) / 2
	root := &TreeNode{Val: nums[mid],
		Left:  sortedArrayToBST(nums[:mid]),
		Right: sortedArrayToBST(nums[mid+1:]),
	}
	return root
}

func isValidBST(root *TreeNode) bool {
	var pre *TreeNode
	var inorder func(node *TreeNode) bool
	//BST 的中序遍历是严格递增序列
	inorder = func(node *TreeNode) bool {
		if node == nil {
			return true
		}
		//左节点
		if !inorder(node.Left) {
			return false
		}
		//和上一个访问的节点比较，看看是否合法（单调递增）
		if pre != nil && node.Val <= pre.Val {
			return false
		}
		pre = node
		//右节点
		return inorder(node.Right)
	}
	return inorder(root)
}

func kthSmallest(root *TreeNode, k int) int {
	var cnt int
	var res int
	var inorder func(node *TreeNode)
	//BST 的中序遍历是严格递增序列
	inorder = func(node *TreeNode) {

		if node == nil {
			return
		}

		//左节点
		inorder(node.Left)
		//访问
		cnt++
		if cnt == k {
			res = node.Val
			return
		}

		//右节点
		inorder(node.Right)

	}
	inorder(root)
	return res
}

func rightSideView(root *TreeNode) []int {
	if root == nil {
		return nil
	}
	que := make([]*TreeNode, 0)
	que = append(que, root)
	res := make([]int, 0)
	for len(que) > 0 {
		size := len(que)
		tmp := make([]int, 0)
		for i := 0; i < size; i++ {
			cur := que[0] //队头出列
			que = que[1:]
			tmp = append(tmp, cur.Val)
			if cur.Left != nil {
				que = append(que, cur.Left)
			}
			if cur.Right != nil {
				que = append(que, cur.Right)
			}
		}
		res = append(res, tmp[len(tmp)-1]) //只要队尾
	}
	return res
}

func flatten(root *TreeNode) {
	var pre *TreeNode //前驱指针
	var dfs func(node *TreeNode)
	dfs = func(node *TreeNode) {
		if node == nil {
			return
		}
		//要求：先序遍历顺序是 根左右
		//现在要从下往上接，必须反过 右左根
		dfs(node.Right)
		dfs(node.Left)

		node.Right = pre
		node.Left = nil
		pre = node
	}
	dfs(root)
}

func buildTree(preorder []int, inorder []int) *TreeNode {
	if len(preorder) == 0 || len(inorder) == 0 {
		return nil
	}
	// 前序遍历第一个节点是根
	rootVal := preorder[0]
	root := &TreeNode{Val: rootVal}

	// 在中序遍历中找到根的位置（因为无重复值）
	var mid int
	for i, v := range inorder {
		if v == rootVal {
			mid = i
			break
		}
	}

	//中序遍历，根的左边是左节点，根的右边是右节点
	root.Left = buildTree(preorder[1:mid+1], inorder[:mid])
	root.Right = buildTree(preorder[mid+1:], inorder[mid+1:])

	return root
}

func subarraySum(nums []int, k int) int {
	res := 0
	m := make(map[int]int)
	m[0] = 1 //前缀和为0的出现过一次
	prefixSum := 0

	//子数组的和：prefixSum[r] - prefixSum[l-1] = k
	//prefixSum[l-1] = prefixSum[r] - k
	for _, num := range nums {
		prefixSum += num
		res += m[prefixSum-k]
		m[prefixSum]++
	}
	return res
}

func pathSum(root *TreeNode, targetSum int) int {
	m := make(map[int]int)
	m[0] = 1 //前缀和为0的出现过一次
	prefixSum := 0
	res := 0
	var dfs func(root *TreeNode)
	dfs = func(node *TreeNode) {
		if node == nil {
			return
		}

		//子数组的和：prefixSum[r] - prefixSum[l-1] = k
		//prefixSum[l-1] = prefixSum[r] - k
		prefixSum += node.Val
		res += m[prefixSum-targetSum]
		m[prefixSum]++ //次数增加

		//遍历左右节点
		dfs(node.Left)
		dfs(node.Right)

		//回溯 这题只走一条线路
		m[prefixSum]--
		prefixSum -= node.Val
	}
	dfs(root)
	return res
}

func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	//其实最近公共只节点有两种情况
	//一.p和q分别在root左右子树里面，那么root就是公共节点
	//二.p和q在root同一个（左或右）子树里面，那么公告节点就是p或者q

	//递归终止条件：找到p/q或者root为nil
	if root == nil || root == p || root == q {
		return root
	}

	//取左右子树中找p和q
	left := lowestCommonAncestor(root.Left, p, q)
	right := lowestCommonAncestor(root.Right, p, q)

	//情况一
	if left != nil && right != nil {
		return root
	}

	//情况二
	if left != nil {
		return left
	}
	if right != nil {
		return right
	}

	return nil
}

func maxPathSum(root *TreeNode) int {
	//最大路径=当前路径值+左右子树提供的最大“单边路径值”（max(leftGain,RightGain)

	res := math.MinInt32
	var dfs func(node *TreeNode) int
	dfs = func(node *TreeNode) int {
		if node == nil {
			return 0
		}
		//左右子树能提供的最大“单边路径”
		leftGain := max(dfs(node.Left), 0) //负数抛弃
		rightGain := max(dfs(node.Right), 0)
		//当前最大路径值
		curSum := node.Val + leftGain + rightGain
		res = max(res, curSum)
		//返回单边路径
		return node.Val + max(leftGain, rightGain)
	}
	dfs(root)
	return res

}

func numIslands(grid [][]byte) int {
	if len(grid) == 0 || len(grid[0]) == 0 {
		return 0
	}

	m, n := len(grid), len(grid[0])
	cnt := 0
	var dfs func(i, j int)
	dfs = func(i, j int) {
		//越界or不是陆地，直接返回
		if i < 0 || i >= m || j < 0 || j >= n || grid[i][j] != '1' {
			return
		}
		//把陆地变成水
		if grid[i][j] == '1' {
			grid[i][j] = '0'
			cnt++
		}
		//往四个方向走
		dfs(i-1, j)
		dfs(i, j-1)
		dfs(i, j+1)
		dfs(i+1, j)
	}
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == '1' {
				cnt++
				dfs(i, j)
			}

		}
	}
	return cnt
}

func orangesRotting(grid [][]int) int {
	if len(grid) == 0 || len(grid[0]) == 0 {
		return 0
	}

	m, n := len(grid), len(grid[0])
	res := 0
	fresh := 0
	var queue [][2]int //队列，记录烂橘子的坐标
	//统计新鲜橘子+腐烂橘子入队
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == 1 {
				fresh++
			}
			if grid[i][j] == 2 {
				queue = append(queue, [2]int{i, j})
			}
		}
	}
	// 没有新鲜橘子 直接返回
	if fresh == 0 {
		return 0
	}
	//BFS
	dirs := [][2]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}
	for len(queue) > 0 {
		//取出一个烂橘子
		x, y := queue[0][0], queue[0][1]
		queue = queue[1:]

		size := len(queue)
		rotted := false //本轮是否腐烂了橘子

		for i := 0; i < size; i++ {
			//遍历四个方向
			for _, d := range dirs {
				xx, yy := x+d[0], y+d[1]
				if xx >= 0 && xx < m && yy >= 0 && yy < n && grid[xx][yy] == 1 {
					grid[xx][yy] = 2
					rotted = true
					fresh--
					queue = append(queue, [2]int{xx, yy})
				}
			}
		}
		//如果有果子腐烂了
		if rotted {
			res++
		}
	}
	if fresh > 0 {
		return -1
	}

	return res

}

func canFinish(numCourses int, prerequisites [][]int) bool {
	graph := make([][]int, numCourses)  //记录后置课程，用map[int][]int也可以
	indegree := make([]int, numCourses) //入度，用map[int]int也可以
	//建图和记录入度
	for _, prerequisite := range prerequisites {
		//[1,0]为例，学1要学0，所以1入度++
		indegree[prerequisite[0]]++
		// [1,0],[2,0]为例,记录0的出度是什么，graph[0]=[1,2]
		graph[prerequisite[1]] = append(graph[prerequisite[1]], prerequisite[0])
	}
	//入度为0的进队列
	queue := make([]int, 0, numCourses)
	for cur, v := range indegree {
		if v == 0 {
			queue = append(queue, cur)
		}
	}
	//BFS
	cnt := 0
	for len(queue) > 0 {
		cur := queue[0]
		queue = queue[1:]
		cnt++
		//学完一门课，后续课程的入度减少
		for _, next := range graph[cur] {
			indegree[next]--
			if indegree[next] == 0 {
				queue = append(queue, next) //入度为0进队列
			}
		}

	}
	return cnt == numCourses
}

type Trie struct {
	children [26]*Trie //26字母
	isEnd    bool
}

func ConstructorV2() Trie {
	return Trie{children: [26]*Trie{}}
}

func (this *Trie) Insert(word string) {
	node := this
	//从 root 出发，找字符，找不到就建立
	for _, ch := range word {
		i := ch - 'a'
		if node.children[i] == nil {
			node.children[i] = &Trie{}
		}
		node = node.children[i]
	}
	node.isEnd = true //标识结束点
}

func (this *Trie) Search(word string) bool {
	node := this
	for _, ch := range word {
		i := ch - 'a'
		if node.children[i] == nil {
			return false
		}
		node = node.children[i]
	}
	return node.isEnd
}

func (this *Trie) StartsWith(prefix string) bool {
	node := this
	for _, ch := range prefix {
		i := ch - 'a'
		if node.children[i] == nil {
			return false
		}
		node = node.children[i]
	}
	return true
}

func permute(nums []int) [][]int {
	visited := make([]bool, len(nums))
	path := make([]int, 0, len(nums))
	result := make([][]int, 0, len(nums))

	var dfs func()
	dfs = func() {
		if len(path) == len(nums) {
			tmp := append([]int{}, path...) //必须拷贝，因为path是切边，本质是指针
			result = append(result, tmp)
			return
		}

		for i, num := range nums {
			if visited[i] {
				continue //访问过的就不走了
			}
			visited[i] = true //标记访问
			path = append(path, num)
			dfs()
			//回溯
			visited[i] = false
			path = path[:len(path)-1]
		}
	}
	dfs()
	return result
}

func subsets(nums []int) [][]int {
	path := make([]int, 0, len(nums))
	result := make([][]int, 0, len(nums))
	var dfs func(start int)
	dfs = func(start int) {
		tmp := append([]int{}, path...) //必须拷贝，因为path是切边，本质是指针
		result = append(result, tmp)

		for i := start; i < len(nums); i++ {
			path = append(path, nums[i])
			dfs(i + 1)
			path = path[:len(path)-1]
		}
	}
	dfs(0)
	return result
}

func letterCombinations(digits string) []string {

	m := map[byte]string{
		'2': "abc",
		'3': "def",
		'4': "ghi",
		'5': "jkl",
		'6': "mno",
		'7': "pqrs",
		'8': "tuv",
		'9': "wxyz",
	}
	result := make([]string, 0)
	var path []byte
	var dfs func(start int)
	dfs = func(start int) {
		if start == len(digits) {
			result = append(result, string(path))
			return
		}
		letter := m[digits[start]]
		for i := range letter {
			path = append(path, letter[i])
			dfs(start + 1)
			path = path[:len(path)-1]
		}
	}
	dfs(0)
	return result
}

func combinationSum(candidates []int, target int) [][]int {
	path := make([]int, 0)
	res := make([][]int, 0)
	cnt := 0 //计数

	sort.Ints(candidates)

	var dfs func(start int)
	dfs = func(start int) {
		if cnt == target {
			tmp := append([]int{}, path...)
			res = append(res, tmp)
			return
		}
		if cnt > target {
			return
		}
		//剪枝 从start开始
		for i := start; i < len(candidates); i++ {
			//剪枝 递增队列，当前结果已经大于target，没必要向后遍历
			if cnt+candidates[i] > target {
				break
			}
			path = append(path, candidates[i])
			cnt += candidates[i]
			dfs(i)
			path = path[:len(path)-1]
			cnt -= candidates[i]
		}

	}
	dfs(0)
	return res
}

func generateParenthesis(n int) []string {
	res := make([]string, 0)
	path := make([]byte, 0)
	var dfs func(eft int, right int)

	//有效括号的条件:任意前缀中左括号数量大于等于右括号
	dfs = func(left int, right int) {
		//如果数量够了就退出
		if len(path) == 2*n {
			res = append(res, string(path))
		}
		//先选'('
		if left < n {
			path = append(path, byte('('))
			dfs(left+1, right)
			path = path[:len(path)-1]
		}
		//如果')'数量小于'(' 选择')'
		if right < left {
			path = append(path, byte(')'))
			dfs(left, right+1)
			path = path[:len(path)-1]
		}
	}
	dfs(0, 0)
	return res
}

func isValid(s string) bool {
	stack := make([]byte, 0)

	// 右括号 -> 对应左括号
	match := map[byte]byte{
		')': '(',
		'}': '{',
		']': '[',
	}

	for i := range s {
		// 左括号：入栈
		if s[i] == '(' || s[i] == '{' || s[i] == '[' {
			stack = append(stack, s[i])
			continue
		}

		// 右括号：必须匹配
		if len(stack) == 0 || stack[len(stack)-1] != match[s[i]] {
			return false
		}

		// 匹配成功，出栈
		stack = stack[:len(stack)-1]
	}

	return len(stack) == 0
}

func exist(board [][]byte, word string) bool {
	//1.先遍历一次找到起点
	//2.dfs
	m := len(board)
	n := len(board[0])
	//题目 m>1,n>1 len(word)>1 不考虑空的情况

	//上下左右
	dirs := [][2]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}
	var dfs func(i, j, loc int) bool
	dfs = func(i, j, loc int) bool {
		// 当前字符不匹配
		if board[i][j] != word[loc] {
			return false
		}

		if loc == len(word)-1 {
			return true
		}

		tmp := board[i][j]
		board[i][j] = '0'
		for _, v := range dirs {
			nextI := v[0] + i
			nextJ := v[1] + j
			//剪枝
			if nextI < m && nextI >= 0 && nextJ < n && nextJ >= 0 {
				if dfs(nextI, nextJ, loc+1) {
					return true
				}
			}
		}
		board[i][j] = tmp
		return false
	}

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			//首字母判断
			if board[i][j] == word[0] && dfs(i, j, 0) {
				return true
			}
		}
	}

	return false
}
