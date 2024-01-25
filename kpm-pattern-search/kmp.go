package kpmpatternsearch

func FindSubStrings(text string, pattern string) []int {
	return make([]int, 0)
}

func ComputeLPSArray(pattern string) []int {
	prefixLength, i := 0, 1
	lpsArray := make([]int, len(pattern))
	lpsArray[0] = 0

	for i < len(pattern) {
		if pattern[prefixLength] == pattern[i] {
			lpsArray[i] = prefixLength + 1
			prefixLength++
			i++
		} else if prefixLength != 0 {
			prefixLength = lpsArray[prefixLength-1]
		} else {
			lpsArray[i] = 0
			i++
		}
	}

	return lpsArray
}
