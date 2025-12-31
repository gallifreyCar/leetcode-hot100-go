# 分发更多题解 - 回溯和其他

# 046_permutations
$content = @'
package permutations

func permute(nums []int) [][]int {
	visited := make([]bool, len(nums))
	path := make([]int, 0, len(nums))
	result := make([][]int, 0, len(nums))

	var dfs func()
	dfs = func() {
		if len(path) == len(nums) {
			tmp := append([]int{}, path...) //必须拷贝，因为path是切片，本质是指针
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
'@
Set-Content -Path "problems/046_permutations/solution.go" -Value $content -Encoding UTF8
Write-Host "Created/Updated problems/046_permutations/solution.go"

$testContent = @'
package permutations

import (
	"reflect"
	"sort"
	"testing"
)

func TestPermute(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want [][]int
	}{
		{
			name: "Example_1",
			nums: []int{1, 2, 3},
			want: [][]int{{1, 2, 3}, {1, 3, 2}, {2, 1, 3}, {2, 3, 1}, {3, 1, 2}, {3, 2, 1}},
		},
		{
			name: "Example_2",
			nums: []int{0, 1},
			want: [][]int{{0, 1}, {1, 0}},
		},
		{
			name: "Single",
			nums: []int{1},
			want: [][]int{{1}},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := permute(tt.nums)
			if !equalPerms(got, tt.want) {
				t.Errorf("permute() = %v, want %v", got, tt.want)
			}
		})
	}
}

func equalPerms(a, b [][]int) bool {
	if len(a) != len(b) {
		return false
	}
	sortPerms(a)
	sortPerms(b)
	return reflect.DeepEqual(a, b)
}

func sortPerms(perms [][]int) {
	sort.Slice(perms, func(i, j int) bool {
		for k := 0; k < len(perms[i]) && k < len(perms[j]); k++ {
			if perms[i][k] != perms[j][k] {
				return perms[i][k] < perms[j][k]
			}
		}
		return len(perms[i]) < len(perms[j])
	})
}
'@
Set-Content -Path "problems/046_permutations/solution_test.go" -Value $testContent -Encoding UTF8
Write-Host "Created/Updated problems/046_permutations/solution_test.go"

# 078_subsets
$content = @'
package subsets

func subsets(nums []int) [][]int {
	path := make([]int, 0, len(nums))
	result := make([][]int, 0, len(nums))
	var dfs func(start int)
	dfs = func(start int) {
		tmp := append([]int{}, path...) //必须拷贝，因为path是切片，本质是指针
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
'@
Set-Content -Path "problems/078_subsets/solution.go" -Value $content -Encoding UTF8
Write-Host "Created/Updated problems/078_subsets/solution.go"

$testContent = @'
package subsets

import (
	"reflect"
	"sort"
	"testing"
)

func TestSubsets(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want [][]int
	}{
		{
			name: "Example_1",
			nums: []int{1, 2, 3},
			want: [][]int{{}, {1}, {1, 2}, {1, 2, 3}, {1, 3}, {2}, {2, 3}, {3}},
		},
		{
			name: "Example_2",
			nums: []int{0},
			want: [][]int{{}, {0}},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := subsets(tt.nums)
			if !equalSubsets(got, tt.want) {
				t.Errorf("subsets() = %v, want %v", got, tt.want)
			}
		})
	}
}

func equalSubsets(a, b [][]int) bool {
	if len(a) != len(b) {
		return false
	}
	sortSubsets(a)
	sortSubsets(b)
	return reflect.DeepEqual(a, b)
}

func sortSubsets(sets [][]int) {
	for _, s := range sets {
		sort.Ints(s)
	}
	sort.Slice(sets, func(i, j int) bool {
		if len(sets[i]) != len(sets[j]) {
			return len(sets[i]) < len(sets[j])
		}
		for k := 0; k < len(sets[i]); k++ {
			if sets[i][k] != sets[j][k] {
				return sets[i][k] < sets[j][k]
			}
		}
		return false
	})
}
'@
Set-Content -Path "problems/078_subsets/solution_test.go" -Value $testContent -Encoding UTF8
Write-Host "Created/Updated problems/078_subsets/solution_test.go"

# 017_letter_combinations
$content = @'
package lettercombinations

func letterCombinations(digits string) []string {
	if len(digits) == 0 {
		return []string{}
	}

	m := map[byte]string{
		'"'"'2'"'"': "abc",
		'"'"'3'"'"': "def",
		'"'"'4'"'"': "ghi",
		'"'"'5'"'"': "jkl",
		'"'"'6'"'"': "mno",
		'"'"'7'"'"': "pqrs",
		'"'"'8'"'"': "tuv",
		'"'"'9'"'"': "wxyz",
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
'@
Set-Content -Path "problems/017_letter_combinations/solution.go" -Value $content -Encoding UTF8
Write-Host "Created/Updated problems/017_letter_combinations/solution.go"

$testContent = @'
package lettercombinations

import (
	"reflect"
	"sort"
	"testing"
)

func TestLetterCombinations(t *testing.T) {
	tests := []struct {
		name   string
		digits string
		want   []string
	}{
		{
			name:   "Example_1",
			digits: "23",
			want:   []string{"ad", "ae", "af", "bd", "be", "bf", "cd", "ce", "cf"},
		},
		{
			name:   "Empty",
			digits: "",
			want:   []string{},
		},
		{
			name:   "Single",
			digits: "2",
			want:   []string{"a", "b", "c"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := letterCombinations(tt.digits)
			sort.Strings(got)
			sort.Strings(tt.want)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("letterCombinations() = %v, want %v", got, tt.want)
			}
		})
	}
}
'@
Set-Content -Path "problems/017_letter_combinations/solution_test.go" -Value $testContent -Encoding UTF8
Write-Host "Created/Updated problems/017_letter_combinations/solution_test.go"

# 039_combination_sum
$content = @'
package combinationsum

import "sort"

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
'@
Set-Content -Path "problems/039_combination_sum/solution.go" -Value $content -Encoding UTF8
Write-Host "Created/Updated problems/039_combination_sum/solution.go"

$testContent = @'
package combinationsum

import (
	"reflect"
	"sort"
	"testing"
)

func TestCombinationSum(t *testing.T) {
	tests := []struct {
		name       string
		candidates []int
		target     int
		want       [][]int
	}{
		{
			name:       "Example_1",
			candidates: []int{2, 3, 6, 7},
			target:     7,
			want:       [][]int{{2, 2, 3}, {7}},
		},
		{
			name:       "Example_2",
			candidates: []int{2, 3, 5},
			target:     8,
			want:       [][]int{{2, 2, 2, 2}, {2, 3, 3}, {3, 5}},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := combinationSum(tt.candidates, tt.target)
			if !equalCombinations(got, tt.want) {
				t.Errorf("combinationSum() = %v, want %v", got, tt.want)
			}
		})
	}
}

func equalCombinations(a, b [][]int) bool {
	if len(a) != len(b) {
		return false
	}
	sortCombinations(a)
	sortCombinations(b)
	return reflect.DeepEqual(a, b)
}

func sortCombinations(combs [][]int) {
	for _, c := range combs {
		sort.Ints(c)
	}
	sort.Slice(combs, func(i, j int) bool {
		if len(combs[i]) != len(combs[j]) {
			return len(combs[i]) < len(combs[j])
		}
		for k := 0; k < len(combs[i]); k++ {
			if combs[i][k] != combs[j][k] {
				return combs[i][k] < combs[j][k]
			}
		}
		return false
	})
}
'@
Set-Content -Path "problems/039_combination_sum/solution_test.go" -Value $testContent -Encoding UTF8
Write-Host "Created/Updated problems/039_combination_sum/solution_test.go"

# 022_generate_parentheses
$content = @'
package generateparentheses

func generateParenthesis(n int) []string {
	res := make([]string, 0)
	path := make([]byte, 0)
	var dfs func(left int, right int)

	//有效括号的条件:任意前缀中左括号数量大于等于右括号
	dfs = func(left int, right int) {
		//如果数量够了就退出
		if len(path) == 2*n {
			res = append(res, string(path))
			return
		}
		//先选'"'"'('"'"'
		if left < n {
			path = append(path, byte('"'"'('"'"'))
			dfs(left+1, right)
			path = path[:len(path)-1]
		}
		//如果'"'"')'"'"'数量小于'"'"'('"'"' 选择'"'"')'"'"'
		if right < left {
			path = append(path, byte('"'"')'"'"'))
			dfs(left, right+1)
			path = path[:len(path)-1]
		}
	}
	dfs(0, 0)
	return res
}
'@
Set-Content -Path "problems/022_generate_parentheses/solution.go" -Value $content -Encoding UTF8
Write-Host "Created/Updated problems/022_generate_parentheses/solution.go"

$testContent = @'
package generateparentheses

import (
	"reflect"
	"sort"
	"testing"
)

func TestGenerateParenthesis(t *testing.T) {
	tests := []struct {
		name string
		n    int
		want []string
	}{
		{
			name: "Example_1",
			n:    3,
			want: []string{"((()))", "(()())", "(())()", "()(())", "()()()"},
		},
		{
			name: "Example_2",
			n:    1,
			want: []string{"()"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := generateParenthesis(tt.n)
			sort.Strings(got)
			sort.Strings(tt.want)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("generateParenthesis() = %v, want %v", got, tt.want)
			}
		})
	}
}
'@
Set-Content -Path "problems/022_generate_parentheses/solution_test.go" -Value $testContent -Encoding UTF8
Write-Host "Created/Updated problems/022_generate_parentheses/solution_test.go"

