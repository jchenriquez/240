package search2d

func search2D (matrix [][]int, target, rowStart, rowEnd, columnStart, columnEnd int) bool {
  if rowEnd - rowStart == 0 && columnEnd - columnStart == 0 {
    return matrix[rowStart][columnStart] == target
  }

  columnMid := (columnEnd+columnStart) / 2
	rowMid := (rowEnd+rowStart) / 2

	var newColumnStart int
	var newColumnEnd int
	var newRowStart int
	var newRowEnd int

	if matrix[rowStart][columnMid] == target {
		return true
	}

	if matrix[rowMid][columnStart] == target {
		return true
	}

	if target > matrix[rowStart][columnMid] {
		newColumnStart = columnMid
		newColumnEnd = columnEnd
	} else {
		newColumnStart = columnStart
		newColumnEnd = columnMid
	}

	if target > matrix[columnMid][rowMid] {
		newRowStart = rowMid
		newColumnEnd = rowEnd
	} else {
		newRowStart = rowStart
		newRowEnd = rowMid
	}

  return search2D(matrix, target, newRowStart, newRowEnd, newColumnStart, newColumnEnd)
}

func SearchMatrix(matrix [][]int, target int) bool {
  if len(matrix) == 0{
    return false
  }
  return search2D(matrix, target, 0, len(matrix), 0, len(matrix[0]))
}
