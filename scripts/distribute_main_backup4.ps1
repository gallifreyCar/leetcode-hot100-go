# 分发更多题解

# 021_merge_two_sorted_lists (这个已经有了，跳过)

# 189_rotate_array
$content = @'
package rotatearray

import "slices"

func rotate(nums []int, k int) {
	k %= len(nums)
	//三次反转，原地算法
	slices.Reverse(nums)
	slices.Reverse(nums[:k])
	slices.Reverse(nums[k:])
}
'@
Set-Content -Path "problems/189_rotate_array/solution.go" -Value $content -Encoding UTF8
Write-Host "Created/Updated problems/189_rotate_array/solution.go"

$testContent = @'
package rotatearray

import (
	"reflect"
	"testing"
)

func TestRotate(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		k    int
		want []int
	}{
		{
			name: "Example_1",
			nums: []int{1, 2, 3, 4, 5, 6, 7},
			k:    3,
			want: []int{5, 6, 7, 1, 2, 3, 4},
		},
		{
			name: "Example_2",
			nums: []int{-1, -100, 3, 99},
			k:    2,
			want: []int{3, 99, -1, -100},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			nums := make([]int, len(tt.nums))
			copy(nums, tt.nums)
			rotate(nums, tt.k)
			if !reflect.DeepEqual(nums, tt.want) {
				t.Errorf("rotate() = %v, want %v", nums, tt.want)
			}
		})
	}
}
'@
Set-Content -Path "problems/189_rotate_array/solution_test.go" -Value $testContent -Encoding UTF8
Write-Host "Created/Updated problems/189_rotate_array/solution_test.go"

# 073_set_matrix_zeroes
$content = @'
package setmatrixzeroes

func setZeroes(matrix [][]int) {
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
'@
Set-Content -Path "problems/073_set_matrix_zeroes/solution.go" -Value $content -Encoding UTF8
Write-Host "Created/Updated problems/073_set_matrix_zeroes/solution.go"

$testContent = @'
package setmatrixzeroes

import (
	"reflect"
	"testing"
)

func TestSetZeroes(t *testing.T) {
	tests := []struct {
		name   string
		matrix [][]int
		want   [][]int
	}{
		{
			name:   "Example_1",
			matrix: [][]int{{1, 1, 1}, {1, 0, 1}, {1, 1, 1}},
			want:   [][]int{{1, 0, 1}, {0, 0, 0}, {1, 0, 1}},
		},
		{
			name:   "Example_2",
			matrix: [][]int{{0, 1, 2, 0}, {3, 4, 5, 2}, {1, 3, 1, 5}},
			want:   [][]int{{0, 0, 0, 0}, {0, 4, 5, 0}, {0, 3, 1, 0}},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			matrix := make([][]int, len(tt.matrix))
			for i := range tt.matrix {
				matrix[i] = make([]int, len(tt.matrix[i]))
				copy(matrix[i], tt.matrix[i])
			}
			setZeroes(matrix)
			if !reflect.DeepEqual(matrix, tt.want) {
				t.Errorf("setZeroes() = %v, want %v", matrix, tt.want)
			}
		})
	}
}
'@
Set-Content -Path "problems/073_set_matrix_zeroes/solution_test.go" -Value $testContent -Encoding UTF8
Write-Host "Created/Updated problems/073_set_matrix_zeroes/solution_test.go"

# 054_spiral_matrix
$content = @'
package spiralmatrix