# 079_word_search
$content = @'
package wordsearch

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
		board[i][j] = '"'"'0'"'"'
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
'@
Set-Content -Path "problems/079_word_search/solution.go" -Value $content -Encoding UTF8
Write-Host "Created/Updated problems/079_word_search/solution.go"

$testContent = @'
package wordsearch

import "testing"

func TestExist(t *testing.T) {
	tests := []struct {
		name  string
		board [][]byte
		word  string
		want  bool
	}{
		{
			name: "Example_1",
			board: [][]byte{
				{'"'"'A'"'"', '"'"'B'"'"', '"'"'C'"'"', '"'"'E'"'"'},
				{'"'"'S'"'"', '"'"'F'"'"', '"'"'C'"'"', '"'"'S'"'"'},
				{'"'"'A'"'"', '"'"'D'"'"', '"'"'E'"'"', '"'"'E'"'"'},
			},
			word: "ABCCED",
			want: true,
		},
		{
			name: "Example_2",
			board: [][]byte{
				{'"'"'A'"'"', '"'"'B'"'"', '"'"'C'"'"', '"'"'E'"'"'},
				{'"'"'S'"'"', '"'"'F'"'"', '"'"'C'"'"', '"'"'S'"'"'},
				{'"'"'A'"'"', '"'"'D'"'"', '"'"'E'"'"', '"'"'E'"'"'},
			},
			word: "SEE",
			want: true,
		},
		{
			name: "Example_3",
			board: [][]byte{
				{'"'"'A'"'"', '"'"'B'"'"', '"'"'C'"'"', '"'"'E'"'"'},
				{'"'"'S'"'"', '"'"'F'"'"', '"'"'C'"'"', '"'"'S'"'"'},
				{'"'"'A'"'"', '"'"'D'"'"', '"'"'E'"'"', '"'"'E'"'"'},
			},
			word: "ABCB",
			want: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := exist(tt.board, tt.word); got != tt.want {
				t.Errorf("exist() = %v, want %v", got, tt.want)
			}
		})
	}
}
'@
Set-Content -Path "problems/079_word_search/solution_test.go" -Value $testContent -Encoding UTF8
Write-Host "Created/Updated problems/079_word_search/solution_test.go"

