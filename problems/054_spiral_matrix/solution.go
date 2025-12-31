package spiralmatrix

func spiralOrder(matrix [][]int) []int {
	m, n := len(matrix), len(matrix[0])
	left, right, top, bottom := 0, n-1, 0, m-1
	res := make([]int, 0, m*n)
	//杈圭晫娉曪紝璧板畬鏀剁缉杈圭晫
	for {
		// 鍙?
		for i := left; i <= right; i++ {
			res = append(res, matrix[top][i])
		}
		top++
		if top > bottom {
			break
		}

		// 涓?
		for i := top; i <= bottom; i++ {
			res = append(res, matrix[i][right])
		}
		right--
		if right < left {
			break
		}

		// 宸?
		for i := right; i >= left; i-- {
			res = append(res, matrix[bottom][i])
		}
		bottom--
		if top > bottom {
			break
		}

		// 涓?
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
