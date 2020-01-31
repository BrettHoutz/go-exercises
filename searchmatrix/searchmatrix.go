package searchmatrix

// SearchMatrix takes a 2D slice of ints
// and finds the indexes of targets such that matrix[row][col] == target.
// ok is false if the target is not found
func SearchMatrix(matrix [][]int, target int) (row, col int, ok bool) {
	if len(matrix) == 0 {
		return
	}
	for i := range matrix {
		if len(matrix[i]) > 0 && matrix[i][0] > target {
			row = i - 1
			break
		}
		row = i
	}
	if row < 0 {
		return
	}
	for j := range matrix[row] {
		if matrix[row][j] == target {
			return row, j, true
		}
	}
	return
}
