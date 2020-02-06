package search2d

func SearchMatrix(matrix [][]int, target int) bool {
	if len(matrix) == 1 && len(matrix[0]) == 1 {
		return matrix[0][0] == target
	}

	columnMid := len(matrix[0]) / 2
	rowMid := len(matrix) / 2

	var newColumnStart int
	var newColumnEnd int
	var newRowStart int
	var newRowEnd int

	if matrix[0][columnMid] == target {
		return true
	}

	if matrix[rowMid][0] == target {
		return true
	}

	if target > matrix[0][columnMid] {
		newColumnStart = columnMid
		newColumnEnd = len(matrix[0])
	} else {
		newColumnStart = 0
		newColumnEnd = columnMid
	}

	if target > matrix[columnMid][rowMid] {
		newRowStart = rowMid
		newColumnEnd = len(matrix)
	} else {
		newRowStart = 0
		newRowEnd = rowMid
	}

}
