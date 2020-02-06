package search2d

import (
	"sort"
)

//SearchMatrix will search a target in a matrix using binary search.
func SearchMatrix(matrix [][]int, target int) bool {
	if len(matrix) == 0 {
		return false
	}

	for i := 0; i < len(matrix); i++ {
		column := matrix[i]

		index := sort.Search(len(column), func(j int) bool {
			return column[j] >= target
		})

		if index < len(column) && column[index] == target {
			return true
		}
	}

	return false
}
