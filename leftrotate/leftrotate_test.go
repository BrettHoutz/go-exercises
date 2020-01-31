package leftrotate

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLeftRotate(t *testing.T) {
	slc := []int{1, 2, 3, 5, 8}

	assert.Equal(t, []int{5, 8, 1, 2, 3}, RotateLeft(slc, 3), "rotates elements left")
	assert.Equal(t, []int{1, 2, 3, 5, 8}, RotateLeft(slc, 0), "does nothing when n = 0")
	assert.Equal(t, []int{5, 8, 1, 2, 3}, RotateLeft(slc, 8), "rotates elements left when n > len(slice)")
	assert.Equal(t, []int{5, 8, 1, 2, 3}, RotateLeft(slc, -2), "rotates elements right when n is negative")
}
