package merge

// TwoSortedIntArrays ...
func TwoSortedIntArrays(a []int, b []int) []int {
	aPtr, bPtr := 0, 0
	c := []int{}

	// Traverse both array
	for aPtr < len(a) && bPtr < len(b) {
		if a[aPtr] < b[bPtr] {
			c = append(c, a[aPtr])
			aPtr++
		} else {
			c = append(c, b[bPtr])
			bPtr++
		}
	}

	if aPtr < len(a) {
		c = append(c, a[aPtr:]...)
	}
	if bPtr < len(b) {
		c = append(c, b[bPtr:]...)
	}

	return c
}
