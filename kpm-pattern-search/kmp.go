package kpmpatternsearch

func FindIndexOfSubStrings(text string, pattern string) []int {
	matchesArray := make([]int, 0)

	textLength, patternLength := len(text), len(pattern)
	textIndex, patternIndex := 0, 0
	lpsArray := ComputeLPSArray(pattern)

	for textIndex < textLength {
		if text[textIndex] == pattern[patternIndex] {
			textIndex++
			patternIndex++

			if patternIndex == patternLength {
				matchesArray = append(matchesArray, textIndex-patternIndex)
				patternIndex = lpsArray[patternIndex-1]
			}
		} else {
			if patternIndex != 0 {
				patternIndex = lpsArray[patternIndex-1]
			} else {
				textIndex++
			}
		}
	}

	return matchesArray
}

// https://www.youtube.com/watch?v=4jY57Ehc14Y&t=551s
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
