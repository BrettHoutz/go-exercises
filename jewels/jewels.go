package jewels

// NumJewels takes a string of runes that represent jewels
// and a string of runes that represent stones you have,
// and returns the number of jewels that you have.
func NumJewels(jewels, stones string) int {
	jewelSet := map[rune]bool{}
	for _, jewel := range jewels {
		jewelSet[jewel] = true
	}

	numJewels := 0
	for _, stone := range stones {
		if jewelSet[stone] {
			numJewels++
		}
	}
	return numJewels
}
