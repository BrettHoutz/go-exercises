package searchmatrix

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSearchMatrix(t *testing.T) {
	matrix := [][]int{
		[]int{1, 2, 3},
		[]int{5, 8, 13},
		[]int{21, 34, 55},
	}

	row, col, ok := SearchMatrix(matrix, 8)
	assert.Equal(t, ok, true, "returns ok=true when the target is present")
	assert.Equal(t, row, 1, "returns the row when the target is present")
	assert.Equal(t, col, 1, "returns the col when the target is present")

	row, col, ok = SearchMatrix(matrix, 7)
	assert.Equal(t, ok, false, "returns ok=false when the target is not present")

	row, col, ok = SearchMatrix(matrix, 1)
	assert.Equal(t, ok, true, "finds the first element")
	assert.Equal(t, row, 0, "returns the row for the first element")
	assert.Equal(t, col, 0, "returns the col for the first element")

	row, col, ok = SearchMatrix(matrix, 55)
	assert.Equal(t, ok, true, "finds the first element")
	assert.Equal(t, row, 2, "returns the row for the first element")
	assert.Equal(t, col, 2, "returns the col for the first element")

	matrix = [][]int{}
	row, col, ok = SearchMatrix(matrix, 8)
	assert.Equal(t, ok, false, "works for an empty slice")

	matrix = [][]int{
		[]int{},
		[]int{},
		[]int{},
	}
	row, col, ok = SearchMatrix(matrix, 8)
	assert.Equal(t, ok, false, "works for a slice of empty slices")

	matrix = [][]int{
		[]int{1, 2, 3},
		[]int{},
		[]int{21, 34, 55},
	}

	row, col, ok = SearchMatrix(matrix, 8)
	assert.Equal(t, ok, false, "works when the expected row is empty")
}
