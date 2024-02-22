package boyermooremajorityvoting

// https://www.geeksforgeeks.org/boyer-moore-majority-voting-algorithm/
func FindMajorityElement(array []int) int {
	candidate, count := 0, 0

	// Finding majority candidate
	for i := 0; i < len(array); i++ {
		if count == 0 {
			candidate = array[i]
			count = 1
		} else if array[i] == candidate {
			count++
		} else {
			count--
		}
	}

	count = 0
	for i := 0; i < len(array); i++ {
		if array[i] == candidate {
			count++
		}
	}

	if count > len(array)/2 {
		return candidate
	}

	return -1
}
