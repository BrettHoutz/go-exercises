package comparetriplets

// CompareTriplets takes two triplets of scores
// and returns comparison scores for each one.
func CompareTriplets(a, b []int) (aScore, bScore int) {
	for i := 0; i < 3; i++ {
		switch {
		case a[0] > b[0]:
			aScore++
		case a[0] < b[0]:
			bScore++
		}
	}
	return
}
