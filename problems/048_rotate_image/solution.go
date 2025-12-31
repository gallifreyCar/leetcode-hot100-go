package rotateimage

func rotate(matrix [][]int) {
	//瀵硅绾垮弽杞?
	for i := 0; i < len(matrix); i++ {
		for j := 0; j < i; j++ {
			matrix[i][j], matrix[j][i] = matrix[j][i], matrix[i][j]
		}
	}
	//宸﹀彸缈昏浆
	for i := 0; i < len(matrix); i++ {
		left, right := 0, len(matrix[0])-1
		for left < right {
			matrix[i][left], matrix[i][right] = matrix[i][right], matrix[i][left]
			left++
			right--
		}
	}
}
