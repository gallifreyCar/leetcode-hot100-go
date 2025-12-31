package setmatrixzeroes

func setZeroes(matrix [][]int) {
	//绫诲瀷excel 鐢ㄧ涓€鍒楀拰绗竴琛屾潵瀛樺偍0淇℃伅

	//璁板綍绗竴鍒楀拰绗竴琛岀殑鏄惁瀛樺湪0
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

	//璁板綍0 璺宠繃绗竴琛?鍒楃殑鍒ゆ柇
	for i := 1; i < len(matrix); i++ {
		for j := 1; j < len(matrix[i]); j++ {
			if matrix[i][j] == 0 {
				matrix[0][j] = 0
				matrix[i][0] = 0
			}
		}
	}

	//鍘熷湴绠楁硶
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
