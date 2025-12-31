# 分发更多题解和测试用例

# 015_three_sum
$content = @'
package threesum

import "sort"

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
				l++
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
'@
Set-Content -Path "problems/015_three_sum/solution.go" -Value $content -Encoding UTF8
Write-Host "Created/Updated problems/015_three_sum/solution.go"

$testContent = @'
package threesum

import (
	"reflect"
	"testing"
)

func TestThreeSum(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want [][]int
	}{
		{
			name: "Example_1",
			nums: []int{-1, 0, 1, 2, -1, -4},
			want: [][]int{{-1, -1, 2}, {-1, 0, 1}},
		},
		{
			name: "Example_2",
			nums: []int{0, 1, 1},
			want: [][]int{},
		},
		{
			name: "Example_3",
			nums: []int{0, 0, 0},
			want: [][]int{{0, 0, 0}},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := threeSum(tt.nums)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("threeSum() = %v, want %v", got, tt.want)
			}
		})
	}
}
'@
Set-Content -Path "problems/015_three_sum/solution_test.go" -Value $testContent -Encoding UTF8
Write-Host "Created/Updated problems/015_three_sum/solution_test.go"

# 042_trapping_rain_water
$content = @'
package trappingrainwater

func trap(height []int) int {
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

	// 两个指针都向中间靠拢
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
'@
Set-Content -Path "problems/042_trapping_rain_water/solution.go" -Value $content -Encoding UTF8
Write-Host "Created/Updated problems/042_trapping_rain_water/solution.go"

$testContent = @'
package trappingrainwater

import "testing"

func TestTrap(t *testing.T) {
	tests := []struct {
		name   string
		height []int
		want   int
	}{
		{
			name:   "Example_1",
			height: []int{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1},
			want:   6,
		},
		{
			name:   "Example_2",
			height: []int{4, 2, 0, 3, 2, 5},
			want:   9,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := trap(tt.height); got != tt.want {
				t.Errorf("trap() = %v, want %v", got, tt.want)
			}
		})
	}
}
'@
Set-Content -Path "problems/042_trapping_rain_water/solution_test.go" -Value $testContent -Encoding UTF8
Write-Host "Created/Updated problems/042_trapping_rain_water/solution_test.go"

# 003_longest_substring_without_repeating
$content = @'
package longestsubstring

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
'@
Set-Content -Path "problems/003_longest_substring_without_repeating/solution.go" -Value $content -Encoding UTF8
Write-Host "Created/Updated problems/003_longest_substring_without_repeating/solution.go"

$testContent = @'
package longestsubstring

import "testing"

func TestLengthOfLongestSubstring(t *testing.T) {
	tests := []struct {
		name string
		s    string
		want int
	}{
		{
			name: "Example_1",
			s:    "abcabcbb",
			want: 3,
		},
		{
			name: "Example_2",
			s:    "bbbbb",
			want: 1,
		},
		{
			name: "Example_3",
			s:    "pwwkew",
			want: 3,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := lengthOfLongestSubstring(tt.s); got != tt.want {
				t.Errorf("lengthOfLongestSubstring() = %v, want %v", got, tt.want)
			}
		})
	}
}
'@
Set-Content -Path "problems/003_longest_substring_without_repeating/solution_test.go" -Value $testContent -Encoding UTF8
Write-Host "Created/Updated problems/003_longest_substring_without_repeating/solution_test.go"

# 076_minimum_window_substring
$content = @'
package minimumwindow

func minWindow(s string, t string) string {
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
'@
Set-Content -Path "problems/076_minimum_window_substring/solution.go" -Value $content -Encoding UTF8
Write-Host "Created/Updated problems/076_minimum_window_substring/solution.go"

$testContent = @'
package minimumwindow

import "testing"

func TestMinWindow(t *testing.T) {
	tests := []struct {
		name string
		s    string
		t    string
		want string
	}{
		{
			name: "Example_1",
			s:    "ADOBECODEBANC",
			t:    "ABC",
			want: "BANC",
		},
		{
			name: "Example_2",
			s:    "a",
			t:    "a",
			want: "a",
		},
		{
			name: "Example_3",
			s:    "a",
			t:    "aa",
			want: "",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := minWindow(tt.s, tt.t); got != tt.want {
				t.Errorf("minWindow() = %v, want %v", got, tt.want)
			}
		})
	}
}
'@
Set-Content -Path "problems/076_minimum_window_substring/solution_test.go" -Value $testContent -Encoding UTF8
Write-Host "Created/Updated problems/076_minimum_window_substring/solution_test.go"

# 053_maximum_subarray
$content = @'
package maximumsubarray

import "math"

func maxSubArray(nums []int) int {

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
'@
Set-Content -Path "problems/053_maximum_subarray/solution.go" -Value $content -Encoding UTF8
Write-Host "Created/Updated problems/053_maximum_subarray/solution.go"

$testContent = @'
package maximumsubarray

import "testing"

func TestMaxSubArray(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want int
	}{
		{
			name: "Example_1",
			nums: []int{-2, 1, -3, 4, -1, 2, 1, -5, 4},
			want: 6,
		},
		{
			name: "Example_2",
			nums: []int{1},
			want: 1,
		},
		{
			name: "Example_3",
			nums: []int{5, 4, -1, 7, 8},
			want: 23,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxSubArray(tt.nums); got != tt.want {
				t.Errorf("maxSubArray() = %v, want %v", got, tt.want)
			}
		})
	}
}
'@
Set-Content -Path "problems/053_maximum_subarray/solution_test.go" -Value $testContent -Encoding UTF8
Write-Host "Created/Updated problems/053_maximum_subarray/solution_test.go"

Write-Host "`nAll files created/updated successfully!"