func spiralOrder(matrix [][]int) []int {
	m, n := len(matrix), len(matrix[0])
	left, right, top, bottom := 0, n-1, 0, m-1
	res := make([]int, 0, m*n)
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
'@
Set-Content -Path "problems/054_spiral_matrix/solution.go" -Value $content -Encoding UTF8
Write-Host "Created/Updated problems/054_spiral_matrix/solution.go"

$testContent = @'
package spiralmatrix

import (
	"reflect"
	"testing"
)

func TestSpiralOrder(t *testing.T) {
	tests := []struct {
		name   string
		matrix [][]int
		want   []int
	}{
		{
			name:   "Example_1",
			matrix: [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}},
			want:   []int{1, 2, 3, 6, 9, 8, 7, 4, 5},
		},
		{
			name:   "Example_2",
			matrix: [][]int{{1, 2, 3, 4}, {5, 6, 7, 8}, {9, 10, 11, 12}},
			want:   []int{1, 2, 3, 4, 8, 12, 11, 10, 9, 5, 6, 7},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := spiralOrder(tt.matrix); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("spiralOrder() = %v, want %v", got, tt.want)
			}
		})
	}
}
'@
Set-Content -Path "problems/054_spiral_matrix/solution_test.go" -Value $testContent -Encoding UTF8
Write-Host "Created/Updated problems/054_spiral_matrix/solution_test.go"

# 048_rotate_image
$content = @'
package rotateimage

func rotate(matrix [][]int) {
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
'@
Set-Content -Path "problems/048_rotate_image/solution.go" -Value $content -Encoding UTF8
Write-Host "Created/Updated problems/048_rotate_image/solution.go"

$testContent = @'
package rotateimage

import (
	"reflect"
	"testing"
)

func TestRotate(t *testing.T) {
	tests := []struct {
		name   string
		matrix [][]int
		want   [][]int
	}{
		{
			name:   "Example_1",
			matrix: [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}},
			want:   [][]int{{7, 4, 1}, {8, 5, 2}, {9, 6, 3}},
		},
		{
			name:   "Example_2",
			matrix: [][]int{{5, 1, 9, 11}, {2, 4, 8, 10}, {13, 3, 6, 7}, {15, 14, 12, 16}},
			want:   [][]int{{15, 13, 2, 5}, {14, 3, 4, 1}, {12, 6, 8, 9}, {16, 7, 10, 11}},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			matrix := make([][]int, len(tt.matrix))
			for i := range tt.matrix {
				matrix[i] = make([]int, len(tt.matrix[i]))
				copy(matrix[i], tt.matrix[i])
			}
			rotate(matrix)
			if !reflect.DeepEqual(matrix, tt.want) {
				t.Errorf("rotate() = %v, want %v", matrix, tt.want)
			}
		})
	}
}
'@
Set-Content -Path "problems/048_rotate_image/solution_test.go" -Value $testContent -Encoding UTF8
Write-Host "Created/Updated problems/048_rotate_image/solution_test.go"

# 041_first_missing_positive
$content = @'
package firstmissingpositive

func firstMissingPositive(nums []int) int {
	//原地算法，用下标来标记这个数有没有出现过
	//用正负数来标记（下标i-1对应数字i)

	// 原来的0或负数直接置为用边界外的任意一个正数（不用处理
	for i, num := range nums {
		if num <= 0 {
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
'@
Set-Content -Path "problems/041_first_missing_positive/solution.go" -Value $content -Encoding UTF8
Write-Host "Created/Updated problems/041_first_missing_positive/solution.go"

$testContent = @'
package firstmissingpositive

import "testing"

func TestFirstMissingPositive(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want int
	}{
		{
			name: "Example_1",
			nums: []int{1, 2, 0},
			want: 3,
		},
		{
			name: "Example_2",
			nums: []int{3, 4, -1, 1},
			want: 2,
		},
		{
			name: "Example_3",
			nums: []int{7, 8, 9, 11, 12},
			want: 1,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			nums := make([]int, len(tt.nums))
			copy(nums, tt.nums)
			if got := firstMissingPositive(nums); got != tt.want {
				t.Errorf("firstMissingPositive() = %v, want %v", got, tt.want)
			}
		})
	}
}
'@
Set-Content -Path "problems/041_first_missing_positive/solution_test.go" -Value $testContent -Encoding UTF8
Write-Host "Created/Updated problems/041_first_missing_positive/solution_test.go"

Write-Host "`nAll files created/updated successfully!"
