package comparetriplets

// CompareTriplets takes two triplets of scores
// and returns comparison scores for each one.
func CompareTriplets(a, b []int) (aScore, bScore int) {
	switch {
	case a[0] > b[0]:
		aScore++
	case a[0] < b[0]:
		bScore++
	}

	switch {
	case a[1] > b[1]:
		aScore++
	case a[1] < b[1]:
		bScore++
	}

	switch {
	case a[2] > b[2]:
		aScore++
	case a[2] < b[2]:
		bScore++
	}
	return
}
