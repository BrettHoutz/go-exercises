package leftrotate

// RotateLeft creates a new slice with the elements in slice rotated by n positions.
func RotateLeft(slice []int, n int) []int {
	rotated := make([]int, len(slice))
	for i, v := range slice {
		rotIndex := (i - n) % len(slice)
		if rotIndex < 0 {
			rotIndex += 5
		}

		rotated[rotIndex] = v
	}
	return rotated
}