# 020_valid_parentheses
$content = @'
package validparentheses

func isValid(s string) bool {
	stack := make([]byte, 0)

	// 右括号 -> 对应左括号
	match := map[byte]byte{
		'"'"')'"'"': '"'"'('"'"',
		'"'"'}'"'"': '"'"'{'"'"',
		'"'"']'"'"': '"'"'['"'"',
	}

	for i := range s {
		// 左括号：入栈
		if s[i] == '"'"'('"'"' || s[i] == '"'"'{'"'"' || s[i] == '"'"'['"'"' {
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
'@
Set-Content -Path "problems/020_valid_parentheses/solution.go" -Value $content -Encoding UTF8
Write-Host "Created/Updated problems/020_valid_parentheses/solution.go"

$testContent = @'
package validparentheses

import "testing"

func TestIsValid(t *testing.T) {
	tests := []struct {
		name string
		s    string
		want bool
	}{
		{
			name: "Example_1",
			s:    "()",
			want: true,
		},
		{
			name: "Example_2",
			s:    "()[]{}",
			want: true,
		},
		{
			name: "Example_3",
			s:    "(]",
			want: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isValid(tt.s); got != tt.want {
				t.Errorf("isValid() = %v, want %v", got, tt.want)
			}
		})
	}
}
'@
Set-Content -Path "problems/020_valid_parentheses/solution_test.go" -Value $testContent -Encoding UTF8
Write-Host "Created/Updated problems/020_valid_parentheses/solution_test.go"

Write-Host "`nAll files created/updated successfully!"
